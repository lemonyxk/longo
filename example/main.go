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
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/lemoyxk/longo"
)

type Example struct {
	Hello string `json:"hello"`
}

type Animal struct {
	Name string `json:"name" bson:"name"`
	Addr string `json:"addr" bson:"addr"`
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

	var mux sync.WaitGroup

	mux.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			// Transaction can not create collection, so you have to create it before you run.
			// maxTransactionLockRequestTimeoutMillis 5ms
			var err = mgo.Transaction(func(sessionContext mongo.SessionContext) error {

				var err error

				var result struct {
					Money int `bson:"money"`
				}

				var test = mgo.DB("Test").C("test").SetContext(sessionContext)
				err = test.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).ReturnDocument().Get(&result)

				var result1 struct {
					Money int `bson:"money"`
				}

				var test1 = mgo.DB("Test").C("test1").SetContext(sessionContext)
				err = test1.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).ReturnDocument().Get(&result1)

				if result.Money != result1.Money {
					panic("error")
				}

				return err
			})

			log.Println(err)

			mux.Done()
		}()
	}

	mux.Wait()

	// var res bson.M
	// err = mgo.DB("Test").C("test").FindOneAndUpdate(bson.M{"name": "a"}, bson.M{"$set": bson.M{"name": "a111123aaaa"}}).ReturnDocument().Get(&res)
	// log.Println(err)
	// log.Println(res)

	// animal.Addr = "hello"
	//
	// log.Println(mgo.DB("Test").C("test").Find(bson.M{}).One(&animal))
	//
	// log.Println(animal)
}
