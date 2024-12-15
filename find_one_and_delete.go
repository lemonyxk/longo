/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-29 14:00
**/

package longo

import (
	"context"
	"github.com/lemonyxk/longo/call"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndDelete struct {
	collection     *mongo.Collection
	option         *options.FindOneAndDeleteOptions
	filter         interface{}
	sessionContext context.Context
}

func NewFindOneAndDelete(ctx context.Context, collection *mongo.Collection, filter interface{}) *FindOneAndDelete {
	return &FindOneAndDelete{collection: collection, option: &options.FindOneAndDeleteOptions{}, filter: filter, sessionContext: ctx}
}

func (f *FindOneAndDelete) Hit(hit interface{}) *FindOneAndDelete {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndDelete) Sort(sort interface{}) *FindOneAndDelete {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndDelete) Projection(projection interface{}) *FindOneAndDelete {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndDelete) Context(ctx context.Context) *FindOneAndDelete {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndDelete) Option(opt *options.FindOneAndDeleteOptions) *FindOneAndDelete {
	f.option = opt
	return f
}

func (f *FindOneAndDelete) Exec(result interface{}) error {

	var t = time.Now()

	var res int64 = 0
	var err error

	defer func() {
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.FindOneAndDelete,
			},
			Query: call.Query{
				Filter:  f.filter,
				Updater: nil,
			},
			Result: call.Result{
				Insert: 0,
				Update: 0,
				Delete: res,
				Match:  res,
				Upsert: 0,
			},
			Consuming: time.Since(t).Microseconds(),
			Error:     err,
		})
	}()

	var cursor = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	err = cursor.Get(result)
	if err != nil {
		return err
	}
	res = 1
	return err
}
