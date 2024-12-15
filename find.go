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
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Find struct {
	collection     *mongo.Collection
	option         *options.FindOptions
	filter         interface{}
	sessionContext context.Context
}

func NewFind(ctx context.Context, collection *mongo.Collection, filter interface{}) *Find {
	return &Find{collection: collection, option: &options.FindOptions{}, filter: filter, sessionContext: ctx}
}

func (f *Find) Hit(hit interface{}) *Find {
	f.option.Hint = hit
	return f
}

func (f *Find) Sort(sort interface{}) *Find {
	f.option.Sort = sort
	return f
}

func (f *Find) Limit(limit int64) *Find {
	f.option.Limit = &limit
	return f
}

func (f *Find) Skip(skip int64) *Find {
	f.option.Skip = &skip
	return f
}

func (f *Find) Projection(projection interface{}) *Find {
	f.option.Projection = projection
	return f
}

func (f *Find) Include(fields ...string) *Find {
	var projection = bson.M{}
	for i := 0; i < len(fields); i++ {
		projection[fields[i]] = 1
	}
	f.option.Projection = projection
	return f
}

func (f *Find) Exclude(fields ...string) *Find {
	var projection = bson.M{}
	for i := 0; i < len(fields); i++ {
		projection[fields[i]] = 0
	}
	f.option.Projection = projection
	return f
}

func (f *Find) Context(ctx context.Context) *Find {
	f.sessionContext = ctx
	return f
}

func (f *Find) Option(opt *options.FindOptions) *Find {
	f.option = opt
	return f
}

func (f *Find) Count(opts ...*options.CountOptions) (int64, error) {

	var t = time.Now()

	var res int64 = 0
	var err error

	defer func() {
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.Count,
			},
			Query: call.Query{
				Filter:  nil,
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

	var ref = reflect.ValueOf(f.filter)
	if ref.IsNil() || (ref.Kind() == reflect.Map && ref.Len() == 0) {
		res, err = f.collection.EstimatedDocumentCount(f.sessionContext)
		return res, err
	}

	res, err = f.collection.CountDocuments(f.sessionContext, f.filter, opts...)
	return res, err
}

func (f *Find) All(result interface{}) error {

	var t = time.Now()

	var res int64 = 0
	var err error

	defer func() {
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.Find,
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

	cursor, err := f.collection.Find(f.sessionContext, f.filter, f.option)
	var all = &MultiResult{cursor: cursor, err: err}
	err = all.All(f.sessionContext, result)
	if err != nil {
		return err
	}
	res = int64(reflect.ValueOf(result).Elem().Len())
	return nil
}

//func (f *Find) One(result interface{}) error {
//	if f.err != nil {
//		return f.err
//	}
//	cursor := f.collection.FindOne(f.sessionContext, f.filter, f.option)
//	var res = &SingleResult{singleResult: cursor}
//	return res.One(result)
//}
