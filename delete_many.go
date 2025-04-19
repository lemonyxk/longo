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

type DeleteMany struct {
	collection       *mongo.Collection
	deleteManyOption options.Lister[options.DeleteManyOptions]
	filter           interface{}
	sessionContext   context.Context

	mustDeleted bool
}

func NewDeleteMany(ctx context.Context, collection *mongo.Collection, filter interface{}) *DeleteMany {
	return &DeleteMany{collection: collection, deleteManyOption: options.DeleteMany(), filter: filter, sessionContext: ctx}
}

func (f *DeleteMany) Option(opt options.Lister[options.DeleteManyOptions]) *DeleteMany {
	f.deleteManyOption = opt
	return f
}

func (f *DeleteMany) Context(ctx context.Context) *DeleteMany {
	f.sessionContext = ctx
	return f
}

func (f *DeleteMany) MustDeleted() *DeleteMany {
	f.mustDeleted = true
	return f
}

func (f *DeleteMany) Exec() (*mongo.DeleteResult, error) {

	res, err := f.collection.DeleteMany(f.sessionContext, f.filter, f.deleteManyOption)
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
