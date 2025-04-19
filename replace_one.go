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

type ReplaceOne struct {
	collection       *mongo.Collection
	replaceOneOption options.Lister[options.ReplaceOptions]
	filter           interface{}
	update           interface{}
	sessionContext   context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
}

func NewReplaceOne(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *ReplaceOne {
	return &ReplaceOne{collection: collection, replaceOneOption: options.Replace(), filter: filter, update: update, sessionContext: ctx}
}

func (f *ReplaceOne) Option(opt options.Lister[options.ReplaceOptions]) *ReplaceOne {
	f.replaceOneOption = opt
	return f
}

func (f *ReplaceOne) Context(ctx context.Context) *ReplaceOne {
	f.sessionContext = ctx
	return f
}

func (f *ReplaceOne) MustModified() *ReplaceOne {
	f.mustModified = true
	return f
}

func (f *ReplaceOne) MustMatched() *ReplaceOne {
	f.mustMatched = true
	return f
}

func (f *ReplaceOne) MustUpsert() *ReplaceOne {
	f.mustUpsert = true
	return f
}

func (f *ReplaceOne) Exec() (*mongo.UpdateResult, error) {
	res, err := f.collection.ReplaceOne(f.sessionContext, f.filter, f.update, f.replaceOneOption)
	if err != nil {
		return nil, err
	}
	if f.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no modified")
		}
	}
	if f.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no matched")
		}
	}
	if f.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update replace one error: %s", "no upsert")
		}
	}
	return res, nil
}
