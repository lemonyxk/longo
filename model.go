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
	"go.mongodb.org/mongo-driver/bson"
)

func NewModel(db, collection string) *model {
	return &model{
		db:         db,
		collection: collection,
	}
}

type model struct {
	handler    *Mgo
	db         string
	collection string
}

type findResult struct {
	find *Find
}

type aggregateResult struct {
	aggregate *Aggregate
}

func (p *aggregateResult) One(res interface{}) error {
	return p.aggregate.One(res)
}

func (p *aggregateResult) All(res interface{}) error {
	return p.aggregate.All(res)
}

func (p *findResult) Sort(sort interface{}) *findResult {
	p.find.Sort(sort)
	return p
}

func (p *findResult) Skip(skip int64) *findResult {
	p.find.Skip(skip)
	return p
}

func (p *findResult) Limit(limit int64) *findResult {
	p.find.Limit(limit)
	return p
}

func (p *findResult) Projection(projection interface{}) *findResult {
	p.find.Projection(projection)
	return p
}

func (p *findResult) One(res interface{}) error {
	return p.find.One(res)
}

func (p *findResult) All(res interface{}) error {
	return p.find.All(res)
}

func (p *model) SetHandler(handler *Mgo) *model {
	p.handler = handler
	return p
}

func (p *model) Query() *Query {
	return p.handler.DB(p.db).C(p.collection)
}

func (p *model) Find(find interface{}) *findResult {
	return &findResult{find: p.handler.DB(p.db).C(p.collection).Find(find)}
}

func (p *model) Count(find interface{}) (int64, error) {
	return p.handler.DB(p.db).C(p.collection).CountDocuments(find)
}

func (p *model) Set(filter interface{}, update interface{}) *UpdateMany {
	return p.handler.DB(p.db).C(p.collection).UpdateMany(filter, bson.M{"$set": update})
}

func (p *model) Update(filter interface{}, update interface{}) *UpdateMany {
	return p.handler.DB(p.db).C(p.collection).UpdateMany(filter, update)
}

func (p *model) Insert(document ...interface{}) *InsertMany {
	return p.handler.DB(p.db).C(p.collection).InsertMany(document)
}

func (p *model) Delete(filter interface{}) *DeleteMany {
	return p.handler.DB(p.db).C(p.collection).DeleteMany(filter)
}

func (p *model) FindAndModify(filter interface{}, update interface{}) *FindOneAndUpdate {
	return p.handler.DB(p.db).C(p.collection).FindOneAndUpdate(filter, update)
}

func (p *model) Aggregate(pipeline interface{}) *aggregateResult {
	return &aggregateResult{aggregate: p.handler.DB(p.db).C(p.collection).Aggregate(pipeline)}
}
