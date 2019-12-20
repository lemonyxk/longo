/**
* @program: mongo
*
* @description:
*
* @author: lemo
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
	collection       *mongo.Collection
	aggregateOptions options.AggregateOptions
	pipeline         interface{}
	sessionContext   context.Context
}

func (f *Aggregate) All(result interface{}) error {
	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, &f.aggregateOptions)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.All(f.sessionContext, result)
}

func (f *Aggregate) One(result interface{}) error {
	cursor, err := f.collection.Aggregate(f.sessionContext, f.pipeline, &f.aggregateOptions)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.One(f.sessionContext, result)
}
