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

	"go.mongodb.org/mongo-driver/bson"
)

func NewModel(db, collection string) *Model {
	return &Model{
		DB:         db,
		Collection: collection,
	}
}

type Model struct {
	Handler    *Mgo
	DB         string
	Collection string
	Ctx        context.Context
}

type FindResult struct {
	Find *Find
}

type AggregateResult struct {
	Aggregate *Aggregate
}

func (p *AggregateResult) One(res interface{}) error {
	return p.Aggregate.One(res)
}

func (p *AggregateResult) All(res interface{}) error {
	return p.Aggregate.All(res)
}

func (p *FindResult) Sort(sort interface{}) *FindResult {
	p.Find.Sort(sort)
	return p
}

func (p *FindResult) Skip(skip int64) *FindResult {
	p.Find.Skip(skip)
	return p
}

func (p *FindResult) Limit(limit int64) *FindResult {
	p.Find.Limit(limit)
	return p
}

func (p *FindResult) Projection(projection interface{}) *FindResult {
	p.Find.Projection(projection)
	return p
}

func (p *FindResult) One(res interface{}) error {
	return p.Find.One(res)
}

func (p *FindResult) All(res interface{}) error {
	return p.Find.All(res)
}

func (p *Model) SetHandler(handler *Mgo) *Model {
	p.Handler = handler
	return p
}

func (p *Model) Query() *Query {
	return p.Handler.DB(p.DB).C(p.Collection).Context(p.Ctx)
}

func (p *Model) Context(ctx context.Context) *Model {
	p.Ctx = ctx
	return p
}

func (p *Model) Find(find interface{}) *FindResult {
	return &FindResult{Find: p.Handler.DB(p.DB).C(p.Collection).Find(find).Context(p.Ctx)}
}

func (p *Model) Count(find interface{}) (int64, error) {
	return p.Handler.DB(p.DB).C(p.Collection).CountDocuments(find)
}

func (p *Model) Set(filter interface{}, update interface{}) *UpdateResult {
	return p.Handler.DB(p.DB).C(p.Collection).UpdateMany(filter, bson.M{"$set": update}).Context(p.Ctx).Do()
}

func (p *Model) Update(filter interface{}, update interface{}) *UpdateResult {
	return p.Handler.DB(p.DB).C(p.Collection).UpdateMany(filter, update).Context(p.Ctx).Do()
}

func (p *Model) Insert(document ...interface{}) *InsertManyResult {
	return p.Handler.DB(p.DB).C(p.Collection).InsertMany(document).Context(p.Ctx).Do()
}

func (p *Model) Delete(filter interface{}) *DeleteResult {
	return p.Handler.DB(p.DB).C(p.Collection).DeleteMany(filter).Context(p.Ctx).Do()
}

func (p *Model) FindAndModify(filter interface{}, update interface{}) *FindOneAndUpdate {
	return p.Handler.DB(p.DB).C(p.Collection).FindOneAndUpdate(filter, update).Context(p.Ctx)
}

func (p *Model) Aggregate(pipeline interface{}) *AggregateResult {
	return &AggregateResult{Aggregate: p.Handler.DB(p.DB).C(p.Collection).Aggregate(pipeline).Context(p.Ctx)}
}
