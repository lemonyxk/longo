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

type FindOneAndReplace struct {
	collection     *mongo.Collection
	option         *options.FindOneAndReplaceOptionsBuilder
	filter         interface{}
	replacement    interface{}
	sessionContext context.Context
}

func NewFindOneAndReplace(ctx context.Context, collection *mongo.Collection, filter interface{}, replacement interface{}) *FindOneAndReplace {
	return &FindOneAndReplace{collection: collection, option: options.FindOneAndReplace(), filter: filter, replacement: replacement, sessionContext: ctx}
}

func (f *FindOneAndReplace) Hit(hit interface{}) *FindOneAndReplace {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndReplaceOptions) error {
		findOptions.Hint = hit
		return nil
	})
	return f
}

func (f *FindOneAndReplace) Sort(sort interface{}) *FindOneAndReplace {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndReplaceOptions) error {
		findOptions.Sort = sort
		return nil
	})
	return f
}

func (f *FindOneAndReplace) Projection(projection interface{}) *FindOneAndReplace {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndReplaceOptions) error {
		findOptions.Projection = projection
		return nil
	})
	return f
}

func (f *FindOneAndReplace) Upsert() *FindOneAndReplace {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndReplaceOptions) error {
		var t = true
		findOptions.Upsert = &t
		return nil
	})
	return f
}

func (f *FindOneAndReplace) ReturnDocument() *FindOneAndReplace {
	f.option.Opts = append(f.option.Opts, func(findOptions *options.FindOneAndReplaceOptions) error {
		var t = options.After
		findOptions.ReturnDocument = &t
		return nil
	})
	return f
}

func (f *FindOneAndReplace) Context(ctx context.Context) *FindOneAndReplace {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndReplace) Option(opt *options.FindOneAndReplaceOptionsBuilder) *FindOneAndReplace {
	f.option = opt
	return f
}

func (f *FindOneAndReplace) Exec(result interface{}) error {
	var cursor = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.option)}
	err := cursor.Get(result)
	if err != nil {
		return err
	}
	return nil
}
