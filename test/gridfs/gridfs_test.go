/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2023-03-01 22:25
**/

package gridfs

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type TestDB struct {
	ID  int     `json:"id" bson:"id" index:"id_1"`
	Add float64 `json:"add" bson:"add"`
}

func (t *TestDB) Empty() bool {
	return t == nil || t.ID == 0
}

var mgo *longo.Mgo

func Test_Connect(t *testing.T) {
	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"

	var err error
	mgo, err = longo.NewClient().Connect(&longo.Config{Url: url, WriteConcern: &longo.WriteConcern{
		W:        1,
		J:        false,
		WTimeout: time.Second * 5,
	}})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}
}

var NewObjectID = longo.NewObjectID()

func Test_GridFS_Upload(t *testing.T) {

	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test_1"))
	assert.True(t, err == nil, err)

	file, err := bucket.OpenUploadStreamWithID(NewObjectID, "gridfs_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, file != nil, file)

	defer func() { _ = file.Close() }()

	f, err := os.Open("gridfs_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, f != nil, f)

	defer func() { _ = f.Close() }()

	_, err = io.Copy(file, f)
	assert.True(t, err == nil, err)
}

func Test_GridFS_Download(t *testing.T) {

	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test_1"))
	assert.True(t, err == nil, err)

	file, err := bucket.OpenDownloadStream(NewObjectID)
	assert.True(t, err == nil, err)
	assert.True(t, file != nil, file)

	defer func() { _ = file.Close() }()

	var f = bytes.NewBuffer(nil)
	_, err = io.Copy(f, file)
	assert.True(t, err == nil, err)

	f1, err := os.Open("gridfs_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, f1 != nil, f1)

	bts, err := io.ReadAll(f1)
	assert.True(t, err == nil, err)
	assert.True(t, string(bts) == f.String(), f.String())
}

func Test_GridFS_Delete(t *testing.T) {
	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test_1"))
	assert.True(t, err == nil, err)

	err = bucket.Delete(NewObjectID)
	assert.True(t, err == nil, err)

	file, err := bucket.OpenDownloadStream(NewObjectID)
	assert.True(t, err != nil, err)
	assert.True(t, file == nil, file)
}

func Test_Clean(t *testing.T) {
	var err = mgo.RawClient().Database("Test_1").Drop(context.Background())
	if err != nil {
		panic(err)
	}
}