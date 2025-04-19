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

type FindOneAndDelete struct {
	collection     *mongo.Collection
	option         *options.FindOneAndDeleteOptionsBuilder
	filter         interface{}
	sessionContext context.Context
}

func NewFindOneAndDelete(ctx context.Context, collection *mongo.Collection, filter interface{}) *FindOneAndDelete {
	return &FindOneAndDelete{collection: collection, option: options.FindOneAndDelete(), filter: filter, sessionContext: ctx}
}

func (f *FindOneAndDelete) Hit(hit interface{}) *FindOneAndDelete {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndDeleteOptions) error {
		findOptions.Hint = hit
		return nil
	})
	return f
}

func (f *FindOneAndDelete) Sort(sort interface{}) *FindOneAndDelete {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndDeleteOptions) error {
		findOptions.Sort = sort
		return nil
	})
	return f
}

func (f *FindOneAndDelete) Projection(projection interface{}) *FindOneAndDelete {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndDeleteOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *FindOneAndDelete) Context(ctx context.Context) *FindOneAndDelete {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndDelete) Option(opt *options.FindOneAndDeleteOptionsBuilder) *FindOneAndDelete {
	f.option = opt
	return f
}

func (f *FindOneAndDelete) Exec(result interface{}) error {
	var cursor = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	err := cursor.Get(result)
	if err != nil {
		return err
	}
	return err
}
