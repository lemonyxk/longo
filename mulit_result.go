/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:28
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MultiResult struct {
	cursor *mongo.Cursor
	err    error
}

func (ag *MultiResult) All(result interface{}) error {
	if ag.err != nil {
		return ag.err
	}
	// refResult := reflect.ValueOf(result)
	// if refResult.Kind() != reflect.Ptr || refResult.Elem().Kind() != reflect.Slice {
	// 	return errors.New("result argument must be a slice address")
	// }
	return ag.cursor.All(context.Background(), result)
}

func (ag *MultiResult) One(result interface{}) error {
	if ag.err != nil {
		return ag.err
	}
	ag.cursor.Next(context.Background())
	return ag.cursor.Decode(result)
}
