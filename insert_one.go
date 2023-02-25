/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2020-12-14 11:50
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InsertOne struct {
	collection      *mongo.Collection
	insertOneOption *options.InsertOneOptions
	document        interface{}
	sessionContext  context.Context
}

func NewInsertOne(ctx context.Context, collection *mongo.Collection, document interface{}) *InsertOne {
	return &InsertOne{collection: collection, insertOneOption: &options.InsertOneOptions{}, document: document, sessionContext: ctx}
}

func (i *InsertOne) Option(opt *options.InsertOneOptions) *InsertOne {
	i.insertOneOption = opt
	return i
}

func (i *InsertOne) Context(ctx context.Context) *InsertOne {
	i.sessionContext = ctx
	return i
}

func (i *InsertOne) Exec() (*mongo.InsertOneResult, error) {
	return i.collection.InsertOne(i.sessionContext, i.document, i.insertOneOption)
}
