/**
* @program: mongo
*
* @description:
*
* @author: lemon
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
	return &readconcern.ReadConcern{Level: strings.ToLower(readConcern)}
}

func NewWriteConcern(writeConcern WriteConcern) *writeconcern.WriteConcern {
	var wc = writeconcern.WriteConcern{}

	if writeConcern.W == Majority {
		wc.W = "majority"
	} else {
		wc.W = writeConcern.W
	}

	wc.Journal = &writeConcern.J
	wc.WTimeout = writeConcern.WTimeout

	return &wc
}
