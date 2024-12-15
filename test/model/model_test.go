/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2023-03-01 15:58
**/

package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestDB struct {
	ID  int     `json:"id" bson:"id" index:"id_1"`
	Add float64 `json:"add" bson:"add"`
}

func (t *TestDB) Empty() bool {
	return t == nil || t.ID == 0
}

var mgo *longo.Mgo

func Test_Connect(t *testing.T) {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	var err error
	mgo, err = longo.NewClient().Connect(&longo.Config{Url: url, WriteConcern: &longo.WriteConcern{
		W:        1,
		J:        false,
		WTimeout: time.Second * 5,
	}})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Database("Test_2").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}

func Test_Model_Insert(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Insert")
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
	assert.True(t, err == nil, err)
}

func Test_Model_InsertMany(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_InsertMany")
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
}

func Test_Model_Find(t *testing.T) {
	// cuz need read after write,
	// so we need to set read preference to primary.
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Find", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOne(bson.M{"id": 1}).Get()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 1, a.Add)

	a, err = test.FindOne(bson.M{"id": 2}).Get()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 2, a.Add)
}

func Test_Model_FindAll(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindAll", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{}).Sort(bson.M{"id": -1}).All()
	assert.True(t, err == nil, err)
	assert.True(t, len(a) == 2, len(a))
	assert.True(t, a[0].Add == 2, a[0].Add)
	assert.True(t, a[1].Add == 1, a[1].Add)
}

func Test_Model_Update(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Update", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	_, err = test.Set(bson.M{"id": 1}, bson.M{"add": 3}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOne(bson.M{"id": 1}).Get()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)
}

func Test_Model_UpdateMany(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_UpdateMany", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	_, err = test.Set(bson.M{}, bson.M{"add": 3}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{}).All()
	assert.True(t, err == nil, err)
	assert.True(t, len(a) == 2, len(a))
	assert.True(t, a[0].Add == 3, a[0].Add)
	assert.True(t, a[1].Add == 3, a[1].Add)
}

func Test_Model_Delete(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Delete", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	_, err = test.Delete(bson.M{"id": 1}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{}).All()
	assert.True(t, err == nil, err)
	assert.True(t, len(a) == 1, len(a))
	assert.True(t, a[0].Add == 2, a[0].Add)
}

func Test_Model_DeleteMany(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_DeleteMany", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	_, err = test.Delete(bson.M{}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{}).All()
	assert.True(t, err == nil, err)
	assert.True(t, len(a) == 0, len(a))
}

func Test_Model_Count(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Count", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{}).Count()
	assert.True(t, err == nil, err)
	assert.True(t, a == 2, a)
}

func Test_Model_CountBy(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_CountBy", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{"id": 1}).Count()
	assert.True(t, err == nil, err)
	assert.True(t, a == 1, a)
}

func Test_Model_Aggregate(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_Aggregate", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	var a []struct {
		Add int `bson:"add"`
	}
	err = test.Aggregate(bson.A{bson.M{"$match": bson.M{"id": 1}}}).All(&a)
	assert.True(t, err == nil, err)
	assert.True(t, len(a) == 1, len(a))
	assert.True(t, a[0].Add == 1, a[0].Add)
}

func Test_Model_AggregateOne(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_AggregateOne", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	var a []struct {
		Add int `bson:"add"`
	}
	err = test.Aggregate(bson.A{bson.M{"$match": bson.M{"id": 1}}}).All(&a)
	assert.True(t, err == nil, err)
	assert.True(t, a[0].Add == 1, a[0].Add)
}

func Test_Model_FindOneAndUpdate(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindOneAndUpdate", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndUpdate(bson.M{"id": 1}, bson.M{"$set": bson.M{"add": 3}}).ReturnDocument().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)

	a, err = test.FindOneAndUpdate(bson.M{"id": 4}, bson.M{"$set": bson.M{"add": 4}}).ReturnDocument().Upsert().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 4, a.Add)
	assert.True(t, a.ID == 4, a.ID)
}

func Test_Model_FindAndDelete(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindAndDelete", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndDelete(bson.M{"id": 1}).Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 1, a.Add)

	a, err = test.FindOneAndDelete(bson.M{"id": 4}).Exec()
	assert.True(t, err != nil, err)
	assert.True(t, a.Empty(), a)
}

func Test_Model_FindOneAndReplace(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindOneAndReplace", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndReplace(bson.M{"id": 1}, &TestDB{ID: 1, Add: 3}).ReturnDocument().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)

	a, err = test.FindOneAndReplace(bson.M{"id": 4}, &TestDB{ID: 4, Add: 4}).ReturnDocument().Upsert().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 4, a.Add)
	assert.True(t, a.ID == 4, a.ID)
}

func Test_Model_FindOneAndUpsert(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindOneAndUpsert", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndReplace(bson.M{"id": 4}, &TestDB{ID: 4, Add: 4}).ReturnDocument().Upsert().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 4, a.Add)
	assert.True(t, a.ID == 4, a.ID)
}

func Test_Model_FindOneAndReturnDocument(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_FindOneAndReturnDocument", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndReplace(bson.M{"id": 1}, &TestDB{ID: 1, Add: 3}).ReturnDocument().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)
}

func Test_Model_BulkInsert(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_BulkWrite", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})

	models := []mongo.WriteModel{
		mongo.NewInsertOneModel().SetDocument(
			bson.M{
				"_id": `1111111`,
			},
		),
		mongo.NewInsertOneModel().SetDocument(
			bson.M{
				"_id": `2222222`,
			},
		),
	}

	res, err := test.BulkWrite(models).Exec()
	assert.True(t, err == nil, err)
	assert.True(t, res.InsertedCount == 2, res)
}

func Test_Model_BulkUpdate(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_BulkWrite", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})

	models := []mongo.WriteModel{
		mongo.NewUpdateOneModel().SetFilter(
			bson.M{
				"_id": `1111111`,
			},
		).SetUpdate(
			bson.M{
				"$set": bson.M{"id": 1},
			},
		),
		mongo.NewUpdateOneModel().SetFilter(
			bson.M{
				"_id": `2222222`,
			},
		).SetUpdate(
			bson.M{
				"$set": bson.M{"id": 2},
			},
		),
	}

	res, err := test.BulkWrite(models).Exec()
	assert.True(t, err == nil, err)
	assert.True(t, res.ModifiedCount == 2, res)
}

func Test_Model_BulkDelete(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_BulkWrite", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})

	models := []mongo.WriteModel{
		mongo.NewDeleteOneModel().SetFilter(
			bson.M{
				"_id": `1111111`,
			},
		),
		mongo.NewDeleteOneModel().SetFilter(
			bson.M{
				"_id": `2222222`,
			},
		),
	}

	res, err := test.BulkWrite(models).Exec()
	assert.True(t, err == nil, err)
	assert.True(t, res.DeletedCount == 2, res)
}

func Test_Model_CreateIndex(t *testing.T) {
	time.Sleep(time.Second * 3)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test_2").C("Test_Model_CreateIndex", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})

	var err = test.CreateIndex()
	assert.True(t, err == nil, err)

	time.Sleep(time.Second * 3)

	var result []longo.Index
	err = test.Collection().Indexes().List().All(context.Background(), &result)
	assert.True(t, err == nil, err)
	assert.True(t, len(result) == 2, len(result))
	assert.True(t, result[1].Key["id"] == 1, result[1].Key["id"])
	assert.True(t, result[1].Name == "id_1", result[1].Name)
}

func Test_Clean(t *testing.T) {
	var err = mgo.RawClient().Database("Test_2").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}
