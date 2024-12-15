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
	"github.com/lemonyxk/longo/call"
	"time"

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

func (f *ReplaceOne) Option(opt *options.ReplaceOptions) *ReplaceOne {
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

	var t = time.Now()
	var res *mongo.UpdateResult
	var err error

	defer func() {
		if res == nil {
			res = &mongo.UpdateResult{}
		}
		call.Default.Call(call.Record{
			Meta: call.Meta{
				Database:   f.collection.Database().Name(),
				Collection: f.collection.Name(),
				Type:       call.ReplaceOne,
			},
			Query: call.Query{
				Filter:  f.filter,
				Updater: f.update,
			},
			Result: call.Result{
				Insert: 0,
				Update: res.ModifiedCount,
				Delete: 0,
				Match:  res.MatchedCount,
				Upsert: res.UpsertedCount,
			},
			Consuming: time.Since(t).Microseconds(),
			Error:     err,
		})
	}()

	res, err = f.collection.ReplaceOne(f.sessionContext, f.filter, f.update, f.replaceOneOption)
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
