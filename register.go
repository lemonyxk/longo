/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2020-01-09 16:52
**/

package longo

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type NullDecoder struct {
	defDecoder bsoncodec.ValueDecoder
	zeroValue  reflect.Value
}

func (d *NullDecoder) DecodeValue(ctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if vr.Type() != bsontype.Null {
		return d.defDecoder.DecodeValue(ctx, vr, val)
	}

	if !val.CanSet() {
		return errors.New("value not settable")
	}
	if err := vr.ReadNull(); err != nil {
		return err
	}
	// Set the zero value of val's type:
	val.Set(d.zeroValue)
	return nil
}

func NewNUll() *bsoncodec.RegistryBuilder {
	var customValues = []interface{}{
		"", // string
		[]string{},
		int32(0),
		int64(0),   // int32
		float32(0), // int32
		float64(0), // int32
		[]int32{},
		[]int64{},
		[]float32{},
		[]float64{},
		false,
		[]bool{},
	}

	var rb = bson.NewRegistryBuilder()

	for i := 0; i < len(customValues); i++ {
		var t = reflect.TypeOf(customValues[i])
		defDecoder, err := bson.DefaultRegistry.LookupDecoder(t)
		if err != nil {
			panic(err)
		}
		rb.RegisterTypeDecoder(t, &NullDecoder{defDecoder, reflect.Zero(t)})
	}

	return rb
}
