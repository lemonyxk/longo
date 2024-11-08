/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:28
**/

package longo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type MultiResult struct {
	cursor *mongo.Cursor
	err    error
}

func (ag *MultiResult) All(sessionContext context.Context, result interface{}) error {
	if ag.err != nil {
		return ag.err
	}
	// refResult := reflect.ValueOf(result)
	// if refResult.Kind() != reflect.Ptr || refResult.Elem().Kind() != reflect.Slice {
	// 	return errors.New("result argument must be a slice address")
	// }

	if ag.cursor == nil {
		return fmt.Errorf("cursor is nil")
	}

	if ag.cursor.Err() != nil {
		return ag.cursor.Err()
	}

	return ag.cursor.All(sessionContext, result)
}

//func (ag *MultiResult) One(sessionContext context.Context, result interface{}) error {
//	if ag.err != nil {
//		return ag.err
//	}
//
//	if ag.cursor == nil {
//		return fmt.Errorf("cursor is nil")
//	}
//
//	if ag.cursor.Err() != nil {
//		return ag.cursor.Err()
//	}
//
//	ag.cursor.Next(sessionContext)
//	var err = ag.cursor.Decode(result)
//	if err == io.EOF {
//		err = nil
//	}
//	return err
//}
