/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2023-03-01 15:58
**/

package test

import (
	"context"
	"testing"
	"time"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func Test_Model_Insert(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Insert")
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
	assert.True(t, err == nil, err)
}

func Test_Model_InsertMany(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_InsertMany")
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
}

func Test_Model_Find(t *testing.T) {
	// cuz need read after write,
	// so we need to set read preference to primary.
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Find", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{"id": 1}).One()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 1, a.Add)

	a, err = test.Find(bson.M{"id": 2}).One()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 2, a.Add)
}

func Test_Model_FindAll(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindAll", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Update", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	_, err = test.Set(bson.M{"id": 1}, bson.M{"add": 3}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Find(bson.M{"id": 1}).One()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)
}

func Test_Model_UpdateMany(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_UpdateMany", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Delete", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_DeleteMany", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Count", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Count(nil)
	assert.True(t, err == nil, err)
	assert.True(t, a == 2, a)
}

func Test_Model_CountBy(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_CountBy", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.Count(bson.M{"id": 1})
	assert.True(t, err == nil, err)
	assert.True(t, a == 1, a)
}

func Test_Model_Aggregate(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_Aggregate", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_AggregateOne", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	var a struct {
		Add int `bson:"add"`
	}
	err = test.Aggregate(bson.A{bson.M{"$match": bson.M{"id": 1}}}).One(&a)
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 1, a.Add)
}

func Test_Model_FindOneAndUpdate(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindOneAndUpdate", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindAndDelete", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindOneAndReplace", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindOneAndUpsert", &options.CollectionOptions{
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
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_FindOneAndReturnDocument", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})
	_, err := test.Insert(&TestDB{ID: 1, Add: 1}, &TestDB{ID: 2, Add: 2}).Exec()
	assert.True(t, err == nil, err)
	a, err := test.FindOneAndReplace(bson.M{"id": 1}, &TestDB{ID: 1, Add: 3}).ReturnDocument().Exec()
	assert.True(t, err == nil, err)
	assert.True(t, a.Add == 3, a.Add)
}

func Test_Model_CreateIndex(t *testing.T) {
	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Model_CreateIndex", &options.CollectionOptions{
		ReadPreference: longo.ReadPreference.Primary,
	})

	test.CreateIndex()

	time.Sleep(time.Second * 3)

	var result []longo.Index
	err := test.Collection().Indexes().List().All(context.Background(), &result)
	assert.True(t, err == nil, err)
	assert.True(t, len(result) == 2, len(result))
	assert.True(t, result[1].Key["id"] == 1, result[1].Key["id"])
	assert.True(t, result[1].Name == "id_1", result[1].Name)
}
