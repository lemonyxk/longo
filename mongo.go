/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 15:55
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadPreference
// primary
// primarypreferred
// secondary
// secondarypreferred
// nearest

// ReadConcern
// local
// majority
// linearizable
// available
// snapshot

// WriteConcern
// w:0|1|majority
// j:false|true
// wtimeout

type Query struct {
	client     *mongo.Client
	config     Config
	db         string
	collection string
	ctx        context.Context
}

func (q *Query) SetContext(ctx context.Context) *Query {
	q.ctx = ctx
	return q
}

func (q *Query) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return q.client.Database(q.db).Collection(q.collection).Clone(opts...)
}

func (q *Query) Name() string {
	return q.client.Database(q.db).Collection(q.collection).Name()
}

func (q *Query) Drop() error {
	return q.client.Database(q.db).Collection(q.collection).Drop(context.Background())
}

func (q *Query) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).CountDocuments(context.Background(), filter, opts...)
}

func (q *Query) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).EstimatedDocumentCount(context.Background(), opts...)
}

func (q *Query) Distinct(fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.client.Database(q.db).Collection(q.collection).Distinct(context.Background(), fieldName, filter, opts...)
}

func (q *Query) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) *ChangeStreamResult {
	res, err := q.client.Database(q.db).Collection(q.collection).Watch(context.Background(), pipeline, opts...)
	return &ChangeStreamResult{changeStream: res, err: err}
}

func (q *Query) Indexes() *IndexView {
	return &IndexView{view: q.client.Database(q.db).Collection(q.collection).Indexes()}
}

// TRANSACTION

func (q *Query) Aggregate(pipeline interface{}) *Aggregate {
	return &Aggregate{collection: q.client.Database(q.db).Collection(q.collection), pipeline: pipeline, sessionContext: q.ctx}
}

func (q *Query) Find(filter interface{}) *Find {
	return &Find{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: q.ctx}
}

func (q *Query) FindOne(filter interface{}) *FindOne {
	return &FindOne{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: q.ctx}
}

func (q *Query) FindOneAndDelete(filter interface{}) *FindOneAndDelete {
	return &FindOneAndDelete{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: q.ctx}
}

func (q *Query) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplace {
	return &FindOneAndReplace{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, replacement: replacement, sessionContext: q.ctx}
}

func (q *Query) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdate {
	return &FindOneAndUpdate{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, update: update, sessionContext: q.ctx}
}

func (q *Query) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertOne(q.ctx, document, opts...)
}

func (q *Query) InsertMany(document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertMany(q.ctx, document, opts...)
}

func (q *Query) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteOne(q.ctx, filter, opts...)
}

func (q *Query) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteMany(q.ctx, filter, opts...)
}

func (q *Query) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateOne(q.ctx, filter, update, opts...)
}

func (q *Query) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateMany(q.ctx, filter, update, opts...)
}

func (q *Query) ReplaceOne(filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).ReplaceOne(q.ctx, filter, update, opts...)
}
