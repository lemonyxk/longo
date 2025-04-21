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

type InsertMany struct {
	collection       *mongo.Collection
	insertManyOption options.Lister[options.InsertManyOptions]
	document         interface{}
	sessionContext   context.Context
}

func NewInsertMany(ctx context.Context, collection *mongo.Collection, document interface{}) *InsertMany {
	return &InsertMany{collection: collection, insertManyOption: options.InsertMany(), document: document, sessionContext: ctx}
}

func (f *InsertMany) Option(opt options.Lister[options.InsertManyOptions]) *InsertMany {
	f.insertManyOption = opt
	return f
}

func (f *InsertMany) Context(ctx context.Context) *InsertMany {
	f.sessionContext = ctx
	return f
}

func (f *InsertMany) Exec() (*mongo.InsertManyResult, error) {
	res, err := f.collection.InsertMany(f.sessionContext, f.document, f.insertManyOption)
	if err != nil {
		return nil, err
	}
	if len(res.InsertedIDs) == 0 {
		return nil, fmt.Errorf("insert many error: %s", "no id")
	}
	return res, nil
}
