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
	"errors"
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

type Money struct {
	Money int `bson:"money"`
}

func main() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	var result = &Money{}

	mgo.DB("Test").C("test").Find(bson.M{"id": 1}).One(result)

	log.Println(result)

	// "renameCollection":"Test.test1", "to":"Test.test2"
	var res interface{}
	err = mgo.DB("admin").RunCommand(bson.D{
		bson.E{Key: "renameCollection", Value: "Test.test1"},
		bson.E{Key: "to", Value: "Test.test2"},
	}).One(&res)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

		var err error

		var result = &Money{}

		var test = mgo.DB("Test").C("test")
		err = test.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).ReturnDocument().Do(&result)

		// time.Sleep(time.Second * 1)

		// log.Println(result, err)

		err = errors.New("1")

		test.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"aaa": 1}}).Context(sessionContext).ReturnDocument().Do(&result)

		return err
	})
}

func test() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	var mux sync.WaitGroup

	mux.Add(100)

	var ctx mongo.SessionContext

	for i := 0; i < 100; i++ {
		go func() {
			// Transaction can not create collection, so you have to create it before you run.
			// maxTransactionLockRequestTimeoutMillis 5ms
			var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

				ctx = sessionContext
				log.Println(ctx.ID().Validate())

				var err error

				var result struct {
					Money int `bson:"money"`
				}

				var test = mgo.DB("Test").C("test")
				err = test.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).ReturnDocument().Do(&result)

				var result1 struct {
					Money int `bson:"money"`
				}

				var test1 = mgo.DB("Test").C("test1")
				err = test1.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).ReturnDocument().Do(&result1)

				// if result.Money != result1.Money {
				// 	panic("error")
				// }
				//
				// log.Println(mgo == handler)

				log.Println(result.Money, result1.Money)

				return err
			})

			log.Println(err)

			mux.Done()
		}()
	}

	mux.Wait()
	var res *bson.M
	err = mgo.DB("Test").C("test").FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(ctx).ReturnDocument().Do(res)
	log.Println(err)
	log.Println(res)
	// animal.Addr = "hello"
	//
	// log.Println(mgo.DB("Test").C("test").Find(bson.M{}).One(&animal))
	//
	// log.Println(animal)
}
