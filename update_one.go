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

type UpdateOne struct {
	collection      *mongo.Collection
	updateOneOption *options.UpdateOptions
	filter          interface{}
	update          interface{}
	sessionContext  context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
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

func (i *UpdateOne) MustModified() *UpdateOne {
	i.mustModified = true
	return i
}

func (i *UpdateOne) MustMatched() *UpdateOne {
	i.mustMatched = true
	return i
}

func (i *UpdateOne) MustUpsert() *UpdateOne {
	i.mustUpsert = true
	return i
}

func (i *UpdateOne) Exec() (*mongo.UpdateResult, error) {
	var res, err = i.collection.UpdateOne(i.sessionContext, i.filter, i.update, i.updateOneOption)
	if err != nil {
		return nil, err
	}
	if i.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no modified")
		}
	}
	if i.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no matched")
		}
	}
	if i.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update one error: %s", "no upsert")
		}
	}
	return res, nil
}
