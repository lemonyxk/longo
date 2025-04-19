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
	"go.mongodb.org/mongo-driver/v2/mongo"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type FindOneAndUpdate struct {
	collection     *mongo.Collection
	option         *options.FindOneAndUpdateOptionsBuilder
	filter         interface{}
	update         interface{}
	sessionContext context.Context
}

func NewFindOneAndUpdate(ctx context.Context, collection *mongo.Collection, filter interface{}, update interface{}) *FindOneAndUpdate {
	return &FindOneAndUpdate{collection: collection, option: options.FindOneAndUpdate(), filter: filter, update: update, sessionContext: ctx}
}

func (f *FindOneAndUpdate) Hit(hit interface{}) *FindOneAndUpdate {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndUpdateOptions) error {
		findOptions.Hint = hit
		return nil
	})
	return f
}

func (f *FindOneAndUpdate) Sort(sort interface{}) *FindOneAndUpdate {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndUpdateOptions) error {
		findOptions.Sort = sort
		return nil
	})
	return f
}

func (f *FindOneAndUpdate) Projection(projection interface{}) *FindOneAndUpdate {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndUpdateOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *FindOneAndUpdate) Upsert() *FindOneAndUpdate {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndUpdateOptions) error {
		var t = true
		findOptions.Upsert = &t
		return nil
	})
	return f
}

func (f *FindOneAndUpdate) ReturnDocument() *FindOneAndUpdate {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndUpdateOptions) error {
		var t = options.After
		findOptions.ReturnDocument = &t
		return nil
	})
	return f
}

func (f *FindOneAndUpdate) Context(ctx context.Context) *FindOneAndUpdate {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndUpdate) Option(opt *options.FindOneAndUpdateOptionsBuilder) *FindOneAndUpdate {
	f.option = opt
	return f
}

func (f *FindOneAndUpdate) Exec(result interface{}) error {
	var cursor = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, f.option)}
	err := cursor.Get(result)
	if err != nil {
		return err
	}
	return nil
}
