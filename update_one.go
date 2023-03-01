/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2020-12-14 11:50
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateOne struct {
	collection      *mongo.Collection
	updateOneOption *options.UpdateOptions
	filter          interface{}
	update          interface{}
	sessionContext  context.Context
}

func NewUpdateOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *UpdateOne {
	return &UpdateOne{collection: collection, updateOneOption: &options.UpdateOptions{}, filter: filter, update: update, sessionContext: ctx}
}

func (i *UpdateOne) Option(opt *options.UpdateOptions) *UpdateOne {
	i.updateOneOption = opt
	return i
}

func (i *UpdateOne) Context(ctx context.Context) *UpdateOne {
	i.sessionContext = ctx
	return i
}

func (i *UpdateOne) Exec() (*mongo.UpdateResult, error) {
	return i.collection.UpdateOne(i.sessionContext, i.filter, i.update, i.updateOneOption)
}
