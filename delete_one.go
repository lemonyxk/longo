/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2020-12-14 11:50
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteOne struct {
	collection      *mongo.Collection
	deleteOneOption *options.DeleteOptions
	filter          interface{}
	sessionContext  context.Context
}

func NewDeleteOne(collection *mongo.Collection, filter interface{}) *DeleteOne {
	return &DeleteOne{collection: collection, deleteOneOption: &options.DeleteOptions{}, filter: filter}
}

func (i *DeleteOne) Option(opt *options.DeleteOptions) *DeleteOne {
	i.deleteOneOption = opt
	return i
}

func (i *DeleteOne) Context(ctx context.Context) *DeleteOne {
	i.sessionContext = ctx
	return i
}

func (i *DeleteOne) Do() *DeleteResult {
	res, err := i.collection.DeleteOne(i.sessionContext, i.filter, i.deleteOneOption)
	return &DeleteResult{deleteResult: res, err: err}
}
