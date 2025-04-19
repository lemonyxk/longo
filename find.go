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
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"reflect"
)

type Find struct {
	collection     *mongo.Collection
	option         *options.FindOptionsBuilder
	filter         interface{}
	sessionContext context.Context
}

func NewFind(ctx context.Context, collection *mongo.Collection, filter interface{}) *Find {
	return &Find{collection: collection, option: options.Find(), filter: filter, sessionContext: ctx}
}

func (f *Find) Hit(hit interface{}) *Find {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Hint = hit
		return nil
	})
	return f
}

func (f *Find) Sort(sort interface{}) *Find {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Sort = sort
		return nil
	})
	return f
}

func (f *Find) Limit(limit int64) *Find {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Limit = &limit
		return nil
	})
	return f
}

func (f *Find) Skip(skip int64) *Find {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Skip = &skip
		return nil
	})
	return f
}

func (f *Find) Projection(projection interface{}) *Find {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *Find) Include(fields ...string) *Find {
	var projection = bson.M{}
	for i := 0; i < len(fields); i++ {
		projection[fields[i]] = 1
	}
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *Find) Exclude(fields ...string) *Find {
	var projection = bson.M{}
	for i := 0; i < len(fields); i++ {
		projection[fields[i]] = 0
	}
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *Find) Context(ctx context.Context) *Find {
	f.sessionContext = ctx
	return f
}

func (f *Find) Option(opt *options.FindOptionsBuilder) *Find {
	f.option = opt
	return f
}

func (f *Find) Count(opts ...options.Lister[options.CountOptions]) (int64, error) {
	var ref = reflect.ValueOf(f.filter)
	if ref.IsNil() || (ref.Kind() == reflect.Map && ref.Len() == 0) {
		res, err := f.collection.EstimatedDocumentCount(f.sessionContext)
		return res, err
	}

	res, err := f.collection.CountDocuments(f.sessionContext, f.filter, opts...)
	return res, err
}

func (f *Find) All(result interface{}) error {
	cursor, err := f.collection.Find(f.sessionContext, f.filter, f.option)
	var all = &MultiResult{cursor: cursor, err: err}
	err = all.All(f.sessionContext, result)
	if err != nil {
		return err
	}
	return nil
}
