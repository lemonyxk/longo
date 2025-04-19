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

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Database struct {
	client          *mongo.Client
	config          Config
	db              string
	databaseOptions []options.Lister[options.DatabaseOptions]
}

func (db *Database) C(collection string, opt ...options.Lister[options.CollectionOptions]) *Collection {
	return &Collection{client: db.client, db: db.db, config: db.config, collection: collection, collectionOptions: opt, databaseOptions: db.databaseOptions}
}

func (db *Database) Drop() error {
	return db.client.Database(db.db, db.databaseOptions...).Drop(context.Background())
}

func (db *Database) RunCommand(command interface{}, opt options.Lister[options.RunCmdOptions]) *Command {
	return &Command{db: db.client.Database(db.db, db.databaseOptions...), command: command, option: opt}
}
