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

type FindOne struct {
	collection     *mongo.Collection
	option         *options.FindOneOptions
	filter         interface{}
	sessionContext context.Context
}

func NewFindOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *FindOne {
	return &FindOne{collection: collection, option: &options.FindOneOptions{}, filter: filter, sessionContext: ctx}
}

func (f *FindOne) Hit(hit interface{}) *FindOne {
	f.option.Hint = hit
	return f
}

func (f *FindOne) Sort(sort interface{}) *FindOne {
	f.option.Sort = sort
	return f
}

func (f *FindOne) Skip(skip int64) *FindOne {
	f.option.Skip = &skip
	return f
}

func (f *FindOne) Projection(projection interface{}) *FindOne {
	f.option.Projection = projection
	return f
}

func (f *FindOne) Context(ctx context.Context) *FindOne {
	f.sessionContext = ctx
	return f
}

func (f *FindOne) Option(opt *options.FindOneOptions) *FindOne {
	f.option = opt
	return f
}

func (f *FindOne) Get(result interface{}) error {

	var t = time.Now()

	var res int64 = 0
	var err error

	defer func() {
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.FindOne,
			},
			Query: call.Query{
				Filter:  f.filter,
				Updater: nil,
			},
			Result: call.Result{
				Insert: 0,
				Update: 0,
				Delete: 0,
				Match:  res,
				Upsert: 0,
			},
			Consuming: time.Since(t).Microseconds(),
			Error:     err,
		})
	}()

	var cursor = &SingleResult{singleResult: f.collection.FindOne(f.sessionContext, f.filter, f.option)}
	err = cursor.Get(result)
	if err == mongo.ErrNoDocuments { // only for FindOne, if modify action, it will return ErrNoDocuments
		err = nil
	}
	if err != nil {
		return err
	}
	res = 1
	return nil
}
