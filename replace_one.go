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

type ReplaceOne struct {
	collection       *mongo.Collection
	replaceOneOption *options.ReplaceOptions
	filter           interface{}
	update           interface{}
	sessionContext   context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
}

func NewReplaceOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *ReplaceOne {
	return &ReplaceOne{collection: collection, replaceOneOption: &options.ReplaceOptions{}, filter: filter, update: update, sessionContext: ctx}
}

func (i *ReplaceOne) Option(opt *options.ReplaceOptions) *ReplaceOne {
	i.replaceOneOption = opt
	return i
}

func (i *ReplaceOne) Context(ctx context.Context) *ReplaceOne {
	i.sessionContext = ctx
	return i
}

func (i *ReplaceOne) MustModified() *ReplaceOne {
	i.mustModified = true
	return i
}

func (i *ReplaceOne) MustMatched() *ReplaceOne {
	i.mustMatched = true
	return i
}

func (i *ReplaceOne) MustUpsert() *ReplaceOne {
	i.mustUpsert = true
	return i
}

func (i *ReplaceOne) Exec() (*mongo.UpdateResult, error) {
	var res, err = i.collection.ReplaceOne(i.sessionContext, i.filter, i.update, i.replaceOneOption)
	if err != nil {
		return nil, err
	}
	if i.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no modified")
		}
	}
	if i.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no matched")
		}
	}
	if i.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no upsert")
		}
	}
	return res, nil
}
