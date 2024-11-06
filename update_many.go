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

type UpdateMany struct {
	collection       *mongo.Collection
	updateManyOption *options.UpdateOptions
	filter           interface{}
	update           interface{}
	sessionContext   context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
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

func (i *UpdateMany) MustModified() *UpdateMany {
	i.mustModified = true
	return i
}

func (i *UpdateMany) MustMatched() *UpdateMany {
	i.mustMatched = true
	return i
}

func (i *UpdateMany) MustUpsert() *UpdateMany {
	i.mustUpsert = true
	return i
}

func (i *UpdateMany) Exec() (*mongo.UpdateResult, error) {
	var res, err = i.collection.UpdateMany(i.sessionContext, i.filter, i.update, i.updateManyOption)
	if err != nil {
		return nil, err
	}
	if i.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no modified")
		}
	}
	if i.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no matched")
		}
	}
	if i.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no upsert")
		}
	}
	return res, nil
}
