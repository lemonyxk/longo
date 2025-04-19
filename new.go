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
	"go.mongodb.org/mongo-driver/v2/bson"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo/writeconcern"
)

func NewClient() *Client {
	return &Client{}
}

func NewObjectID() bson.ObjectID {
	return bson.NewObjectID()
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
		wc.W = int(writeConcern.W)
	}

	wc.Journal = &writeConcern.J
	//wc.WTimeout = writeConcern.WTimeout

	return &wc
}
