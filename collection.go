/**
* @program: lottery-server
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-25 15:55
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadPreference
//
// Primary
// PrimaryPreferred
// Secondary
// SecondaryPreferred
// Nearest

// ReadConcern
//
// Local
// Majority
// Linearizable
// Available
// Snapshot

// WriteConcern
//
// w:0|1|majority
// j:false|true
// wTimeout

type Collection struct {
	client            *mongo.Client
	config            Config
	db                string
	collection        string
	context           context.Context
	collectionOptions []*options.CollectionOptions
	databaseOptions   []*options.DatabaseOptions
}

func (q *Collection) Clone() (*mongo.Collection, error) {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Clone()
}

func (q *Collection) Name() string {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Name()
}

func (q *Collection) Drop() error {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Drop(context.Background())
}

func (q *Collection) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).CountDocuments(context.Background(), filter, opts...)
}

func (q *Collection) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).EstimatedDocumentCount(context.Background(), opts...)
}

func (q *Collection) Distinct(fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Distinct(context.Background(), fieldName, filter, opts...)
}

func (q *Collection) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) *ChangeStreamResult {
	res, err := q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Watch(context.Background(), pipeline, opts...)
	return &ChangeStreamResult{changeStream: res, err: err}
}

func (q *Collection) Indexes() *IndexView {
	return &IndexView{view: q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...).Indexes()}
}

// Context
// TRANSACTION
// QUERY
func (q *Collection) Context(ctx context.Context) *Collection {
	q.context = ctx
	return q
}

func (q *Collection) Aggregate(pipeline interface{}) *Aggregate {
	return NewAggregate(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), pipeline)
}

func (q *Collection) Find(filter interface{}) *Find {
	return NewFind(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter)
}

func (q *Collection) FindOne(filter interface{}) *FindOne {
	return NewFindOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter)
}

func (q *Collection) FindByID(id interface{}) *FindOne {
	return NewFindOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), bson.M{"_id": id})
}

func (q *Collection) FindOneAndDelete(filter interface{}) *FindOneAndDelete {
	return NewFindOneAndDelete(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter)
}

func (q *Collection) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplace {
	return NewFindOneAndReplace(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter, replacement)
}

func (q *Collection) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdate {
	return NewFindOneAndUpdate(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter, update)
}

// InsertOne
// INSERT
func (q *Collection) InsertOne(document interface{}) *InsertOne {
	return NewInsertOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), document)
}

func (q *Collection) InsertMany(document []interface{}) *InsertMany {
	return NewInsertMany(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), document)
}

func (q *Collection) DeleteOne(filter interface{}) *DeleteOne {
	return NewDeleteOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter)
}

func (q *Collection) DeleteMany(filter interface{}) *DeleteMany {
	return NewDeleteMany(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter)
}

func (q *Collection) UpdateOne(filter interface{}, update interface{}) *UpdateOne {
	return NewUpdateOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter, update)
}

func (q *Collection) UpdateMany(filter interface{}, update interface{}) *UpdateMany {
	return NewUpdateMany(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter, update)
}

func (q *Collection) ReplaceOne(filter interface{}, update interface{}) *ReplaceOne {
	return NewReplaceOne(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), filter, update)
}

// BulkWrite
// WRITE
func (q *Collection) BulkWrite(models []mongo.WriteModel) *BulkWrite {
	return NewBulkWrite(q.context, q.client.Database(q.db, q.databaseOptions...).Collection(q.collection, q.collectionOptions...), models)
}
