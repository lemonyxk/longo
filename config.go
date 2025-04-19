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
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Config struct {
	User           string
	Pass           string
	Hosts          []string
	Url            string
	TLS            bool
	Direct         bool
	ReadPreference *readpref.ReadPref
	ReadConcern    *readconcern.ReadConcern
	WriteConcern   *WriteConcern
	ConnectTimeout time.Duration
	Timeout        time.Duration
	Register       *bson.Registry

	// Messages are compressed when both parties enable network compression. Otherwise, messages between the parties are uncompressed.
	// If you specify multiple compressors, then the order in which you list the compressors matter as well as the communication initiator.
	// For example, if mongosh specifies the following network compressors zlib,snappy and the mongod specifies snappy,zlib, messages between mongosh and mongod uses zlib.
	//
	// If the parties do not share at least one common compressor, messages between the parties are uncompressed.
	// For example, if mongosh specifies the network compressor zlib and mongod specifies snappy, messages between mongosh and mongod are not compressed.
	Compressors []string
}
