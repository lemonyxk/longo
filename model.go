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
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewModel[T ~[]*E, E any](ctx context.Context, handler *Mgo) *DBModel[T, E] {
	return &DBModel[T, E]{
		Handler: handler,
		Ctx:     ctx,
	}
}

type DBModel[T ~[]*E, E any] struct {
	Handler *Mgo
	Ctx     context.Context
}

func (p *DBModel[T, E]) DB(db string, opt ...*options.DatabaseOptions) *CModel[T, E] {
	return &CModel[T, E]{
		DB:              db,
		DatabaseOptions: opt,
		DBModel:         p,
	}
}

type CModel[T ~[]*E, E any] struct {
	DB              string
	DatabaseOptions []*options.DatabaseOptions
	*DBModel[T, E]
}

func (p *CModel[T, E]) C(collection string, opt ...*options.CollectionOptions) *Model[T, E] {
	return &Model[T, E]{
		DB:                p.DB,
		C:                 collection,
		DatabaseOptions:   p.DatabaseOptions,
		CollectionOptions: opt,
		Ctx:               p.Ctx,
		Handler:           p.Handler,
	}
}

// Model is a mongodb model
type Model[T ~[]*E, E any] struct {
	Handler           *Mgo
	DB                string
	C                 string
	Ctx               context.Context
	DatabaseOptions   []*options.DatabaseOptions
	CollectionOptions []*options.CollectionOptions
}

// func (p *Model[T,E]) SetHandler(ctx context.Context, handler *Mgo) *Model[T,E] {
// 	p.Handler = handler
// 	p.Ctx = ctx
// 	return p
// }

func (p *Model[T, E]) CreateIndex() error {
	var t E
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
			return err
		}
	}

	return nil
}

func (p *Model[T, E]) Database() *Database {
	return p.Handler.DB(p.DB, p.DatabaseOptions...)
}

func (p *Model[T, E]) Collection() *Collection {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Context(p.Ctx)
}

// func (p *Model[T,E]) Context(ctx context.Context) *Model[T,E] {
// 	p.Ctx = ctx
// 	return p
// }

func (p *Model[T, E]) Count(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	var ref = reflect.ValueOf(filter)
	if ref.IsNil() || (ref.Kind() == reflect.Map && ref.Len() == 0) {
		return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).EstimatedDocumentCount()
	}
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).CountDocuments(filter, opts...)
}

func (p *Model[T, E]) Set(filter interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(filter, bson.M{"$set": update}).Context(p.Ctx)
}

func (p *Model[T, E]) SetByID(id interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(bson.M{"_id": id}, bson.M{"$set": update}).Context(p.Ctx)
}

func (p *Model[T, E]) Update(filter interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(filter, update).Context(p.Ctx)
}

func (p *Model[T, E]) UpdateByID(id interface{}, update interface{}) *UpdateMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).UpdateMany(bson.M{"_id": id}, update).Context(p.Ctx)
}

func (p *Model[T, E]) Insert(document ...*E) *InsertMany {
	var docs = make([]interface{}, len(document))
	for i := 0; i < len(docs); i++ {
		docs[i] = document[i]
	}
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).InsertMany(docs).Context(p.Ctx)
}

func (p *Model[T, E]) Delete(filter interface{}) *DeleteMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).DeleteMany(filter).Context(p.Ctx)
}

func (p *Model[T, E]) DeleteByID(id interface{}) *DeleteMany {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).DeleteMany(bson.M{"_id": id}).Context(p.Ctx)
}

func (p *Model[T, E]) BulkWrite(models []mongo.WriteModel) *BulkWrite {
	return p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).BulkWrite(models).Context(p.Ctx)
}

// FindOneAndReplaceResult is the result of a FindOneAndReplace method.

func (p *Model[T, E]) FindOneAndReplace(filter interface{}, replacement interface{}) *FindOneAndReplaceResult[T, E] {
	return &FindOneAndReplaceResult[T, E]{FindOneAndReplace: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndReplace(filter, replacement).Context(p.Ctx)}
}

