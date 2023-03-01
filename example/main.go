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
	"context"
	"log"

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
	var mgo, err = longo.NewClient().Connect(&longo.Config{Url: url})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	var test1 = longo.NewModel[Test2](context.Background(), mgo).DB("Test").C("User")
	var test2 = longo.NewModel[Test2](context.Background(), mgo).DB("Test").C("Test1")

	mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

		test1.Set(bson.M{"id": 2}, bson.M{"money": 100}).Context(sessionContext).Exec()

		test2.Set(bson.M{"id": 2}, bson.M{"money": 100}).Context(sessionContext).Exec()

		return nil
	})

	log.Println(test1.Set(bson.M{"id": 3}, bson.M{"money": 200}).Exec())

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

	// bucket, err := gridfs.NewBucket(mgo.RawClient().Database("test1"))
	// if err != nil {
	// 	panic(err)
	// }

	// stream, err := bucket.OpenUploadStreamWithID("abc1", "ok", &options.UploadOptions{Metadata: bson.M{"uuid": 111}})
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer stream.Close()
	//
	// var f, _ = os.Open("client.go")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// defer f.Close()
	//
	// io.Copy(stream, f)

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
