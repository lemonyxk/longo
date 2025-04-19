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

type FindOne struct {
	collection     *mongo.Collection
	option         *options.FindOneOptionsBuilder
	filter         interface{}
	sessionContext context.Context
}

func NewFindOne(ctx context.Context, collection *mongo.Collection, filter interface{}) *FindOne {
	return &FindOne{collection: collection, option: options.FindOne(), filter: filter, sessionContext: ctx}
}

func (f *FindOne) Hit(hit interface{}) *FindOne {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneOptions) error {
		findOptions.Hint = hit
		return nil
	})
	return f
}

func (f *FindOne) Sort(sort interface{}) *FindOne {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneOptions) error {
		findOptions.Sort = sort
		return nil
	})
	return f
}

func (f *FindOne) Skip(skip int64) *FindOne {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneOptions) error {
		findOptions.Skip = &skip
		return nil
	})
	return f
}

func (f *FindOne) Projection(projection interface{}) *FindOne {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *FindOne) Context(ctx context.Context) *FindOne {
	f.sessionContext = ctx
	return f
}

func (f *FindOne) Option(opt *options.FindOneOptionsBuilder) *FindOne {
	f.option = opt
	return f
}

func (f *FindOne) Get(result interface{}) error {
	var cursor = &SingleResult{singleResult: f.collection.FindOne(f.sessionContext, f.filter, f.option)}
	err := cursor.Get(result)
	if err == mongo.ErrNoDocuments { // only for FindOne, if modify action, it will return ErrNoDocuments
		err = nil
	}
	if err != nil {
		return err
	}
	return nil
}
