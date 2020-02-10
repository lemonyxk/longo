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

type FindOneAndUpdate struct {
	collection              *mongo.Collection
	findOneAndUpdateOptions options.FindOneAndUpdateOptions
	filter                  interface{}
	update                  interface{}
	sessionContext          context.Context
}

func (f *FindOneAndUpdate) Sort(sort interface{}) *FindOneAndUpdate {
	f.findOneAndUpdateOptions.Sort = sort
	return f
}

func (f *FindOneAndUpdate) Projection(projection interface{}) *FindOneAndUpdate {
	f.findOneAndUpdateOptions.Projection = projection
	return f
}

func (f *FindOneAndUpdate) Upsert() *FindOneAndUpdate {
	var t = true
	f.findOneAndUpdateOptions.Upsert = &t
	return f
}

func (f *FindOneAndUpdate) ReturnDocument() *FindOneAndUpdate {
	var t = options.After
	f.findOneAndUpdateOptions.ReturnDocument = &t
	return f
}

func (f *FindOneAndUpdate) Get(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, &f.findOneAndUpdateOptions)}
	return res.Get(result)
}
