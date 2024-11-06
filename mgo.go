/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:35
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mgo struct {
	client *mongo.Client
	config Config
	mux    sync.Mutex
}

func (m *Mgo) RawClient() *mongo.Client {
	return m.client
}

func (m *Mgo) Ping() error {
	return m.client.Ping(context.Background(), ReadPreference.Primary)
}

func (m *Mgo) DB(db string, opt ...*options.DatabaseOptions) *Database {
	return &Database{client: m.client, db: db, config: m.config, databaseOptions: opt}
}

func (m *Mgo) Bucket(db string, opt ...*options.DatabaseOptions) *Bucket {
	return &Bucket{client: m.client, db: db, config: m.config, mgo: m, databaseOptions: opt}
}

// TransactionWithLock one by one
func (m *Mgo) TransactionWithLock(fn func(handler *Mgo, sessionContext mongo.SessionContext) error, opts ...*options.TransactionOptions) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.Transaction(fn, opts...)
}

// Transaction can not create collection, so you have to create it before you run.
// maxTransactionLockRequestTimeoutMillis 5ms
func (m *Mgo) Transaction(fn func(handler *Mgo, sessionContext mongo.SessionContext) error, opts ...*options.TransactionOptions) error {
	//return m.client.UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
	//	var err = sessionContext.StartTransaction(opts...)
	//	if err != nil {
	//		return err
	//	}
	//	err = fn(m, sessionContext)
	//	if err != nil {
	//		var e = sessionContext.AbortTransaction(sessionContext)
	//		if e != nil {
	//			return e
	//		}
	//		return err
	//	}
	//	return sessionContext.CommitTransaction(sessionContext)
	//})

	var sess, err = m.client.StartSession()
	if err != nil {
		return err
	}

	defer sess.EndSession(context.Background())

	wc := writeconcern.Majority()
	txnWrite := options.Transaction().SetWriteConcern(wc)

	mc := readconcern.Majority()
	txnRead := options.Transaction().SetReadConcern(mc)

	ac := readpref.Primary()
	txnReadP := options.Transaction().SetReadPreference(ac)

	if len(opts) == 0 {
		opts = append(opts, txnWrite, txnRead, txnReadP)
	}

	_, err = sess.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		return nil, fn(m, ctx)
	}, opts...)

	return err
}
