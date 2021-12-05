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
	client     *mongo.Client
	config     Config
	db         string
	collection string
	context    context.Context
}

func (q *Collection) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return q.client.Database(q.db).Collection(q.collection).Clone(opts...)
}

func (q *Collection) Name() string {
	return q.client.Database(q.db).Collection(q.collection).Name()
}

func (q *Collection) Drop() error {
	return q.client.Database(q.db).Collection(q.collection).Drop(context.Background())
}

func (q *Collection) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).CountDocuments(context.Background(), filter, opts...)
}

func (q *Collection) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).EstimatedDocumentCount(context.Background(), opts...)
}

func (q *Collection) Distinct(fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.client.Database(q.db).Collection(q.collection).Distinct(context.Background(), fieldName, filter, opts...)
}

func (q *Collection) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) *ChangeStreamResult {
	res, err := q.client.Database(q.db).Collection(q.collection).Watch(context.Background(), pipeline, opts...)
	return &ChangeStreamResult{changeStream: res, err: err}
}

func (q *Collection) Indexes() *IndexView {
	return &IndexView{view: q.client.Database(q.db).Collection(q.collection).Indexes()}
}

// Context
// TRANSACTION
// QUERY
func (q *Collection) Context(ctx context.Context) *Collection {
	q.context = ctx
	return q
}

func (q *Collection) Aggregate(pipeline interface{}) *Aggregate {
	return NewAggregate(q.context, q.client.Database(q.db).Collection(q.collection), pipeline)
}

func (q *Collection) Find(filter interface{}) *Find {
	return NewFind(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Collection) FindOne(filter interface{}) *FindOne {
	return NewFindOne(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Collection) FindOneAndDelete(filter interface{}) *FindOneAndDelete {
	return NewFindOneAndDelete(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Collection) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplace {
	return NewFindOneAndReplace(q.context, q.client.Database(q.db).Collection(q.collection), filter, replacement)
}

func (q *Collection) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdate {
	return NewFindOneAndUpdate(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

// InsertOne
// INSERT
func (q *Collection) InsertOne(document interface{}) *InsertOne {
	return NewInsertOne(q.context, q.client.Database(q.db).Collection(q.collection), document)
}

func (q *Collection) InsertMany(document []interface{}) *InsertMany {
	return NewInsertMany(q.context, q.client.Database(q.db).Collection(q.collection), document)
}

func (q *Collection) DeleteOne(filter interface{}) *DeleteOne {
	return NewDeleteOne(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Collection) DeleteMany(filter interface{}) *DeleteMany {
	return NewDeleteMany(q.context, q.client.Database(q.db).Collection(q.collection), filter)
}

func (q *Collection) UpdateOne(filter interface{}, update interface{}) *UpdateOne {
	return NewUpdateOne(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

func (q *Collection) UpdateMany(filter interface{}, update interface{}) *UpdateMany {
	return NewUpdateMany(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}

func (q *Collection) ReplaceOne(filter interface{}, update interface{}) *ReplaceOne {
	return NewReplaceOne(q.context, q.client.Database(q.db).Collection(q.collection), filter, update)
}
