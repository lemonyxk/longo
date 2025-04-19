/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-26 19:10
**/

package main

import (
	"github.com/lemonyxk/longo"
	"time"
)

type Test2 map[string]interface{}

type List []*Test2

func (l List) Len() int {
	return len(l)
}

func main() {

	//var filter = filter2.New()
	//
	//var data = &Test2{
	//	ID:    1,
	//	Money: 1.1,
	//	Test: &Test2{
	//		ID: 2,
	//	},
	//}
	//
	//data.Test.Test = data
	//
	//var res = filter.Zero(data)
	//
	//log.Println(res)

	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
	var mgo, err = longo.NewClient().Connect(&longo.Config{
		Url: url, WriteConcern: &longo.WriteConcern{W: -1, J: true, WTimeout: 10 * time.Second},
	})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	//call.Default.Database("*").Collection().Type().Watch(func(info call.Record) {
	//	bts, err := json.Marshal(info)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	fmt.Println(string(bts))
	//})

	//var start = time.Now()
	//
	//var res interface{}
	//err = mgo.DB("test").C("test").FindOneAndUpdate(bson.M{"_id": 1}, bson.M{"$set": bson.M{"name": "2"}}).Sort(bson.M{"tid": -1}).Exec(&res)
	//if err != nil {
	//	panic(err)
	//}
	//
	//log.Printf("time: %+v\n", time.Since(start))
	//log.Println(res)

	//for i := 0; i < len(res); i++ {
	//	log.Printf("%+v\n", res[i])
	//}

	//var test = longo.NewModel[[]*Test2](context.Background(), mgo).DB("test").C("test")
	//var test1 = longo.NewModel[[]*Test2](context.Background(), mgo).DB("test").C("test1")

	//log.Println(test.Insert(&Test2{"_id": 1, "add": 1}, &Test2{"_id": 2, "add": 1}).Option(options.InsertMany().SetOrdered(false)).Exec())

	//_ = test
	//
	//// test1
	//var url1 = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
	//mgo1, err := longo.NewClient().Connect(&longo.Config{Url: url1, WriteConcern: &longo.WriteConcern{W: -1, J: true, WTimeout: 10 * time.Second}})
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = mgo1.RawClient().Ping(nil, longo.ReadPreference.Primary)
	//if err != nil {
	//	panic(err)
	//}
	//
	//var test1 = longo.NewModel[[]*Test2](context.Background(), mgo1).DB("test").C("test1")
	//
	//_ = test1

	//res, err := test.Update(bson.M{"_id": 1}, bson.M{"$set": bson.M{"id": 1, "add": 1}}).Exec()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", res)
	//
	//a, err := test.Find(bson.M{"_id": 999}).All()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", a)

	//models := []mongo.WriteModel{
	//	mongo.NewUpdateOneModel().SetFilter(
	//		bson.M{
	//			"_id": `1111111`,
	//		},
	//	).SetUpdate(
	//		bson.M{
	//			"$set": bson.M{"id1": 1},
	//		},
	//	),
	//	mongo.NewUpdateOneModel().SetFilter(
	//		bson.M{
	//			"_id": `2222222`,
	//		},
	//	).SetUpdate(
	//		bson.M{
	//			"$set": bson.M{"id1": 2},
	//		},
	//	),
	//}
	//
	//res, err := test1.BulkWrite(models).Exec()
	//if err != nil {
	//	panic(err)
	//}
	//
	//log.Printf("%+v\n", res)

	//var wait sync.WaitGroup
	//
	//wait.Add(1500)
	//
	//var start = time.Now()
	//
	//go func() {
	//	for i := 0; i < 1500; i++ {
	//		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
	//			res1, err := test.Set(bson.M{"_id": 1}, bson.M{"add": i}).Context(sessionContext).Exec()
	//			if err != nil {
	//				return err
	//			}
	//
	//			if res1.ModifiedCount == 0 {
	//				return errors.New("modified count is 0")
	//			}
	//
	//			res, err := test1.Insert(&Test2{"id": i, "add": i}).Context(sessionContext).Exec()
	//			if err != nil {
	//				return err
	//			}
	//
	//			if res.InsertedIDs == nil {
	//				return errors.New("inserted ids is nil")
	//			}
	//
	//			return nil
	//		})
	//		if err != nil {
	//			log.Printf("1: %+v\n", err)
	//		}
	//		wait.Done()
	//	}
	//}()
	//
	//wait.Wait()
	//
	//log.Printf("time: %+v\n", time.Since(start))

	//var wait sync.WaitGroup
	//
	//wait.Add(2)
	//
	//go func() {
	//	var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
	//
	//		_, err := test.Find(bson.M{}).Context(sessionContext).All()
	//		if err != nil {
	//			return errors.New("repeatable read 1: " + err.Error())
	//		}
	//
	//		time.Sleep(time.Millisecond * 500)
	//
	//		_, err = test.Insert(&Test2{"id": 1, "add": 1}).Context(sessionContext).Exec()
	//		if err != nil {
	//			return errors.New("repeatable write 2: " + err.Error())
	//		}
	//
	//		return nil
	//	})
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	wait.Done()
	//}()
	//
	//go func() {
	//	time.Sleep(time.Millisecond * 200)
	//	var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
	//		_, err := test.Insert(&Test2{"id": 2, "add": 2}).Context(sessionContext).Exec()
	//		return err
	//	})
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	wait.Done()
	//}()
	//
	//wait.Wait()

	//
	//a, err := test2.FindOneAndUpdate(bson.M{"_id": 96, "id": 3}, bson.M{"$set": Test2{
	//	ID:    3,
	//	Money: 5454,
	//}}).Upsert().ReturnDocument().Exec()
	//log.Println(a, err)
	//
	//_,err = test2.FindOneAndUpdate(bson.M{"_id": 99}, bson.M{"$inc": bson.M{"a": 1}}).Upsert().Exec()
	//log.Println(mongo.ErrNoDocuments == err)

	// var test1 = longo.NewModel[Test2](context.Background(), mgo).DB("Test").C("User")
	// var test2 = longo.NewModel[Test2](context.Background(), mgo).DB("Test").C("Test1")
	//
	// mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
	//
	// 	test1.Set(bson.M{"id": 2}, bson.M{"money": 100}).Context(sessionContext).Exec()
	//
	// 	test2.Set(bson.M{"id": 2}, bson.M{"money": 100}).Context(sessionContext).Exec()
	//
	// 	return nil
	// })
	//
	// log.Println(test1.Set(bson.M{"id": 3}, bson.M{"money": 200}).Exec())

	// var test2 = longo.NewModel[Test2]().DB("Test").C("test2").SetHandler(mgo)
	// test2.FindOneAndUpdate(bson.M{"_id": 999}, bson.M{"$setOnInsert": bson.M{"xixi": 111}, "$inc": bson.M{"a": 1}}).Upsert().Exec()

	// mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
	// 	var test2 = longo.NewModel[Test2]("Test", "test2").Context(sessionContext).SetHandler(mgo)
	// 	var res = test2.FindOneAndUpdate(bson.M{"id": 2}, bson.M{"$inc": bson.M{"money": 1}}).Upsert().ReturnDocument().Exec()
	// 	log.Println(res)
	// 	return nil
	// 	return errors.New("1")
	// })

	//
	// res, err := test2.Find(bson.M{"id": 1111111111}).Exec()
	// log.Println(res, err)

	// tranIsolationRepeatableOutside()
	// tranIsolationRepeatable()
	// tranIsolationRepeatableOutsideWithWrite()
	// tranIsolationRepeatableWithWrite()

	// var model = longo.NewModel[Person]("Test", "Person").SetHandler(mgo)
	//
	// log.Println(model.CreateIndex())

	// var buf bytes.Buffer
	// dStream, err := bucket.DownloadToStream(`ok`, &buf)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("File size to download: %v\n", dStream)
	// fmt.Println(buf.String())

	// var a, _ = primitive.ObjectIDFromHex("63b1451d7f733ad2c0be5db9")
	// bucket.Delete(a)

	// var a = bucket.GetFilesCollection().FindOne(context.Background(), bson.M{"metadata.uuid": 1111})
	// var r bson.M
	// a.Decode(&r)
	// log.Println(r)

	// mgo.Bucket("test1").New()

	// var b,_ = mgo.Bucket("test1").New()
	// b.Delete("abc")
	// b.Delete("abc1")
	// b.Delete("heihei")

	// var files,err = mgo.Bucket("test1").NewFilesModel().Find(bson.M{}).All()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// for i := 0; i < len(files); i++ {
	// 	log.Printf("%+v\n",files[i])
	// }
	//
	// var id, _ = primitive.ObjectIDFromHex("63b26ff1ff871ccb3807d7d3")
	// chunks, err := mgo.Bucket("test1").NewChunksModel().FindByID(id)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// log.Printf("%+v\n", string(chunks.Data))

	// var id, _ = primitive.ObjectIDFromHex("63b13ca965c4b836937d5620")
	// var b,_ = mgo.Bucket("test1").New()
	// var a,_ = b.Find(bson.M{"_id":id})

	// var home, _ = os.UserHomeDir()
	// var f, err = os.OpenFile(filepath.Join(home, "a.test"), os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(filepath.Join(home, "a.test"))
	// var buf = &bytes.Buffer{}
	// mgo.Bucket("test1").DownloadFile(id).Read(f)
	// log.Println(buf.String())

	// var a, _ = mgo.Bucket("test1").NewFilesModel().Find(bson.M{}).All()
	// for i := 0; i < len(a); i++ {
	// 	log.Printf("%+v\n", a[i])
	// }
}
