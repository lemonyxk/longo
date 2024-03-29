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

func (f *Find) Context(ctx context.Context) *Find {
	f.sessionContext = ctx
	return f
}

func (f *Find) Option(opt *options.FindOptions) *Find {
	f.option = opt
	return f
}

func (f *Find) Count(opts ...*options.CountOptions) (int64, error) {
	if f.filter == nil {
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

func (f *Find) One(result interface{}) error {
	if f.err != nil {
		return f.err
	}
	cursor, err := f.collection.Find(f.sessionContext, f.filter, f.option)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.One(f.sessionContext, result)
}
