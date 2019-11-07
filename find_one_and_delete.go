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

type FindOneAndDelete struct {
	collection              *mongo.Collection
	findOneAndDeleteOptions options.FindOneAndDeleteOptions
	filter                  interface{}
	sessionContext          context.Context
}

func (f *FindOneAndDelete) Sort(sort interface{}) *FindOneAndDelete {
	f.findOneAndDeleteOptions.Sort = sort
	return f
}

func (f *FindOneAndDelete) Projection(projection interface{}) *FindOneAndDelete {
	f.findOneAndDeleteOptions.Projection = projection
	return f
}

func (f *FindOneAndDelete) Get(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, &f.findOneAndDeleteOptions)}
	return res.Get(result)
}
