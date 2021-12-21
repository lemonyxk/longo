/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:34
**/

package longo

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func NewClient() *Client {
	return &Client{}
}

func NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func NewReadPreference(readPreference string) *readpref.ReadPref {
	rf, _ := readpref.ModeFromString(strings.ToLower(readPreference))
	rp, _ := readpref.New(rf)
	return rp
}

func NewReadConcern(readConcern string) *readconcern.ReadConcern {
	return readconcern.New(readconcern.Level(strings.ToLower(readConcern)))
}

func NewWriteConcern(writeConcern WriteConcern) *writeconcern.WriteConcern {
	var opts []writeconcern.Option
	if writeConcern.W == -1 {
		opts = append(opts, writeconcern.WMajority())
	} else {
		opts = append(opts, writeconcern.W(writeConcern.W))
	}
	opts = append(opts, writeconcern.J(writeConcern.J))
	opts = append(opts, writeconcern.WTimeout(writeConcern.WTimeout))
	return writeconcern.New(opts...)
}
