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

type Find struct {
	collection  *mongo.Collection
	findOptions options.FindOptions
	filter      interface{}
}

func (f *Find) Sort(sort interface{}) *Find {
	f.findOptions.Sort = sort
	return f
}

func (f *Find) Limit(limit int64) *Find {
	f.findOptions.Limit = &limit
	return f
}

func (f *Find) Skip(skip int64) *Find {
	f.findOptions.Skip = &skip
	return f
}

func (f *Find) Projection(projection interface{}) *Find {
	f.findOptions.Projection = projection
	return f
}

func (f *Find) All(result interface{}) error {
	cursor, err := f.collection.Find(context.Background(), f.filter, &f.findOptions)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.All(result)
}
