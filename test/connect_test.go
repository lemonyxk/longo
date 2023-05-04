/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2023-03-01 15:43
**/

package test

import (
	"context"
	"testing"

	"github.com/lemonyxk/longo"
)

type TestDB struct {
	ID   int     `json:"id" bson:"id" index:"id_1"`
	Add  float64 `json:"add" bson:"add"`
}

func (t *TestDB) Empty() bool {
	return t == nil || t.ID == 0
}

var mgo *longo.Mgo

func connect() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	var err error
	mgo, err = longo.NewClient().Connect(&longo.Config{Url: url})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

}

func clean() {
	var err = mgo.RawClient().Database("Test").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}

func TestMain(t *testing.M) {

	connect()

	t.Run()

	clean()
}