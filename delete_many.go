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

type DeleteMany struct {
	collection       *mongo.Collection
	deleteManyOption *options.DeleteOptions
	filter           interface{}
	sessionContext   context.Context

	mustDeleted bool
}

func NewDeleteMany(ctx context.Context, collection *mongo.Collection, filter interface{}) *DeleteMany {
	return &DeleteMany{collection: collection, deleteManyOption: &options.DeleteOptions{}, filter: filter, sessionContext: ctx}
}

func (f *DeleteMany) Option(opt *options.DeleteOptions) *DeleteMany {
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

	//var t = time.Now()
	//var res *mongo.DeleteResult
	//var err error
	//
	//defer func() {
	//	if res == nil {
	//		res = &mongo.DeleteResult{}
	//	}
	//	call.Default.Call(call.Record{
	//		Meta: call.Meta{
	//			Database:   f.collection.Database().Name(),
	//			Collection: f.collection.Name(),
	//			Type:       call.DeleteMany,
	//		},
	//		Query: call.Query{
	//			Filter:  f.filter,
	//			Updater: nil,
	//		},
	//		Result: call.Result{
	//			Insert: 0,
	//			Update: 0,
	//			Delete: res.DeletedCount,
	//			Match:  0,
	//			Upsert: 0,
	//		},
	//		Consuming: time.Since(t).Microseconds(),
	//		Error:     err,
	//	})
	//}()

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
