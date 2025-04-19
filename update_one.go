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

type UpdateOne struct {
	collection      *mongo.Collection
	updateOneOption options.Lister[options.UpdateOneOptions]
	filter          interface{}
	update          interface{}
	sessionContext  context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
}

func NewUpdateOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *UpdateOne {
	return &UpdateOne{collection: collection, updateOneOption: options.UpdateOne(), filter: filter, update: update, sessionContext: ctx}
}

func (f *UpdateOne) Option(opt options.Lister[options.UpdateOneOptions]) *UpdateOne {
	f.updateOneOption = opt
	return f
}

func (f *UpdateOne) Context(ctx context.Context) *UpdateOne {
	f.sessionContext = ctx
	return f
}

func (f *UpdateOne) MustModified() *UpdateOne {
	f.mustModified = true
	return f
}

func (f *UpdateOne) MustMatched() *UpdateOne {
	f.mustMatched = true
	return f
}

func (f *UpdateOne) MustUpsert() *UpdateOne {
	f.mustUpsert = true
	return f
}

func (f *UpdateOne) Exec() (*mongo.UpdateResult, error) {
	res, err := f.collection.UpdateOne(f.sessionContext, f.filter, f.update, f.updateOneOption)
	if err != nil {
		return nil, err
	}
	if f.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no modified")
		}
	}
	if f.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no matched")
		}
	}
	if f.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no upsert")
		}
	}
	return res, nil
}
