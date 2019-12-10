/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:35
**/

package longo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mgo struct {
	client *mongo.Client
	config Config
}

func (m *Mgo) RawClient() *mongo.Client {
	return m.client
}

func (m *Mgo) Ping() error {
	return m.client.Ping(context.Background(), ReadPreference.Primary)
}

func (m *Mgo) DB(db string) *DB {
	return &DB{client: m.client, db: db, config: m.config}
}

func (m *Mgo) Transaction(fn func(sessionContext mongo.SessionContext) error, opts ...*options.TransactionOptions) error {
	return m.client.UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		var err = sessionContext.StartTransaction(opts...)
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			var e = sessionContext.AbortTransaction(sessionContext)
			if e != nil {
				return e
			}
			return err
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}
