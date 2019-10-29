/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-29 14:00
**/

package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOne struct {
	collection  *mongo.Collection
	findOptions options.FindOneOptions
	filter      interface{}
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

func (f *FindOne) All(result interface{}) error {
	var res = &SingleResult{singleResult: f.collection.FindOne(context.Background(), f.filter, &f.findOptions)}
	return res.Get(result)
}
