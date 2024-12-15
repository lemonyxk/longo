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

func (f *UpdateOne) Option(opt *options.UpdateOptions) *UpdateOne {
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
				Type:       call.UpdateOne,
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

	res, err = f.collection.UpdateOne(f.sessionContext, f.filter, f.update, f.updateOneOption)
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