type FindOneAndReplaceResult[T ~[]*E, E any] struct {
	*FindOneAndReplace
}

func (f *FindOneAndReplaceResult[T, E]) Hit(hit interface{}) *FindOneAndReplaceResult[T, E] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Sort(sort interface{}) *FindOneAndReplaceResult[T, E] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Projection(projection interface{}) *FindOneAndReplaceResult[T, E] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Upsert() *FindOneAndReplaceResult[T, E] {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndReplaceResult[T, E]) ReturnDocument() *FindOneAndReplaceResult[T, E] {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Context(ctx context.Context) *FindOneAndReplaceResult[T, E] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Option(opt *options.FindOneAndReplaceOptions) *FindOneAndReplaceResult[T, E] {
	f.option = opt
	return f
}

func (f *FindOneAndReplaceResult[T, E]) Exec() (*E, error) {
	var res E
	var sig = &SingleResult{singleResult: f.collection.FindOneAndReplace(f.sessionContext, f.filter, f.replacement, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindOneAndDeleteResult is the result of a FindOneAndDelete operation.

func (p *Model[T, E]) FindOneAndDelete(filter interface{}) *FindOneAndDeleteResult[T, E] {
	return &FindOneAndDeleteResult[T, E]{FindOneAndDelete: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndDelete(filter).Context(p.Ctx)}
}

type FindOneAndDeleteResult[T ~[]*E, E any] struct {
	*FindOneAndDelete
}

func (f *FindOneAndDeleteResult[T, E]) Hit(hit interface{}) *FindOneAndDeleteResult[T, E] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndDeleteResult[T, E]) Sort(sort interface{}) *FindOneAndDeleteResult[T, E] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndDeleteResult[T, E]) Projection(projection interface{}) *FindOneAndDeleteResult[T, E] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndDeleteResult[T, E]) Context(ctx context.Context) *FindOneAndDeleteResult[T, E] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndDeleteResult[T, E]) Option(opt *options.FindOneAndDeleteOptions) *FindOneAndDeleteResult[T, E] {
	f.option = opt
	return f
}

