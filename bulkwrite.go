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

type BulkWrite struct {
	collection       *mongo.Collection
	bulkWriteOptions options.Lister[options.BulkWriteOptions]
	models           []mongo.WriteModel
	sessionContext   context.Context

	mustModified bool
	mustMatched  bool
	mustUpsert   bool
	mustDeleted  bool
	mustInserted bool
}

func NewBulkWrite(ctx context.Context, collection *mongo.Collection, models []mongo.WriteModel) *BulkWrite {
	return &BulkWrite{collection: collection, bulkWriteOptions: options.BulkWrite(), models: models, sessionContext: ctx}
}

func (f *BulkWrite) Option(opt options.Lister[options.BulkWriteOptions]) *BulkWrite {
	f.bulkWriteOptions = opt
	return f
}

func (f *BulkWrite) Context(ctx context.Context) *BulkWrite {
	f.sessionContext = ctx
	return f
}

func (f *BulkWrite) MustModified() *BulkWrite {
	f.mustModified = true
	return f
}

func (f *BulkWrite) MustMatched() *BulkWrite {
	f.mustMatched = true
	return f
}

func (f *BulkWrite) MustUpsert() *BulkWrite {
	f.mustUpsert = true
	return f
}

func (f *BulkWrite) MustDeleted() *BulkWrite {
	f.mustDeleted = true
	return f
}

func (f *BulkWrite) MustInserted() *BulkWrite {
	f.mustInserted = true
	return f
}

func (f *BulkWrite) Exec() (*mongo.BulkWriteResult, error) {
	res, err := f.collection.BulkWrite(f.sessionContext, f.models, f.bulkWriteOptions)
	if err != nil {
		return nil, err
	}
	if f.mustModified {
		if res.ModifiedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no modified")
		}
	}
	if f.mustMatched {
		if res.MatchedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no matched")
		}
	}
	if f.mustUpsert {
		if res.UpsertedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no upsert")
		}
	}
	if f.mustDeleted {
		if res.DeletedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no deleted")
		}
	}
	if f.mustInserted {
		if res.InsertedCount == 0 {
			return nil, fmt.Errorf("update bulkwrite error: %s", "no inserted")
		}
	}
	return res, nil
}
