/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2020-12-14 11:50
**/

package longo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOne struct {
	collection      *mongo.Collection
	insertOneOption *options.InsertOneOptions
	document        interface{}
	sessionContext  context.Context
}

func NewInsertOne(ctx context.Context, collection *mongo.Collection, document interface{}) *InsertOne {
	return &InsertOne{collection: collection, insertOneOption: &options.InsertOneOptions{}, document: document, sessionContext: ctx}
}

func (f *InsertOne) Option(opt *options.InsertOneOptions) *InsertOne {
	f.insertOneOption = opt
	return f
}

func (f *InsertOne) Context(ctx context.Context) *InsertOne {
	f.sessionContext = ctx
	return f
}

func (f *InsertOne) Exec() (*mongo.InsertOneResult, error) {

	//var t = time.Now()
	//var res *mongo.InsertOneResult
	//var err error
	//
	//defer func() {
	//	if res == nil {
	//		res = &mongo.InsertOneResult{}
	//	}
	//	call.Default.Call(call.Record{
	//		Meta: call.Meta{
	//			Database:   f.collection.Database().Name(),
	//			Collection: f.collection.Name(),
	//			Type:       call.InsertOne,
	//		},
	//		Query: call.Query{
	//			Filter:  nil,
	//			Updater: nil,
	//		},
	//		Result: call.Result{
	//			Insert: 1,
	//			Update: 0,
	//			Delete: 0,
	//			Match:  0,
	//			Upsert: 0,
	//		},
	//		Consuming: time.Since(t).Microseconds(),
	//		Error:     err,
	//	})
	//}()

	res, err := f.collection.InsertOne(f.sessionContext, f.document, f.insertOneOption)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == nil {
		return nil, fmt.Errorf("insert one error: %s", "no id")
	}
	return res, nil
}
