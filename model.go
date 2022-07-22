/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2020-04-08 17:39
**/

package longo

import (
	"context"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewModel[T any](db, collection string) *Model[T] {
	return &Model[T]{
		DB: db,
		C:  collection,
	}
}

// Model is a mongodb model
type Model[T any] struct {
	Handler *Mgo
	M       T
	DB      string
	C       string
	Ctx     context.Context
}

func (p *Model[T]) SetHandler(handler *Mgo) *Model[T] {
	p.Handler = handler
	return p
}

func (p *Model[T]) CreateIndex() *Model[T] {
	var srcType = reflect.TypeOf(p.M)
	if srcType.Kind() != reflect.Struct {
		panic("model must be struct")
	}

	// get all indexes of the model
	var result []Index
	_ = p.Collection().Indexes().List().All(context.Background(), &result)

	// ignore _id index
	var mr = make(map[string]bool)
	for i := 0; i < len(result); i++ {
		var n = result[i].Name
		if n != "_id_" {
			mr[n] = true
		}
	}

	var index []string
	var indexes []string
	var n = srcType.NumField()
	for i := 0; i < n; i++ {
		var field = srcType.Field(i)
		var indexName = field.Tag.Get("index")
		if indexName != "" && !mr[indexName+"_1"] {
			index = append(index, indexName)
		}

		var indexesName = field.Tag.Get("indexes")
		if indexesName != "" && !mr[strings.ReplaceAll(indexesName, ",", "_1_")+"_1"] {
			indexes = append(indexes, indexesName)
		}
	}

	var create []mongo.IndexModel

	for i := 0; i < len(index); i++ {
		create = append(create, mongo.IndexModel{
			Keys:    bson.M{index[i]: 1},
			Options: &options.IndexOptions{},
		})
	}

	for i := 0; i < len(indexes); i++ {
		var keys = bson.D{}
		var im = mongo.IndexModel{Keys: bson.E{}, Options: &options.IndexOptions{}}

		var list = strings.Split(indexes[i], ",")
		for j := 0; j < len(list); j++ {
			keys = append(keys, bson.E{Key: list[j], Value: 1})
		}

		im.Keys = keys

		create = append(create, im)
	}

	if len(create) == 0 {
		return p
	}

	_, _ = p.Collection().Indexes().CreateMany(create)

	return p
}

func (p *Model[T]) Collection() *Collection {
	return p.Handler.DB(p.DB).C(p.C).Context(p.Ctx)
}

func (p *Model[T]) Context(ctx context.Context) *Model[T] {
	p.Ctx = ctx
	return p
}

func (p *Model[T]) Find(filter interface{}) *FindResult[T] {
	return &FindResult[T]{Find: p.Handler.DB(p.DB).C(p.C).Find(filter).Context(p.Ctx)}
}

func (p *Model[T]) Count(filter interface{}) (int64, error) {
	return p.Handler.DB(p.DB).C(p.C).CountDocuments(filter)
}

func (p *Model[T]) Set(filter interface{}, update interface{}) *UpdateResult {
	return p.Handler.DB(p.DB).C(p.C).UpdateMany(filter, bson.M{"$set": update}).Context(p.Ctx).Do()
}

func (p *Model[T]) Update(filter interface{}, update interface{}) *UpdateResult {
	return p.Handler.DB(p.DB).C(p.C).UpdateMany(filter, update).Context(p.Ctx).Do()
}

func (p *Model[T]) Insert(document ...*T) *InsertManyResult {
	var docs = make([]interface{}, len(document))
	for i := 0; i < len(docs); i++ {
		docs[i] = document[i]
	}
	return p.Handler.DB(p.DB).C(p.C).InsertMany(docs).Context(p.Ctx).Do()
}

func (p *Model[T]) Delete(filter interface{}) *DeleteResult {
	return p.Handler.DB(p.DB).C(p.C).DeleteMany(filter).Context(p.Ctx).Do()
}

func (p *Model[T]) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdateResult[T] {
	return &FindOneAndUpdateResult[T]{FindOneAndUpdate: p.Handler.DB(p.DB).C(p.C).FindOneAndUpdate(filter, update).Context(p.Ctx)}
}

func (p *Model[T]) FindOneAndDelete(filter interface{}) *FindOneAndDeleteResult[T] {
	return &FindOneAndDeleteResult[T]{FindOneAndDelete: p.Handler.DB(p.DB).C(p.C).FindOneAndDelete(filter).Context(p.Ctx)}
}

func (p *Model[T]) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplaceResult[T] {
	return &FindOneAndReplaceResult[T]{FindOneAndReplace: p.Handler.DB(p.DB).C(p.C).FindOneAndReplace(filter, replacement).Context(p.Ctx)}
}

// FindOneAndReplaceResult is the result of a FindOneAndReplace method.
type FindOneAndReplaceResult[T any] struct {
	*FindOneAndReplace
}

func (f *FindOneAndReplaceResult[T]) Hit(hit interface{}) *FindOneAndReplaceResult[T] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndReplaceResult[T]) Sort(sort interface{}) *FindOneAndReplaceResult[T] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndReplaceResult[T]) Projection(projection interface{}) *FindOneAndReplaceResult[T] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndReplaceResult[T]) Upsert() *FindOneAndReplaceResult[T] {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndReplaceResult[T]) ReturnDocument() *FindOneAndReplaceResult[T] {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndReplaceResult[T]) Context(ctx context.Context) *FindOneAndReplaceResult[T] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndReplaceResult[T]) Option(opt *options.FindOneAndReplaceOptions) *FindOneAndReplaceResult[T] {
	f.option = opt
	return f
}

