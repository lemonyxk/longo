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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DeleteOne struct {
	collection      *mongo.Collection
	deleteOneOption options.Lister[options.DeleteOneOptions]
	filter          interface{}
	sessionContext  context.Context

	mustDeleted bool
}

func NewDeleteOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *DeleteOne {
	return &DeleteOne{collection: collection, deleteOneOption: options.DeleteOne(), filter: filter, sessionContext: ctx}
}

func (f *DeleteOne) Option(opt options.Lister[options.DeleteOneOptions]) *DeleteOne {
	f.deleteOneOption = opt
	return f
}

func (f *DeleteOne) Context(ctx context.Context) *DeleteOne {
	f.sessionContext = ctx
	return f
}

func (f *DeleteOne) MustDeleted() *DeleteOne {
	f.mustDeleted = true
	return f
}

func (f *DeleteOne) Exec() (*mongo.DeleteResult, error) {

	res, err := f.collection.DeleteOne(f.sessionContext, f.filter, f.deleteOneOption)
	if err != nil {
		return nil, err
	}

	if f.mustDeleted {
		if res.DeletedCount == 0 {
			return nil, fmt.Errorf("no document deleted")
		}
	}

	return res, nil
}
