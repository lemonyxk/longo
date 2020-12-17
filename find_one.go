/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-29 14:00
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOne struct {
	collection     *mongo.Collection
	findOptions    *options.FindOneOptions
	filter         interface{}
	sessionContext context.Context
}

func NewFindOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *FindOne {
	return &FindOne{collection: collection, findOptions: &options.FindOneOptions{}, filter: filter, sessionContext: ctx}
}

func (f *FindOne) Sort(sort interface{}) *FindOne {
	f.findOptions.Sort = sort
	return f
}

func (f *FindOne) Skip(skip int64) *FindOne {
	f.findOptions.Skip = &skip
	return f
}

func (f *FindOne) Projection(projection interface{}) *FindOne {
	f.findOptions.Projection = projection
	return f
}

func (f *FindOne) Context(ctx context.Context) *FindOne {
	f.sessionContext = ctx
	return f
}

func (f *FindOne) Option(opt *options.FindOneOptions) *FindOne {
	f.findOptions = opt
	return f
}

func (f *FindOne) Do(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOne(f.sessionContext, f.filter, f.findOptions)}
	return res.Do(result)
}
