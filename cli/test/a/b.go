/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-07 03:57
**/

package a

import (
	"github.com/lemoyxk/longo/longo"
	"github.com/lemoyxk/longo/model"
)

type Name struct {
	MongoID string `bson:"_id"`
}

func (t *Name) Empty() bool {
	return t == nil || t.MongoID == ""
}

func NewNameModel() *NameModel {
	return &NameModel{
		Model: &model.Model{DB: "222", C: "111"},
	}
}

type NameModel struct {
	*model.Model
}

func (p *NameModel) Find(find interface{}) *nameResult {
	return &nameResult{FindResult: p.Model.Find(find)}
}

func (p *NameModel) SetHandler(handler *longo.Mgo) *NameModel {
	p.Handler = handler
	return p
}

type nameResult struct {
	*model.FindResult
}

func (p *nameResult) One() *Name {
	var res Name
	_ = p.Find.One(&res)
	return &res
}

func (p *nameResult) All() []*Name {
	var res []*Name
	_ = p.Find.All(&res)
	return res
}

func (p *nameResult) Sort(sort interface{}) *nameResult {
	p.Find.Sort(sort)
	return p
}

func (p *nameResult) Skip(skip int64) *nameResult {
	p.Find.Skip(skip)
	return p
}

func (p *nameResult) Limit(limit int64) *nameResult {
	p.Find.Limit(limit)
	return p
}

func (p *nameResult) Hit(res interface{}) *nameResult {
	p.Find.Hit(res)
	return p
}

func (p *nameResult) Projection(projection interface{}) *nameResult {
	p.Find.Projection(projection)
	return p
}
