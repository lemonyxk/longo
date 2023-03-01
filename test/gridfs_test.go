/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2023-03-01 22:25
**/

package test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/lemonyxk/longo"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var NewObjectID = longo.NewObjectID()

func Test_GridFS_Upload(t *testing.T) {

	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test"))
	assert.True(t, err == nil, err)

	file, err := bucket.OpenUploadStreamWithID(NewObjectID, "trans_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, file != nil, file)

	defer func() { _ = file.Close() }()

	f, err := os.Open("trans_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, f != nil, f)

	defer func() { _ = f.Close() }()

	_, err = io.Copy(file, f)
	assert.True(t, err == nil, err)
}

func Test_GridFS_Download(t *testing.T) {

	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test"))
	assert.True(t, err == nil, err)

	file, err := bucket.OpenDownloadStream(NewObjectID)
	assert.True(t, err == nil, err)
	assert.True(t, file != nil, file)

	defer func() { _ = file.Close() }()

	var f = bytes.NewBuffer(nil)
	_, err = io.Copy(f, file)
	assert.True(t, err == nil, err)

	f1, err := os.Open("trans_test.go")
	assert.True(t, err == nil, err)
	assert.True(t, f1 != nil, f1)

	bts, err := io.ReadAll(f1)
	assert.True(t, err == nil, err)
	assert.True(t, string(bts) == f.String(), f.String())
}

func Test_GridFS_Delete(t *testing.T) {
	var bucket, err = gridfs.NewBucket(mgo.RawClient().Database("Test"))
	assert.True(t, err == nil, err)

	err = bucket.Delete(NewObjectID)
	assert.True(t, err == nil, err)

	file, err := bucket.OpenDownloadStream(NewObjectID)
	assert.True(t, err != nil, err)
	assert.True(t, file == nil, file)
}
