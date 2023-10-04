/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:33
**/

package longo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	User           string
	Pass           string
	Hosts          []string
	Url            string
	ReadPreference *readpref.ReadPref
	ReadConcern    *readconcern.ReadConcern
	WriteConcern   *WriteConcern
	ConnectTimeout time.Duration
	Register       *bsoncodec.Registry
}
