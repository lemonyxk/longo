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

type DB struct {
	client *mongo.Client
	config Config
	db     string
}

func (db *DB) C(collection string) *Collection {
	return &Collection{client: db.client, db: db.db, config: db.config, collection: collection}
}

func (db *DB) Drop() error {
	return db.client.Database(db.db).Drop(context.Background())
}

func (db *DB) RunCommand(command interface{}) *Command {
	return &Command{db: db.client.Database(db.db), command: command, option: &options.RunCmdOptions{}}
}
