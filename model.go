/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2020-04-08 17:39
**/

package longo

import (
	"context"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewModel[T any](ctx context.Context, handler *Mgo) *DBModel[T] {
	return &DBModel[T]{
		Handler: handler,
		Ctx:     ctx,
	}
}

type DBModel[T any] struct {
	Handler *Mgo
	Ctx     context.Context
}

func (p *DBModel[T]) DB(db string, opt ...*options.DatabaseOptions) *CModel[T] {
	return &CModel[T]{
		DB:              db,
		DatabaseOptions: opt,
		DBModel:         p,
	}
}

type CModel[T any] struct {
	DB              string
	DatabaseOptions []*options.DatabaseOptions
	*DBModel[T]
}

func (p *CModel[T]) C(collection string, opt ...*options.CollectionOptions) *Model[T] {
	return &Model[T]{
		DB:                p.DB,
		C:                 collection,
		DatabaseOptions:   p.DatabaseOptions,
		CollectionOptions: opt,
		Ctx:               p.Ctx,
		Handler:           p.Handler,
	}
}

// Model is a mongodb model
type Model[T any] struct {
	Handler           *Mgo
	DB                string
	C                 string
	Ctx               context.Context
	DatabaseOptions   []*options.DatabaseOptions
	CollectionOptions []*options.CollectionOptions
}

// func (p *Model[T]) SetHandler(ctx context.Context, handler *Mgo) *Model[T] {
// 	p.Handler = handler
// 	p.Ctx = ctx
// 	return p
// }

func (p *Model[T]) CreateIndex() *Model[T] {
	var t T
	var srcType = reflect.TypeOf(t)
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
		var indexNameArr = strings.Split(indexName, " ")
		if indexName != "" && len(indexNameArr) > 0 {
			for j := 0; j < len(indexNameArr); j++ {
				if !inMapArray(mr, indexNameArr[j]) && indexNameArr[j] != "" {
					index = append(index, indexNameArr[j])
				}
			}
		}

		var indexesName = field.Tag.Get("indexes")
		var indexesNameArr = strings.Split(indexesName, " ")
		if indexesName != "" && len(indexesNameArr) > 0 {
			for j := 0; j < len(indexesNameArr); j++ {
				if !inMapArray(mr, indexesNameArr[j]) && indexesNameArr[j] != "" {
					indexes = append(indexes, indexesNameArr[j])
				}
			}
		}
	}

	var create = parseIndex(index)
	create = append(create, parseIndexes(indexes)...)

	if len(create) > 0 {
		_, err := p.Collection().Indexes().CreateMany(create)
		if err != nil {
			panic(err)
		}
	}

	return p
}

func (p *Model[T]) Database() *Database {
	return p.Handler.DB(p.DB, p.DatabaseOptions...)
}

func (p *Model[T]) Collection() *Collection {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Context(p.Ctx)
}

// func (p *Model[T]) Context(ctx context.Context) *Model[T] {
// 	p.Ctx = ctx
// 	return p
// }

func (p *Model[T]) Count(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	if filter == nil {
		return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).EstimatedDocumentCount()
	}
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).CountDocuments(filter, opts...)
}

func (p *Model[T]) Set(filter interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(filter, bson.M{"$set": update}).Context(p.Ctx)
}

func (p *Model[T]) SetByID(id interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(bson.M{"_id": id}, bson.M{"$set": update}).Context(p.Ctx)
}

func (p *Model[T]) Update(filter interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(filter, update).Context(p.Ctx)
}

func (p *Model[T]) UpdateByID(id interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(bson.M{"_id": id}, update).Context(p.Ctx)
}

func (p *Model[T]) Insert(document ...*T) *InsertMany {
	var docs = make([]interface{}, len(document))
	for i := 0; i < len(docs); i++ {
		docs[i] = document[i]
	}
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).InsertMany(docs).Context(p.Ctx)
}

func (p *Model[T]) Delete(filter interface{}) *DeleteMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).DeleteMany(filter).Context(p.Ctx)
}

func (p *Model[T]) DeleteByID(id interface{}) *DeleteMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).DeleteMany(bson.M{"_id": id}).Context(p.Ctx)
}

// FindOneAndReplaceResult is the result of a FindOneAndReplace method.

func (p *Model[T]) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplaceResult[T] {
	return &FindOneAndReplaceResult[T]{FindOneAndReplace: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndReplace(filter, replacement).Context(p.Ctx)}
}

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

func (f *FindOneAndReplaceResult[T]) Exec() (*T, error) {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindOneAndDeleteResult is the result of a FindOneAndDelete operation.

func (p *Model[T]) FindOneAndDelete(filter interface{}) *FindOneAndDeleteResult[T] {
	return &FindOneAndDeleteResult[T]{FindOneAndDelete: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndDelete(filter).Context(p.Ctx)}
}

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

func (f *FindOneAndDeleteResult[T]) Exec() (*T, error) {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindOneAndUpdateResult is a find one and update

func (p *Model[T]) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdateResult[T] {
	return &FindOneAndUpdateResult[T]{FindOneAndUpdate: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndUpdate(filter, update).Context(p.Ctx)}
}

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

func (f *FindOneAndUpdateResult[T]) Exec() (*T, error) {
	var res T
	var sig = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindResult is the result of a Find operation.

func (p *Model[T]) Find(filter interface{}) *FindResult[T] {
	return &FindResult[T]{Find: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Find(filter).Context(p.Ctx)}
}

func (p *Model[T]) FindByID(id interface{}) *FindResult[T] {
	return &FindResult[T]{Find: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Find(bson.M{"_id": id}).Context(p.Ctx)}
}

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

func (p *FindResult[T]) Hit(res interface{}) *FindResult[T] {
	p.Find.Hit(res)
	return p
}

func (p *FindResult[T]) Context(ctx context.Context) *FindResult[T] {
	p.Find.Context(ctx)
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

// AggregateResult is the result from an aggregate operation.

func (p *Model[T]) Aggregate(pipeline interface{}) *AggregateResult {
	return &AggregateResult{Aggregate: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Aggregate(pipeline).Context(p.Ctx)}
}

type AggregateResult struct {
	*Aggregate
}

func (p *AggregateResult) One(res interface{}) error {
	return p.Aggregate.One(res)
}

func (p *AggregateResult) All(res interface{}) error {
	return p.Aggregate.All(res)
}

func (p *AggregateResult) Context(ctx context.Context) *AggregateResult {
	p.Aggregate.Context(ctx)
	return p
}
