/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-06 02:01
**/

package model

import (
	"embed"
	"strings"
)

//go:embed temp
var temp embed.FS

func create(packageName, name, db, collection, id, t string) []byte {

	if name == "" || id == "" {
		return nil
	}

	var tv = ""
	if t == "int" {
		tv = `0`
	} else if t == "string" {
		tv = `""`
	}

	if tv == "" {
		return nil
	}

	var res = createEmptyMethod(name, id, tv)

	res += createModel(name, db, collection)

	return []byte(res)
}

func createModel(name, db, collection string) string {
	var model, _ = temp.ReadFile("temp/model.temp")
	var _name = []byte(name)
	_name[0] = _name[0] + 32
	var res = replaceString(string(model), []string{"@name", "@_name", "@db", "@collection"}, []string{name, string(_name), db, collection})
	return res + "\n\n"
}

func createEmptyMethod(name string, id string, tv string) string {
	var emptyMethod, _ = temp.ReadFile("temp/empty_method.temp")
	var res = replaceString(string(emptyMethod), []string{"@name", "@id", "@tv"}, []string{name, id, tv})
	return res + "\n\n"
}

func replaceString(s string, oList []string, nList []string) string {
	for i := 0; i < len(oList); i++ {
		s = strings.ReplaceAll(s, oList[i], nList[i])
	}
	return s
}
