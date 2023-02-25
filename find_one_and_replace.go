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

type FindOneAndReplace struct {
	collection     *mongo.Collection
	option         *options.FindOneAndReplaceOptions
	filter         interface{}
	replacement    interface{}
	sessionContext context.Context
}

func NewFindOneAndReplace(ctx context.Context, collection *mongo.Collection, filter interface{}, replacement interface{}) *FindOneAndReplace {
	return &FindOneAndReplace{collection: collection, option: &options.FindOneAndReplaceOptions{}, filter: filter, replacement: replacement, sessionContext: ctx}
}

func (f *FindOneAndReplace) Hit(hit interface{}) *FindOneAndReplace {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndReplace) Sort(sort interface{}) *FindOneAndReplace {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndReplace) Projection(projection interface{}) *FindOneAndReplace {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndReplace) Upsert() *FindOneAndReplace {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndReplace) ReturnDocument() *FindOneAndReplace {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndReplace) Context(ctx context.Context) *FindOneAndReplace {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndReplace) Option(opt *options.FindOneAndReplaceOptions) *FindOneAndReplace {
	f.option = opt
	return f
}

func (f *FindOneAndReplace) Exec(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.option)}
	return res.One(result)
}
