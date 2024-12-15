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
	"github.com/lemonyxk/longo/call"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertMany struct {
	collection       *mongo.Collection
	insertManyOption *options.InsertManyOptions
	document         []interface{}
	sessionContext   context.Context
}

func NewInsertMany(ctx context.Context, collection *mongo.Collection, document []interface{}) *InsertMany {
	return &InsertMany{collection: collection, insertManyOption: &options.InsertManyOptions{}, document: document, sessionContext: ctx}
}

func (f *InsertMany) Option(opt *options.InsertManyOptions) *InsertMany {
	f.insertManyOption = opt
	return f
}

func (f *InsertMany) Context(ctx context.Context) *InsertMany {
	f.sessionContext = ctx
	return f
}

func (f *InsertMany) Exec() (*mongo.InsertManyResult, error) {

	var t = time.Now()
	var res *mongo.InsertManyResult
	var err error

	defer func() {
		if res == nil {
			res = &mongo.InsertManyResult{}
		}
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.InsertMany,
			},
			Query: call.Query{
				Filter:  nil,
				Updater: nil,
			},
			Result: call.Result{
				Insert: int64(len(res.InsertedIDs)),
				Update: 0,
				Delete: 0,
				Match:  0,
				Upsert: 0,
			},
			Consuming: time.Since(t).Microseconds(),
			Error:     err,
		})
	}()

	res, err = f.collection.InsertMany(f.sessionContext, f.document, f.insertManyOption)
	if err != nil {
		return nil, err
	}
	if len(res.InsertedIDs) == 0 {
		return nil, fmt.Errorf("insert many error: %s", "no id")
	}
	return res, nil
}
