/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-26 19:10
**/

package main

import (
	"encoding/json"
	"log"

	"github.com/Lemo-yxk/longo"
)

type Example struct {
	Hello string `json:"hello"`
}

type Animal struct {
	Name string `json:"name" bson:"name"`
	Eye
	Hand
}

type Hand struct {
	Size int `json:"size" bson:"size"`
}

type Eye struct {
	Color string `json:"color" bson:"color"`
}

func main() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// Transaction can not create collection, so you have to create it before you run.
	// mgo.Transaction(func(sessionContext mongo.SessionContext) error {
	//
	// 	var err error
	//
	// 	_, err = mgo.DB("Test").C("test1").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
	//
	// 	_, err = mgo.DB("Test").C("test2").InsertOneWithSession(sessionContext, bson.M{"hello": "world"})
	//
	// 	return err
	// })

	var animal = Animal{}

	b, _ := json.Marshal(animal)
	log.Println(string(b))

	_, _ = mgo.DB("Test").C("test").InsertOne(animal)

}
