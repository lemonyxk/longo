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
	var res = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	return res.One(result)
}
