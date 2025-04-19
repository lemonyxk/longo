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
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (m *Mgo) DB(db string, opt ...options.Lister[options.DatabaseOptions]) *Database {
	return &Database{client: m.client, db: db, config: m.config, databaseOptions: opt}
}

func (m *Mgo) Bucket(db string, opt ...options.Lister[options.DatabaseOptions]) *Bucket {
	return &Bucket{client: m.client, db: db, config: m.config, mgo: m, databaseOptions: opt}
}

// TransactionWithLock one by one
func (m *Mgo) TransactionWithLock(fn func(handler *Mgo, ctx context.Context) error, opts ...options.Lister[options.TransactionOptions]) error {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.Transaction(fn, opts...)
}

// Transaction can not create collection, so you have to create it before you run.
// maxTransactionLockRequestTimeoutMillis 5ms
func (m *Mgo) Transaction(fn func(handler *Mgo, ctx context.Context) error, opts ...options.Lister[options.TransactionOptions]) error {
	var session, err = m.client.StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(context.Background())

	if len(opts) == 0 {
		opts = append(opts,
			options.Transaction().
				SetWriteConcern(writeconcern.Majority()).
				SetReadConcern(readconcern.Majority()).
				SetReadPreference(readpref.Primary()),
		)
	}

	var f = func(ctx context.Context) (interface{}, error) {
		return nil, fn(m, ctx)
	}

	_, err = session.WithTransaction(context.TODO(), f, opts...)

	return err
}
