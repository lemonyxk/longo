/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:34
**/

package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func NewMongoClient() *Client {
	return &Client{}
}

func NewReadPreference(readPreference string) *readpref.ReadPref {
	rf, _ := readpref.ModeFromString(readPreference)
	rp, _ := readpref.New(rf)
	return rp
}

func NewReadConcern(readConcern string) *readconcern.ReadConcern {
	return readconcern.New(readconcern.Level(readConcern))
}

func NewWriteConcern(writeConcern WriteConcern) *writeconcern.WriteConcern {
	var opts []writeconcern.Option
	opts = append(opts, writeconcern.W(writeConcern.W))
	opts = append(opts, writeconcern.J(writeConcern.J))
	opts = append(opts, writeconcern.WTimeout(writeConcern.Wtimeout))
	return writeconcern.New(opts...)
}