func (f *FindOneAndDeleteResult[T, E]) Exec() (*E, error) {
	var res E
	var sig = &SingleResult{singleResult: f.collection.FindOneAndDelete(f.sessionContext, f.filter, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindOneAndUpdateResult is a find one and update

func (p *Model[T, E]) FindOneAndUpdate(filter interface{}, update interface{}) *FindOneAndUpdateResult[T, E] {
	return &FindOneAndUpdateResult[T, E]{FindOneAndUpdate: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOneAndUpdate(filter, update).Context(p.Ctx)}
}

type FindOneAndUpdateResult[T ~[]*E, E any] struct {
	*FindOneAndUpdate
}

func (f *FindOneAndUpdateResult[T, E]) Hit(hit interface{}) *FindOneAndUpdateResult[T, E] {
	f.option.Hint = hit
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Sort(sort interface{}) *FindOneAndUpdateResult[T, E] {
	f.option.Sort = sort
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Projection(projection interface{}) *FindOneAndUpdateResult[T, E] {
	f.option.Projection = projection
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Upsert() *FindOneAndUpdateResult[T, E] {
	var t = true
	f.option.Upsert = &t
	return f
}

func (f *FindOneAndUpdateResult[T, E]) ReturnDocument() *FindOneAndUpdateResult[T, E] {
	var t = options.After
	f.option.ReturnDocument = &t
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Context(ctx context.Context) *FindOneAndUpdateResult[T, E] {
	f.sessionContext = ctx
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Option(opt *options.FindOneAndUpdateOptions) *FindOneAndUpdateResult[T, E] {
	f.option = opt
	return f
}

func (f *FindOneAndUpdateResult[T, E]) Exec() (*E, error) {
	var res E
	var sig = &SingleResult{singleResult: f.collection.FindOneAndUpdate(f.sessionContext, f.filter, f.update, f.option)}
	var err = sig.One(&res)
	return &res, err
}

// FindResult is the result of a Find operation.

func (p *Model[T, E]) Find(filter interface{}) *FindResult[T, E] {
	return &FindResult[T, E]{Find: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Find(filter).Context(p.Ctx)}
}

type FindResult[T ~[]*E, E any] struct {
	Find *Find
}

func (p *FindResult[T, E]) Sort(sort interface{}) *FindResult[T, E] {
	p.Find.Sort(sort)
	return p
}

func (p *FindResult[T, E]) Skip(skip int64) *FindResult[T, E] {
	p.Find.Skip(skip)
	return p
}

func (p *FindResult[T, E]) Limit(limit int64) *FindResult[T, E] {
	p.Find.Limit(limit)
	return p
}

func (p *FindResult[T, E]) Projection(projection interface{}) *FindResult[T, E] {
	p.Find.Projection(projection)
	return p
}

func (p *FindResult[T, E]) Hit(res interface{}) *FindResult[T, E] {
	p.Find.Hit(res)
	return p
}

func (p *FindResult[T, E]) Context(ctx context.Context) *FindResult[T, E] {
	p.Find.Context(ctx)
	return p
}

//func (p *FindResult[T, E]) One() (*E, error) {
//	var res E
//	var err = p.Find.One(&res)
//	return &res, err
//}

func (p *FindResult[T, E]) All() (T, error) {
	var res T
	var err = p.Find.All(&res)
	return res, err
}

// FindByID findOneResult is the result of a FindOne operation.
func (p *Model[T, E]) FindByID(id interface{}) *FindOneResult[T, E] {
	return &FindOneResult[T, E]{FindOne: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOne(bson.M{"_id": id}).Context(p.Ctx)}
}

func (p *Model[T, E]) FindOne(filter interface{}) *FindOneResult[T, E] {
	return &FindOneResult[T, E]{FindOne: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).FindOne(filter).Context(p.Ctx)}
}

type FindOneResult[T ~[]*E, E any] struct {
	FindOne *FindOne
}

func (p *FindOneResult[T, E]) Sort(sort interface{}) *FindOneResult[T, E] {
	p.FindOne.Sort(sort)
	return p
}

func (p *FindOneResult[T, E]) Skip(skip int64) *FindOneResult[T, E] {
	p.FindOne.Skip(skip)
	return p
}

func (p *FindOneResult[T, E]) Projection(projection interface{}) *FindOneResult[T, E] {
	p.FindOne.Projection(projection)
	return p
}

func (p *FindOneResult[T, E]) Hit(res interface{}) *FindOneResult[T, E] {
	p.FindOne.Hit(res)
	return p
}

func (p *FindOneResult[T, E]) Context(ctx context.Context) *FindOneResult[T, E] {
	p.FindOne.Context(ctx)
	return p
}

func (p *FindOneResult[T, E]) One() (*E, error) {
	var res E
	var err = p.FindOne.One(&res)
	return &res, err
}

// AggregateResult is the result from an aggregate operation.

func (p *Model[T, E]) Aggregate(pipeline interface{}) *AggregateResult {
	return &AggregateResult{Aggregate: p.Handler.DB(p.DB, p.DatabaseOptions...).C(p.C, p.CollectionOptions...).Aggregate(pipeline).Context(p.Ctx)}
}

type AggregateResult struct {
	*Aggregate
}

//func (p *AggregateResult) One(res interface{}) error {
//	return p.Aggregate.One(res)
//}

func (p *AggregateResult) All(res interface{}) error {
	return p.Aggregate.All(res)
}

func (p *AggregateResult) Context(ctx context.Context) *AggregateResult {
	p.Aggregate.Context(ctx)
	return p
}
