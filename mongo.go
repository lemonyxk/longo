/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 15:55
**/

package mongo

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

func (q *Query) DropWithSession(sessionContext context.Context) error {
	return q.client.Database(q.db).Collection(q.collection).Drop(sessionContext)
}

func (q *Query) Drop() error {
	return q.DropWithSession(context.Background())
}

func (q *Query) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) *MultiResult {
	cursor, err := q.client.Database(q.db).Collection(q.collection).Aggregate(context.Background(), pipeline, opts...)
	return &MultiResult{cursor: cursor, err: err}
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

func (q *Query) Find(filter interface{}, opts ...*options.FindOptions) *MultiResult {
	cursor, err := q.client.Database(q.db).Collection(q.collection).Find(context.Background(), filter, opts...)
	return &MultiResult{cursor: cursor, err: err}
}

func (q *Query) FindOne(filter interface{}, opts ...*options.FindOneOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOne(context.Background(), filter, opts...)}
}

func (q *Query) FindOneAndDelete(filter interface{}, opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndDelete(context.Background(), filter, opts...)}
}

func (q *Query) FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndReplace(context.Background(), filter, replacement, opts...)}
}

func (q *Query) FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndUpdate(context.Background(), filter, update, opts...)}
}

func (q *Query) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) *ChangeStreamResult {
	res, err := q.client.Database(q.db).Collection(q.collection).Watch(context.Background(), pipeline, opts...)
	return &ChangeStreamResult{changeStream: res, err: err}
}

func (q *Query) Indexes() *IndexView {
	return &IndexView{view: q.client.Database(q.db).Collection(q.collection).Indexes()}
}
