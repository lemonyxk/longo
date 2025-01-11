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

func (f *UpdateMany) Option(opt *options.UpdateOptions) *UpdateMany {
	f.updateManyOption = opt
	return f
}

func (f *UpdateMany) Context(ctx context.Context) *UpdateMany {
	f.sessionContext = ctx
	return f
}

func (f *UpdateMany) MustModified() *UpdateMany {
	f.mustModified = true
	return f
}

func (f *UpdateMany) MustMatched() *UpdateMany {
	f.mustMatched = true
	return f
}

func (f *UpdateMany) MustUpsert() *UpdateMany {
	f.mustUpsert = true
	return f
}

func (f *UpdateMany) Exec() (*mongo.UpdateResult, error) {

	//var t = time.Now()
	//var res *mongo.UpdateResult
	//var err error
	//
	//defer func() {
	//	if res == nil {
	//		res = &mongo.UpdateResult{}
	//	}
	//	call.Default.Call(call.Record{
	//		Meta: call.Meta{
	//			Database:   f.collection.Database().Name(),
	//			Collection: f.collection.Name(),
	//			Type:       call.UpdateMany,
	//		},
	//		Query: call.Query{
	//			Filter:  f.filter,
	//			Updater: f.update,
	//		},
	//		Result: call.Result{
	//			Insert: 0,
	//			Update: res.ModifiedCount,
	//			Delete: 0,
	//			Match:  res.MatchedCount,
	//			Upsert: res.UpsertedCount,
	//		},
	//		Consuming: time.Since(t).Microseconds(),
	//		Error:     err,
	//	})
	//}()

	res, err := f.collection.UpdateMany(f.sessionContext, f.filter, f.update, f.updateManyOption)
	if err != nil {
		return nil, err
	}
	if f.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no modified")
		}
	}
	if f.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no matched")
		}
	}
	if f.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update many error: %s", "no upsert")
		}
	}
	return res, nil
}
