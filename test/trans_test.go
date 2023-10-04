/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2023-03-01 21:09
**/

package test

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Test_Transaction_Success(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test1 = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_Success1")
	var test2 = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_Success2")

	time.AfterFunc(time.Millisecond*100, func() {
		var a, err = test1.Find(bson.M{"id": 1}).One()
		assert.True(t, err == nil, err)
		assert.True(t, a.Add != 1, a.Add)
		wait.Done()
	})

	time.AfterFunc(time.Millisecond*500, func() {
		var a, err = test2.Find(bson.M{"id": 1}).One()
		assert.True(t, err == nil, err)
		assert.True(t, a.Add == 1, a.Add)
		wait.Done()
	})

	var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
		_, err := test1.Insert(&TestDB{ID: 1, Add: 1}).Context(sessionContext).Exec()
		if err != nil {
			return err
		}

		time.Sleep(time.Millisecond * 150)

		_, err = test2.Insert(&TestDB{ID: 1, Add: 1}).Context(sessionContext).Exec()
		if err != nil {
			return err
		}

		return nil
	})

	assert.True(t, err == nil, err)

	wait.Wait()
}

func Test_Transaction_RepeatableOutsideWithRead(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_RepeatableOutsideWithRead")

	go func() {
		var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			res1, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 1: " + err.Error())
			}

			time.Sleep(time.Millisecond * 500)

			res2, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 2: " + err.Error())
			}

			if len(res1) != len(res2) {
				return errors.New("repeatable read 3")
			}

			return nil
		})

		assert.True(t, strings.HasPrefix(err.Error(), "repeatable read 2: "), err)

		wait.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		var _, err = test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
		assert.True(t, err == nil, err)

		wait.Done()
	}()

	wait.Wait()
}

func Test_Transaction_RepeatableWithRead(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_RepeatableWithRead")

	go func() {
		var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			res1, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 1: " + err.Error())
			}

			time.Sleep(time.Millisecond * 500)

			res2, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 2: " + err.Error())
			}

			if len(res1) != len(res2) {
				return errors.New("repeatable read 3")
			}

			return nil
		})

		assert.True(t, strings.HasPrefix(err.Error(), "repeatable read 2: "), err)

		wait.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)

		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Context(sessionContext).Exec()
			return err
		})

		assert.True(t, err == nil, err)

		wait.Done()
	}()

	wait.Wait()
}

func Test_Transaction_RepeatableOutsideWithWrite(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_RepeatableOutsideWithWrite")

	go func() {
		var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			_, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 1: " + err.Error())
			}

			time.Sleep(time.Millisecond * 500)

			_, err = test.Insert(&TestDB{ID: 1, Add: 1}).Context(sessionContext).Exec()
			if err != nil {
				return errors.New("repeatable write 2: " + err.Error())
			}

			return nil
		})

		assert.True(t, strings.HasPrefix(err.Error(), "repeatable write 2: "), err)

		wait.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		var _, err = test.Insert(&TestDB{ID: 2, Add: 2}).Exec()
		assert.True(t, err == nil, err)

		wait.Done()
	}()

	wait.Wait()
}

func Test_Transaction_RepeatableWithWrite(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_RepeatableWithWrite")

	go func() {
		var err = mgo.TransactionWithLock(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			_, err := test.Find(bson.M{}).Context(sessionContext).All()
			if err != nil {
				return errors.New("repeatable read 1: " + err.Error())
			}

			time.Sleep(time.Millisecond * 500)

			_, err = test.Insert(&TestDB{ID: 1, Add: 1}).Context(sessionContext).Exec()
			if err != nil {
				return errors.New("repeatable write 2: " + err.Error())
			}

			return nil
		})

		assert.True(t, strings.HasPrefix(err.Error(), "repeatable write 2: "), err)

		wait.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 200)
		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			_, err := test.Insert(&TestDB{ID: 2, Add: 2}).Context(sessionContext).Exec()
			return err
		})

		assert.True(t, err == nil, err)

		wait.Done()
	}()

	wait.Wait()
}

func Test_Transaction_Set(t *testing.T) {

	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_Set")

	_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
	assert.True(t, err == nil, err)

	go func() {
		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			time.Sleep(time.Millisecond * 100)

			c, err := test.Update(bson.M{"id": 1}, bson.M{"$set": bson.M{"add": 1}}).Context(sessionContext).Exec()
			if err != nil {
				return err
			}

			assert.True(t, c == nil, c)

			time.Sleep(time.Millisecond * 1000)

			return nil
		})

		assert.True(t, err != nil, err)
		wait.Done()
	}()

	go func() {
		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {

			c, err := test.Update(bson.M{"id": 1}, bson.M{"$set": bson.M{"add": 2}}).Context(sessionContext).Exec()
			if err != nil {
				return err
			}

			assert.True(t, c.ModifiedCount == 1, c)

			time.Sleep(time.Millisecond * 500)

			return nil
		})

		assert.True(t, err == nil, err)

		wait.Done()
	}()

	wait.Wait()
}

func Test_Transaction_Set_Wait(t *testing.T) {
	var wait sync.WaitGroup

	wait.Add(2)

	var test = longo.NewModel[[]*TestDB](context.Background(), mgo).DB("Test").C("Test_Transaction_RepeatableWithWrite")

	_, err := test.Insert(&TestDB{ID: 1, Add: 1}).Exec()
	assert.True(t, err == nil, err)

	go func() {
		var err = mgo.Transaction(func(handler *longo.Mgo, sessionContext mongo.SessionContext) error {
			time.Sleep(time.Millisecond * 500)
			_, err := test.Set(bson.M{"id": 1}, bson.M{"add": 2}).Context(sessionContext).Exec()
			if err != nil {
				return err
			}

			time.Sleep(time.Millisecond * 3000)

			return nil
		})

		assert.True(t, err == nil, err)

		wait.Done()
	}()

	go func() {
		time.Sleep(time.Millisecond * 1000)
		// will wait for transaction until timeout
		_, err := test.FindOneAndUpdate(bson.M{"id": 1, "add": 1}, bson.M{"$set": bson.M{"add": 2}}).Exec()
		assert.True(t, errors.Is(err, mongo.ErrNoDocuments), err)
		wait.Done()
	}()

	wait.Wait()
}
