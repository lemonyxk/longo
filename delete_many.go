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

type DeleteMany struct {
	collection       *mongo.Collection
	deleteManyOption *options.DeleteOptions
	filter           interface{}
	sessionContext   context.Context
}

func NewDeleteMany(ctx context.Context, collection *mongo.Collection, filter interface{}) *DeleteMany {
	return &DeleteMany{collection: collection, deleteManyOption: &options.DeleteOptions{}, filter: filter, sessionContext: ctx}
}

func (i *DeleteMany) Option(opt *options.DeleteOptions) *DeleteMany {
	i.deleteManyOption = opt
	return i
}

func (i *DeleteMany) Context(ctx context.Context) *DeleteMany {
	i.sessionContext = ctx
	return i
}

func (i *DeleteMany) Exec() (*mongo.DeleteResult, error) {
	return i.collection.DeleteMany(i.sessionContext, i.filter, i.deleteManyOption)
}
