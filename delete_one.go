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
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeleteOne struct {
	collection      *mongo.Collection
	deleteOneOption *options.DeleteOptions
	filter          interface{}
	sessionContext  context.Context

	mustDeleted bool
}

func NewDeleteOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *DeleteOne {
	return &DeleteOne{collection: collection, deleteOneOption: &options.DeleteOptions{}, filter: filter, sessionContext: ctx}
}

func (i *DeleteOne) Option(opt *options.DeleteOptions) *DeleteOne {
	i.deleteOneOption = opt
	return i
}

func (i *DeleteOne) Context(ctx context.Context) *DeleteOne {
	i.sessionContext = ctx
	return i
}

func (i *DeleteOne) MustDeleted() *DeleteOne {
	i.mustDeleted = true
	return i
}

func (i *DeleteOne) Exec() (*mongo.DeleteResult, error) {
	var res, err = i.collection.DeleteOne(i.sessionContext, i.filter, i.deleteOneOption)
	if err != nil {
		return nil, err
	}
	if i.mustDeleted {
		if res.DeletedCount == 0 {
			return nil, fmt.Errorf("no document deleted")
		}
	}
	return res, nil
}
