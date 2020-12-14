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

type InsertMany struct {
	collection       *mongo.Collection
	insertManyOption *options.InsertManyOptions
	document         []interface{}
	sessionContext   context.Context
}

func NewInsertMany(collection *mongo.Collection, document []interface{}) *InsertMany {
	return &InsertMany{collection: collection, insertManyOption: &options.InsertManyOptions{}, document: document}
}

func (i *InsertMany) Option(opt *options.InsertManyOptions) *InsertMany {
	i.insertManyOption = opt
	return i
}

func (i *InsertMany) Context(ctx context.Context) *InsertMany {
	i.sessionContext = ctx
	return i
}

func (i *InsertMany) Do() *InsertManyResult {
	res, err := i.collection.InsertMany(i.sessionContext, i.document, i.insertManyOption)
	return &InsertManyResult{insertManyResult: res, err: err}
}
