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

type Aggregate struct {
	collection     *mongo.Collection
	option         *options.AggregateOptions
	pipeline       interface{}
	sessionContext context.Context
}

func NewAggregate(ctx context.Context, collection *mongo.Collection, pipeline interface{}) *Aggregate {
	return &Aggregate{collection: collection, option: &options.AggregateOptions{}, pipeline: pipeline, sessionContext: ctx}
}

func (f *Aggregate) Context(ctx context.Context) *Aggregate {
	f.sessionContext = ctx
	return f
}

func (f *Aggregate) Hit(hit interface{}) *Aggregate {
	f.option.Hint = hit
	return f
}

func (f *Aggregate) Option(opt *options.AggregateOptions) *Aggregate {
	f.option = opt
	return f
}

func (f *Aggregate) All(result interface{}) error {
	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, f.option)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.All(f.sessionContext, result)
}

func (f *Aggregate) One(result interface{}) error {
	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, f.option)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.One(f.sessionContext, result)
}
