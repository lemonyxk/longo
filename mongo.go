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

func (q *Query) AggregateWithSession(sessionContext context.Context, pipeline interface{}, opts ...*options.AggregateOptions) *MultiResult {
	cursor, err := q.client.Database(q.db).Collection(q.collection).Aggregate(sessionContext, pipeline, opts...)
	return &MultiResult{cursor: cursor, err: err}
}

func (q *Query) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) *MultiResult {
	return q.AggregateWithSession(context.Background(), pipeline, opts...)
}

func (q *Query) FindWithSession(sessionContext context.Context, filter interface{}) *Find {
	return &Find{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: sessionContext}
}

func (q *Query) Find(filter interface{}) *Find {
	return q.FindWithSession(context.Background(), filter)
}

func (q *Query) FindOneWithSession(sessionContext context.Context, filter interface{}) *FindOne {
	return &FindOne{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: sessionContext}
}

func (q *Query) FindOne(filter interface{}) *FindOne {
	return q.FindOneWithSession(context.Background(), filter)
}

func (q *Query) FindOneAndDeleteWithSession(sessionContext context.Context, filter interface{}) *FindOneAndDelete {
	return &FindOneAndDelete{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, sessionContext: sessionContext}
}

func (q *Query) FindOneAndDelete(filter interface{}) *FindOneAndDelete {
	return q.FindOneAndDeleteWithSession(context.Background(), filter)
}

func (q *Query) FindOneAndReplaceWithSession(sessionContext context.Context, filter interface{}, replacement interface{}) *FindOneAndReplace {
	return &FindOneAndReplace{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, replacement: replacement, sessionContext: sessionContext}
}

func (q *Query) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplace {
	return q.FindOneAndReplaceWithSession(context.Background(), filter, replacement)
}

func (q *Query) FindOneAndUpdateWithSession(sessionContext context.Context, filter interface{}, update interface{}) *FindOneAndUpdate {
	return &FindOneAndUpdate{collection: q.client.Database(q.db).Collection(q.collection), filter: filter, update: update, sessionContext: sessionContext}
}

func (q *Query) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdate {
	return q.FindOneAndUpdateWithSession(context.Background(), filter, update)
}

func (q *Query) InsertOneWithSession(sessionContext context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertOne(sessionContext, document, opts...)
}

func (q *Query) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return q.InsertOneWithSession(context.Background(), document, opts...)
}

func (q *Query) InsertManyWithSession(sessionContext context.Context, document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertMany(sessionContext, document, opts...)
}

func (q *Query) InsertMany(document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return q.InsertManyWithSession(context.Background(), document, opts...)
}

func (q *Query) DeleteOneWithSession(sessionContext context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteOne(sessionContext, filter, opts...)
}

func (q *Query) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.DeleteOneWithSession(context.Background(), filter, opts...)
}

func (q *Query) DeleteManyWithSession(sessionContext context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteMany(sessionContext, filter, opts...)
}

func (q *Query) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.DeleteManyWithSession(context.Background(), filter, opts...)
}

func (q *Query) UpdateOneWithSession(sessionContext context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateOne(sessionContext, filter, update, opts...)
}

func (q *Query) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.UpdateOneWithSession(context.Background(), filter, update, opts...)
}

func (q *Query) UpdateManyWithSession(sessionContext context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateMany(sessionContext, filter, update, opts...)
}

func (q *Query) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.UpdateManyWithSession(context.Background(), filter, update, opts...)
}

func (q *Query) ReplaceOneWithSession(sessionContext context.Context, filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).ReplaceOne(sessionContext, filter, update, opts...)
}

func (q *Query) ReplaceOne(filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return q.ReplaceOneWithSession(context.Background(), filter, update, opts...)
}
