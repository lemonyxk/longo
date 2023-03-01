/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2023-03-01 15:58
**/

package test

import (
	"context"
	"testing"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func Test_Model_Insert(t *testing.T) {

	var test = longo.NewModel[TestDB](context.Background(), mgo).DB("Test").C("test")

	_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
	assert.True(t, err == nil, err)

	testes, err := test.Find(bson.M{}).All()
	assert.True(t, err == nil, err)
	assert.True(t, len(testes) == 1, len(testes))
}
