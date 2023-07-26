/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-27 01:05
**/

package filter

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strings"
	"sync"
)

type Type int

const (
	Integer Type = 1 << iota
	Float
	String
	Bool
	All = Integer | Float | String | Bool
)

type Filter struct {
	visited map[uintptr]bool
	deep    int
	tag     string
	flag    Type
	mux     sync.Mutex
}

//func (m *Filter) Force(b bool) {
//	m.force = b
//}

func (m *Filter) Tag(tag string) *Filter {
	m.tag = tag
	return m
}

func (m *Filter) Flag(flag Type) *Filter {
	m.flag = flag
	return m
}

func New() *Filter {
	return &Filter{
		visited: make(map[uintptr]bool),
		deep:    0,
		tag:     "bson",
		flag:    All,
	}
}

func (m *Filter) Zero(t any) bson.M {

	m.mux.Lock()
	defer m.mux.Unlock()

	var rv = reflect.ValueOf(t)

	var properties = bson.M{}

	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil
	}

	m.printStruct(properties, "", rv)

	return properties
}

func (m *Filter) format(mapping map[string]any, key string, rv reflect.Value, tag reflect.StructTag) {

	var tagStr = tag.Get(m.tag)
	if tagStr == "" {
		return
	}

	var tagArr = strings.Split(tagStr, ",")
	if tagArr[0] == "-" {
		return
	}

	if rv.IsZero() {
		if m.shouldIgnore(rv) {
			return
		}

		if len(tagArr) > 1 && tagArr[1] == "omitempty" {
			return
		}
	}

	var name = tagArr[0]

	if name == "" {
		name = key
	}

	switch rv.Kind() {

	// SIMPLE TYPE
	case reflect.Bool:
		// ignore
		mapping[name] = rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		// ignore
		mapping[name] = rv.Int()
	case reflect.Float32, reflect.Float64:
		// ignore
		mapping[name] = rv.Float()
	case reflect.String:
		// ignore
		mapping[name] = rv.String()
	case reflect.Func:
		// ignore
	case reflect.UnsafePointer:
		// ignore
	case reflect.Chan:
		// ignore
	case reflect.Invalid:
		// ignore

	// COMPLEX TYPE
	case reflect.Map:
		var newMapping = bson.M{}
		mapping[name] = newMapping
		m.printMap(newMapping, name, rv, tag)
	case reflect.Struct:
		var newMapping = bson.M{}
		mapping[name] = newMapping
		m.printStruct(newMapping, name, rv)
	case reflect.Array, reflect.Slice:
		m.printSlice(mapping, name, rv, tag)
	case reflect.Ptr:
		m.printPtr(mapping, name, rv, tag)
	case reflect.Interface:
		m.format(mapping, name, rv.Elem(), tag)
	default:
		// ignore
	}
}

func (m *Filter) printMap(mapping map[string]any, key string, v reflect.Value, tag reflect.StructTag) {

	var d = m.deep
	m.deep++

	if v.Len() == 0 {
		// ignore
		// cuz you don't know the type of key and value
		m.deep = d
		return
	}

	if m.visited[v.Pointer()] {
		// ignore
		// repeat reference
		m.deep = d
		return
	}

	m.visited[v.Pointer()] = true

	keys := v.MapKeys()
	for i := 0; i < v.Len(); i++ {
		value := v.MapIndex(keys[i])
		var name = keys[i].String()
		m.format(mapping, name, value, tag)
	}

	m.deep = d
}

func (m *Filter) printStruct(mapping map[string]any, key string, v reflect.Value) {

	var d = m.deep
	m.deep++

	if v.NumField() == 0 {
		// ignore
		// cuz you don't know the type of key and value
		m.deep = d
		return
	}

	for i := 0; i < v.NumField(); i++ {
		var field = v.Type().Field(i)

		value := v.Field(i)

		if value.CanInterface() {
			var name = field.Name
			m.format(mapping, name, value, field.Tag)
		}
	}

	m.deep = d
}

func (m *Filter) printSlice(mapping map[string]any, key string, v reflect.Value, tag reflect.StructTag) {

	var d = m.deep
	m.deep++

	if v.Len() == 0 {
		// ignore
		// cuz you don't know the type of key and value
		m.deep = d
		return
	}

	//  if is array, will be handled in printPtr
	if v.Kind() == reflect.Slice {
		if m.visited[v.Pointer()] {
			// repeat reference
			m.deep = d
			return
		}
		m.visited[v.Pointer()] = true
	}

	var arr = make([]any, 0)
	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)
		var v = m.removeZero(value)
		if v == nil {
			continue
		}
		arr = append(arr, v)
	}
	mapping[key] = arr
	m.deep = d
}

func (m *Filter) removeZero(i reflect.Value) any {
	if i.IsZero() {
		if m.shouldIgnore(i) {
			return nil
		}
	}
	switch i.Kind() {
	case reflect.Struct:
		var res = bson.M{}
		for j := 0; j < i.NumField(); j++ {
			var field = i.Type().Field(j)
			var value = i.Field(j)
			if value.IsZero() {
				continue
			}
			res[field.Name] = m.removeZero(value)
		}
		return res
	case reflect.Slice, reflect.Array:
		var arr = make([]any, 0)
		for j := 0; j < i.Len(); j++ {
			var value = i.Index(j)
			if value.IsZero() {
				continue
			}
			arr = append(arr, m.removeZero(value))
		}
		return arr
	case reflect.Map:
		var res = bson.M{}
		var keys = i.MapKeys()
		for j := 0; j < i.Len(); j++ {
			var value = i.MapIndex(keys[j])
			if value.IsZero() {
				continue
			}
			res[keys[j].String()] = m.removeZero(value)
		}
		return res
	case reflect.Interface:
		if i.IsNil() {
			return nil
		}
		return m.removeZero(i.Elem())
	case reflect.Ptr:
		if i.Pointer() == 0 {
			return nil
		}
		return m.removeZero(i.Elem())
	case reflect.Invalid:
		return nil
	default:
		if i.CanInterface() {
			return i.Interface()
		}
		return nil
	}
}

func (m *Filter) printPtr(mapping map[string]any, key string, v reflect.Value, tag reflect.StructTag) {

	if m.visited[v.Pointer()] {
		// repeat reference
		return
	}

	if v.Pointer() != 0 {
		m.visited[v.Pointer()] = true
	}

	if v.Elem().IsValid() {
		m.format(mapping, key, v.Elem(), tag)
	}
}

func (m *Filter) shouldIgnore(i reflect.Value) bool {
	switch i.Kind() {
	case reflect.Bool:
		return m.flag&Bool != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		return m.flag&Integer != 0
	case reflect.Float32, reflect.Float64:
		return m.flag&Float != 0
	case reflect.String:
		return m.flag&String != 0
	default:
		return true
	}
}
