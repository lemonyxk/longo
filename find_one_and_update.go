/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-29 14:00
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FindOneAndUpdate struct {
	collection     *mongo.Collection
	option         *options.FindOneAndUpdateOptions
	filter         interface{}
	update         interface{}
	sessionContext context.Context
}

func NewFindOneAndUpdate(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *FindOneAndUpdate {
	return &FindOneAndUpdate{collection: collection, option: &options.FindOneAndUpdateOptions{}, filter: filter, update: update, sessionContext: ctx}
}

func (f *FindOneAndUpdate) Hit(hit interface{}) *FindOneAndUpdate {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndUpdate) Sort(sort interface{}) *FindOneAndUpdate {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndUpdate) Projection(projection interface{}) *FindOneAndUpdate {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndUpdate) Upsert() *FindOneAndUpdate {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndUpdate) ReturnDocument() *FindOneAndUpdate {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndUpdate) Context(ctx context.Context) *FindOneAndUpdate {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndUpdate) Option(opt *options.FindOneAndUpdateOptions) *FindOneAndUpdate {
	f.option = opt
	return f
}

func (f *FindOneAndUpdate) Exec(result interface{}) error {

	//var t = time.Now()
	//var res int64 = 0
	//var err error
	//
	//defer func() {
	//	call.Default.Call(call.Record{
	//		Meta: call.Meta{
	//			Database:   f.collection.Database().Name(),
	//			Collection: f.collection.Name(),
	//			Type:       call.FindOneAndUpdate,
	//		},
	//		Query: call.Query{
	//			Filter:  f.filter,
	//			Updater: f.update,
	//		},
	//		Result: call.Result{
	//			Insert: 0,
	//			Update: res,
	//			Delete: 0,
	//			Match:  res,
	//			Upsert: 0,
	//		},
	//		Consuming: time.Since(t).Microseconds(),
	//		Error:     err,
	//	})
	//}()

	var cursor = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, f.option)}
	err := cursor.Get(result)
	if err != nil {
		return err
	}
	//res = 1
	return nil
}
