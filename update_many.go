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

type UpdateMany struct {
	collection       *mongo.Collection
	updateManyOption *options.UpdateOptions
	filter           interface{}
	update           interface{}
	sessionContext   context.Context
}

func NewUpdateMany(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *UpdateMany {
	return &UpdateMany{collection: collection, updateManyOption: &options.UpdateOptions{}, filter: filter, update: update, sessionContext: ctx}
}

func (i *UpdateMany) Option(opt *options.UpdateOptions) *UpdateMany {
	i.updateManyOption = opt
	return i
}

func (i *UpdateMany) Context(ctx context.Context) *UpdateMany {
	i.sessionContext = ctx
	return i
}

func (i *UpdateMany) Exec() (*mongo.UpdateResult, error) {
	return i.collection.UpdateMany(i.sessionContext, i.filter, i.update, i.updateManyOption)
}
