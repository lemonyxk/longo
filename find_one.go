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
	option         *options.FindOneOptions
	filter         interface{}
	sessionContext context.Context
	err            error
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

func (f *FindOne) Exec(result interface{}) error {
	if f.err != nil {
		return f.err
	}
	var res = &SingleResult{singleResult: f.collection.FindOne(f.sessionContext, f.filter, f.option)}
	return res.One(result)
}
