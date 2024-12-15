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
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Aggregate struct {
	collection     *mongo.Collection
	option         *options.AggregateOptions
	pipeline       interface{}
	sessionContext context.Context
}

func NewAggregate(ctx context.Context, collection *mongo.Collection, pipeline interface{}) *Aggregate {
	return &Aggregate{collection: collection, option: &options.AggregateOptions{}, pipeline: pipeline, sessionContext: ctx}
}

func (f *Aggregate) Context(ctx context.Context) *Aggregate {
	f.sessionContext = ctx
	return f
}

func (f *Aggregate) Hit(hit interface{}) *Aggregate {
	f.option.Hint = hit
	return f
}

func (f *Aggregate) Option(opt *options.AggregateOptions) *Aggregate {
	f.option = opt
	return f
}

func (f *Aggregate) All(result interface{}) error {

	var t = time.Now()
	var res int64 = 0
	var err error

	defer func() {
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.Aggregate,
			},
			Query: call.Query{
				Filter:  f.pipeline,
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

	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, f.option)
	var all = &MultiResult{cursor: cursor, err: err}
	err = all.All(f.sessionContext, result)
	if err != nil {
		return err
	}
	res = int64(reflect.ValueOf(result).Elem().Len())
	return nil
}
