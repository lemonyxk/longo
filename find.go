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
	"go.mongodb.org/mongo-driver/bson"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Find struct {
	collection     *mongo.Collection
	option         *options.FindOptions
	filter         interface{}
	sessionContext context.Context
	err            error
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
	var ref = reflect.ValueOf(f.filter)
	if ref.IsNil() || (ref.Kind() == reflect.Map && ref.Len() == 0) {
		return f.collection.EstimatedDocumentCount(f.sessionContext)
	}
	return f.collection.CountDocuments(f.sessionContext, f.filter, opts...)
}

func (f *Find) All(result interface{}) error {
	if f.err != nil {
		return f.err
	}
	cursor, err := f.collection.Find(f.sessionContext, f.filter, f.option)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.All(f.sessionContext, result)
}

//func (f *Find) One(result interface{}) error {
//	if f.err != nil {
//		return f.err
//	}
//	cursor := f.collection.FindOne(f.sessionContext, f.filter, f.option)
//	var res = &SingleResult{singleResult: cursor}
//	return res.One(result)
//}
