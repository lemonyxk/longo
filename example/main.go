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
	"reflect"
	"time"

	"github.com/lemonyxk/longo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Test2 struct {
	ID    int     `json:"id" bson:"id"`
	Money float64 `json:"money" bson:"money"`
}

func main() {

	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})
	//
	// err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// var test2 = longo.NewModel[Test2]("Test", "test2").SetHandler(mgo)
	//
	// res, err := test2.Find(bson.M{"id": 1111111111}).One()
	// log.Println(res, err == nil)

	// tranIsolationRepeatableOutside()
	// tranIsolationRepeatable()
	// tranIsolationRepeatableOutsideWithWrite()
	// tranIsolationRepeatableWithWrite()

	var model = longo.NewModel[Person]("Test", "Person").SetHandler(mgo)

	log.Println(model.CreateIndex())
}

// MONGO与MYSQL的事务区别在于
// 如果在可重复读的事务隔离性下
// 两次读取外界干扰的情况下,如果继续写 (A:R-B:W-A:W)
// MONGO默认使用悲观锁解决,即返回错误
// MYSQL需要自己update check来检测是否是原来的值

// 1.在事务提交期间,外部读取操作可能会尝试读取将被事务修改的相同文档,如果事务写入多个分片,则在跨分片提交尝试期间,
// 使用读取关注"snapshot"或"linearizable",或因果一致会话的一部分(即包括afterClusterTime)的外部读取等待事务的所有写入可见,
// 使用其他读取问题的外部读取不会等待事务的所有写入可见,而是读取可用文档的事务前版本.
// 2.如果包含读取操作,则ReadPreference必须选用Primary.
// 3.ReadConcern和WriteConcern根据需求选用,路由到同一节点即可.(或者都选择Majority)
// 4.事务采用表锁.

// 验证MONGO事务隔离性
// 非事务可重复读
// 即外界非事务改变不会干扰事务中的读取
func tranIsolationRepeatableOutside() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// 事务读取
	// 外界写入
	// 事务读取 NO CHANGE

	go func() {
		err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			var err error

			var res1 []Test2

			var res2 []Test2

			log.Println("tran start")

			var test2 = mgo.DB("Test").C("test2")

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res1)

			log.Println("tran:", res1)

			time.Sleep(time.Second * 2)

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res2)

			log.Println("tran:", res2)

			if !reflect.DeepEqual(res1, res2) {
				err = errors.New("inconsistent data")
			}

			return err
		})

		log.Println(err)
	}()

	go func() {
		time.Sleep(time.Second * 1)

		var res Test2
		var test2 = mgo.DB("Test").C("test2")

		var err = test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).
			ReturnDocument().Do(&res)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("outside:", res)
		}
	}()

	select {}
}

// 验证MONGO事务隔离性
// 事务可重复读
// 即外界事务改变不会干扰事务中的读取
// maxTransactionLockRequestTimeoutMillis 事务之间等待锁最长5MS
func tranIsolationRepeatable() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// 事务1读取
	// 事务2写入
	// 事务1读取 NO CHANGE

	go func() {
		err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			var err error

			var res1 []Test2

			var res2 []Test2

			log.Println("tran1 start")

			var test2 = mgo.DB("Test").C("test2")

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res1)

			log.Println("tran1:", res1)

			time.Sleep(time.Millisecond * 2)

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res2)

			log.Println("tran1:", res2)

			if !reflect.DeepEqual(res1, res2) {
				err = errors.New("inconsistent data")
			}

			return err
		})

		log.Println(err)
	}()

	go func() {
		time.Sleep(time.Millisecond)

		err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			var err error

			var res Test2

			log.Println("tran2 start")

			var test2 = mgo.DB("Test").C("test2")

			err = test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).
				ReturnDocument().Do(&res)

			log.Println("tran2:", res)

			return err
		})

		log.Println(err)
	}()

	select {}
}

// 验证MONGO事务隔离性
// 非事务可重复读再写
// 即外界非事务改变不会干扰事务中的读取之后再写入
// MONGO:事务查询之后外界有写入,此时如果事务继续写,则会返回错误
//
//	如果外界后写入则会等待事务完成
//
// MYSQL:事务查询之后外界有写入,此时如果事务继续写,则需要自己使用悲观锁(for update)或者乐观锁(update check)来解决外界非事务更改
//
//	如果外界后写入则会等待事务完成
func tranIsolationRepeatableOutsideWithWrite() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// 事务读取
	// 外界写入
	// 事务写入 ERROR

	go func() {

		err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			// var res Test2

			// var res1 []Test2

			// var res2 []Test2

			log.Println("tran start")

			var test2 = longo.NewModel[Test2]("Test", "test2").Context(sessionContext).SetHandler(handler)

			res1, err := test2.Find(bson.M{}).All()

			log.Println("tran:", res1)

			time.Sleep(time.Second * 2)

			res2, err := test2.Find(bson.M{}).All()

			log.Println("tran:", res2)

			if !reflect.DeepEqual(res1, res2) {
				err = errors.New("inconsistent data")
			}

			// will get a error!!!
			res := test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).ReturnDocument().Do()

			log.Println("tran:", res.Result())

			err = res.Error()

			return err
		})

		log.Println(err)
	}()

	go func() {
		time.Sleep(time.Second)
		var res Test2
		var test2 = mgo.DB("Test").C("test2")

		var err = test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).
			ReturnDocument().Do(&res)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("outside:", res)
		}
	}()

	select {}
}

// 验证MONGO事务隔离性
// 事务可重复读再写
// MONGO:事务查询之后其他事务有写入,此时如果事务继续写,则会返回错误
//
//	如果其他事务后写入提交则会等待事务完成
//
// MYSQL:事务查询之后其他事务有写入并提交,此时如果该事务继续写,则需要自己使用悲观锁(for update)或者乐观锁(update check)来解决外界事务更改
//
//	如果其他事务后写入提交则会等待该事务完成
//
// maxTransactionLockRequestTimeoutMillis 事务之间等待锁最长5MS
func tranIsolationRepeatableWithWrite() {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	mgo, _ := longo.NewClient().Connect(&longo.Config{Url: url})

	err := mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	// 事务1读取
	// 事务2读取
	// 事务2写入
	// 事务2提交
	// 事务1写入 ERROR

	go func() {
		err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			var err error

			var res Test2

			var res1 []Test2

			log.Println("tran1 start")

			var test2 = mgo.DB("Test").C("test2")

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res1)

			log.Println("tran1:", res1)

			time.Sleep(time.Millisecond * 2)

			// will get a error!!!
			err = test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).
				ReturnDocument().Do(&res)

			log.Println("tran1:", res)

			return err
		})

		log.Println(err)
	}()

	go func() {
		time.Sleep(time.Millisecond * 1)

		err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			var err error

			var res Test2

			var res1 []Test2

			log.Println("tran2 start")

			var test2 = mgo.DB("Test").C("test2")

			err = test2.Find(bson.M{}).Context(sessionContext).All(&res1)

			log.Println("tran2:", res1)

			err = test2.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$inc": bson.M{"money": 1}}).Context(sessionContext).
				ReturnDocument().Do(&res)

			log.Println("tran2:", res)

			return err
		})

		log.Println(err)
	}()

	select {}
}
