/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:33
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IndexView struct {
	view mongo.IndexView
}

func (iv *IndexView) List(opts ...*options.ListIndexesOptions) *MultiResult {
	cursor, err := iv.view.List(context.Background(), opts...)
	return &MultiResult{cursor: cursor, err: err}
}

func (iv *IndexView) DropAll(opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return iv.view.DropAll(context.Background(), opts...)
}

func (iv *IndexView) DropOne(name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return iv.view.DropOne(context.Background(), name, opts...)
}

func (iv *IndexView) CreateMany(models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return iv.view.CreateMany(context.Background(), models, opts...)
}

func (iv *IndexView) CreateOne(model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return iv.view.CreateOne(context.Background(), model, opts...)
}
