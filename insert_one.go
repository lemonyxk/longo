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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InsertOne struct {
	collection      *mongo.Collection
	insertOneOption options.Lister[options.InsertOneOptions]
	document        interface{}
	sessionContext  context.Context
}

func NewInsertOne(ctx context.Context, collection *mongo.Collection, document interface{}) *InsertOne {
	return &InsertOne{collection: collection, insertOneOption: options.InsertOne(), document: document, sessionContext: ctx}
}

func (f *InsertOne) Option(opt options.Lister[options.InsertOneOptions]) *InsertOne {
	f.insertOneOption = opt
	return f
}

func (f *InsertOne) Context(ctx context.Context) *InsertOne {
	f.sessionContext = ctx
	return f
}

func (f *InsertOne) Exec() (*mongo.InsertOneResult, error) {
	res, err := f.collection.InsertOne(f.sessionContext, f.document, f.insertOneOption)
	if err != nil {
		return nil, err
	}
	if res.InsertedID == nil {
		return nil, fmt.Errorf("insert one error: %s", "no id")
	}
	return res, nil
}
