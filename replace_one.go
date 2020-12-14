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

type ReplaceOne struct {
	collection       *mongo.Collection
	replaceOneOption *options.ReplaceOptions
	filter           interface{}
	update           interface{}
	sessionContext   context.Context
}

func NewReplaceOne(collection *mongo.Collection, filter interface{}, update interface{}) *ReplaceOne {
	return &ReplaceOne{collection: collection, replaceOneOption: &options.ReplaceOptions{}, filter: filter, update: update}
}

func (i *ReplaceOne) Option(opt *options.ReplaceOptions) *ReplaceOne {
	i.replaceOneOption = opt
	return i
}

func (i *ReplaceOne) Context(ctx context.Context) *ReplaceOne {
	i.sessionContext = ctx
	return i
}

func (i *ReplaceOne) Do() *UpdateResult {
	res, err := i.collection.ReplaceOne(i.sessionContext, i.filter, i.update, i.replaceOneOption)
	return &UpdateResult{updateResult: res, err: err}
}
