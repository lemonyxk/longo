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

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BulkWrite struct {
	collection       *mongo.Collection
	bulkWriteOptions *options.BulkWriteOptions
	models           []mongo.WriteModel
	sessionContext   context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
	mustDeleted  bool
	mustInserted bool
}

func NewBulkWrite(ctx context.Context, collection *mongo.Collection, models []mongo.WriteModel) *BulkWrite {
	return &BulkWrite{collection: collection, bulkWriteOptions: &options.BulkWriteOptions{}, models: models, sessionContext: ctx}
}

func (i *BulkWrite) Option(opt *options.BulkWriteOptions) *BulkWrite {
	i.bulkWriteOptions = opt
	return i
}

func (i *BulkWrite) Context(ctx context.Context) *BulkWrite {
	i.sessionContext = ctx
	return i
}

func (i *BulkWrite) MustModified() *BulkWrite {
	i.mustModified = true
	return i
}

func (i *BulkWrite) MustMatched() *BulkWrite {
	i.mustMatched = true
	return i
}

func (i *BulkWrite) MustUpsert() *BulkWrite {
	i.mustUpsert = true
	return i
}

func (i *BulkWrite) MustDeleted() *BulkWrite {
	i.mustDeleted = true
	return i
}

func (i *BulkWrite) MustInserted() *BulkWrite {
	i.mustInserted = true
	return i
}

func (i *BulkWrite) Exec() (*mongo.BulkWriteResult, error) {
	var res, err = i.collection.BulkWrite(i.sessionContext, i.models, i.bulkWriteOptions)
	if err != nil {
		return nil, err
	}
	if i.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no modified")
		}
	}
	if i.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no matched")
		}
	}
	if i.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no upsert")
		}
	}
	if i.mustDeleted {
		if res.DeletedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no deleted")
		}
	}
	if i.mustInserted {
		if res.InsertedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no inserted")
		}
	}
	return res, nil
}
