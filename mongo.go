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
	context    context.Context
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
// QUERY
func (q *Query) Context(ctx context.Context) *Query {
	q.context = ctx
	return q
}

func (q *Query) Aggregate(pipeline interface{}) *Aggregate {
	return NewAggregate(q.context, q.client.Database(q.db).Collection(q.collection), pipeline)
}

func (q *Query) Find(filter interface{}) *Find {
	return NewFind(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Query) FindOne(filter interface{}) *FindOne {
	return NewFindOne(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Query) FindOneAndDelete(filter interface{}) *FindOneAndDelete {
	return NewFindOneAndDelete(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Query) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplace {
	return NewFindOneAndReplace(q.context, q.client.Database(q.db).Collection(q.collection), filter, replacement)
}

func (q *Query) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdate {
	return NewFindOneAndUpdate(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

// INSERT
func (q *Query) InsertOne(document interface{}) *InsertOne {
	return NewInsertOne(q.context, q.client.Database(q.db).Collection(q.collection), document)
}

func (q *Query) InsertMany(document []interface{}) *InsertMany {
	return NewInsertMany(q.context, q.client.Database(q.db).Collection(q.collection), document)
}

func (q *Query) DeleteOne(filter interface{}) *DeleteOne {
	return NewDeleteOne(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Query) DeleteMany(filter interface{}) *DeleteMany {
	return NewDeleteMany(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Query) UpdateOne(filter interface{}, update interface{}) *UpdateOne {
	return NewUpdateOne(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

func (q *Query) UpdateMany(filter interface{}, update interface{}) *UpdateMany {
	return NewUpdateMany(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

func (q *Query) ReplaceOne(filter interface{}, update interface{}) *ReplaceOne {
	return NewReplaceOne(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}
