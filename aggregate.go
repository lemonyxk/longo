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

type Aggregate struct {
	collection     *mongo.Collection
	option         options.Lister[options.AggregateOptions]
	pipeline       interface{}
	sessionContext context.Context
}

func NewAggregate(ctx context.Context, collection *mongo.Collection, pipeline interface{}) *Aggregate {
	return &Aggregate{collection: collection, option: options.Aggregate(), pipeline: pipeline, sessionContext: ctx}
}

func (f *Aggregate) Context(ctx context.Context) *Aggregate {
	f.sessionContext = ctx
	return f
}

func (f *Aggregate) Option(opt options.Lister[options.AggregateOptions]) *Aggregate {
	f.option = opt
	return f
}

func (f *Aggregate) All(result interface{}) error {
	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, f.option)
	var all = &MultiResult{cursor: cursor, err: err}
	err = all.All(f.sessionContext, result)
	if err != nil {
		return err
	}
	return nil
}
