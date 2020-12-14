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
	collection               *mongo.Collection
	findOneAndReplaceOptions *options.FindOneAndReplaceOptions
	filter                   interface{}
	replacement              interface{}
	sessionContext           context.Context
}

func NewFindOneAndReplace(collection *mongo.Collection, filter interface{}, replacement interface{}) *FindOneAndReplace {
	return &FindOneAndReplace{collection: collection, findOneAndReplaceOptions: &options.FindOneAndReplaceOptions{}, filter: filter, replacement: replacement}
}

func (f *FindOneAndReplace) Sort(sort interface{}) *FindOneAndReplace {
	f.findOneAndReplaceOptions.Sort = sort
	return f
}

func (f *FindOneAndReplace) Projection(projection interface{}) *FindOneAndReplace {
	f.findOneAndReplaceOptions.Projection = projection
	return f
}

func (f *FindOneAndReplace) Upsert() *FindOneAndReplace {
	var t = true
	f.findOneAndReplaceOptions.Upsert = &t
	return f
}

func (f *FindOneAndReplace) ReturnDocument() *FindOneAndReplace {
	var t = options.After
	f.findOneAndReplaceOptions.ReturnDocument = &t
	return f
}

func (f *FindOneAndReplace) Context(ctx context.Context) *FindOneAndReplace {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndReplace) Option(opt *options.FindOneAndReplaceOptions) *FindOneAndReplace {
	f.findOneAndReplaceOptions = opt
	return f
}

func (f *FindOneAndReplace) Get(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.findOneAndReplaceOptions)}
	return res.Get(result)
}
