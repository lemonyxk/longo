/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-06 01:08
**/

package test

import "github.com/lemoyxk/longo"

type Test struct {
	MongoID int `json:"_id" bson:"_id" mapstructure:"_id"`
	ID      int `json:"id" bson:"id" mapstructure:"id"`
}

func (t *Test) Empty() bool {
	return t == nil || t.MongoID == 0
}

func NewTestModel() *TestModel {
	return &TestModel{
		Model: &longo.Model{DB: "222", C: "111"},
	}
}

type TestModel struct {
	*longo.Model
}

func (p *TestModel) Find(find interface{}) *testResult {
	return &testResult{FindResult: p.Model.Find(find)}
}

func (p *TestModel) SetHandler(handler *longo.Mgo) *TestModel {
	p.Handler = handler
	return p
}

type testResult struct {
	*longo.FindResult
}

func (p *testResult) One() *Test {
	var res Test
	_ = p.Find.One(&res)
	return &res
}

func (p *testResult) All() []*Test {
	var res []*Test
	_ = p.Find.All(&res)
	return res
}

func (p *testResult) Sort(sort interface{}) *testResult {
	p.Find.Sort(sort)
	return p
}

func (p *testResult) Skip(skip int64) *testResult {
	p.Find.Skip(skip)
	return p
}

func (p *testResult) Limit(limit int64) *testResult {
	p.Find.Limit(limit)
	return p
}

func (p *testResult) Hit(res interface{}) *testResult {
	p.Find.Hit(res)
	return p
}

func (p *testResult) Projection(projection interface{}) *testResult {
	p.Find.Projection(projection)
	return p
}
