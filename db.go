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
)

type DB struct {
	client *mongo.Client
	config Config
	db     string
}

func (db *DB) C(collection string) *Query {
	return &Query{client: db.client, db: db.db, config: db.config, collection: collection}
}

func (db *DB) Drop() error {
	return db.client.Database(db.db).Drop(context.Background())
}