func (f *FindOneAndReplaceResult[T]) Do() Result[T] {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.option)}
	var err = sig.Do(&res)
	return Result[T]{res, err}
}

// FindOneAndDeleteResult is the result of a FindOneAndDelete operation.
type FindOneAndDeleteResult[T any] struct {
	*FindOneAndDelete
}

func (f *FindOneAndDeleteResult[T]) Hit(hit interface{}) *FindOneAndDeleteResult[T] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndDeleteResult[T]) Sort(sort interface{}) *FindOneAndDeleteResult[T] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndDeleteResult[T]) Projection(projection interface{}) *FindOneAndDeleteResult[T] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndDeleteResult[T]) Context(ctx context.Context) *FindOneAndDeleteResult[T] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndDeleteResult[T]) Option(opt *options.FindOneAndDeleteOptions) *FindOneAndDeleteResult[T] {
	f.option = opt
	return f
}

func (f *FindOneAndDeleteResult[T]) Do() Result[T] {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	var err = sig.Do(&res)
	return Result[T]{res, err}
}

// FindOneAndUpdateResult is a find one and update
type FindOneAndUpdateResult[T any] struct {
	*FindOneAndUpdate
}

func (f *FindOneAndUpdateResult[T]) Hit(hit interface{}) *FindOneAndUpdateResult[T] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndUpdateResult[T]) Sort(sort interface{}) *FindOneAndUpdateResult[T] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndUpdateResult[T]) Projection(projection interface{}) *FindOneAndUpdateResult[T] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndUpdateResult[T]) Upsert() *FindOneAndUpdateResult[T] {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndUpdateResult[T]) ReturnDocument() *FindOneAndUpdateResult[T] {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndUpdateResult[T]) Context(ctx context.Context) *FindOneAndUpdateResult[T] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndUpdateResult[T]) Option(opt *options.FindOneAndUpdateOptions) *FindOneAndUpdateResult[T] {
	f.option = opt
	return f
}

func (f *FindOneAndUpdateResult[T]) Do() Result[T] {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, f.option)}
	var err = sig.Do(&res)
	return Result[T]{res, err}
}

// FindResult is the result of a Find operation.
type FindResult[T any] struct {
	Find *Find
}

func (p *FindResult[T]) Sort(sort interface{}) *FindResult[T] {
	p.Find.Sort(sort)
	return p
}

func (p *FindResult[T]) Skip(skip int64) *FindResult[T] {
	p.Find.Skip(skip)
	return p
}

func (p *FindResult[T]) Limit(limit int64) *FindResult[T] {
	p.Find.Limit(limit)
	return p
}

func (p *FindResult[T]) Projection(projection interface{}) *FindResult[T] {
	p.Find.Projection(projection)
	return p
}

func (p *FindResult[T]) One() (*T, error) {
	var res T
	var err = p.Find.One(&res)
	return &res, err
}

func (p *FindResult[T]) All() ([]*T, error) {
	var res []*T
	var err = p.Find.All(&res)
	return res, err
}

func (p *FindResult[T]) Hit(res interface{}) *FindResult[T] {
	p.Find.Hit(res)
	return p
}

// AggregateResult is the result from an aggregate operation.
type AggregateResult struct {
	*Aggregate
}

func (p *AggregateResult) One(res interface{}) error {
	return p.Aggregate.One(res)
}

func (p *AggregateResult) All(res interface{}) error {
	return p.Aggregate.All(res)
}

func (p *Model[T]) Aggregate(pipeline interface{}) *AggregateResult {
	return &AggregateResult{Aggregate: p.Handler.DB(p.DB).C(p.C).Aggregate(pipeline).Context(p.Ctx)}
}
