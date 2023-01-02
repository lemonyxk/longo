/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2023-01-02 10:18
**/

package longo

import (
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UploadFile struct {
	dataBase *mongo.Database
	fileName string
	option   *options.UploadOptions
	id       primitive.ObjectID
	// sessionContext context.Context
}

func NewUploadFile(dataBase *mongo.Database, id primitive.ObjectID, fileName string) *UploadFile {
	return &UploadFile{fileName: fileName, dataBase: dataBase, id: id}
}

func (u *UploadFile) MetaData(metadata interface{}) *UploadFile {
	u.option.Metadata = metadata
	return u
}

func (u *UploadFile) Write(reader io.Reader) error {
	var bucket, err = gridfs.NewBucket(u.dataBase)
	if err != nil {
		return err
	}

	st, err := bucket.OpenUploadStreamWithID(u.id, u.fileName, u.option)
	if err != nil {
		return err
	}

	defer func() { _ = st.Close() }()

	_, err = io.Copy(st, reader)
	if err != nil {
		return err
	}

	return nil
}

func NewDeleteFile(dataBase *mongo.Database, id primitive.ObjectID) error {
	var bucket, err = gridfs.NewBucket(dataBase)
	if err != nil {
		return err
	}
	return bucket.Delete(id)
}

type DownloadFile struct {
	dataBase *mongo.Database
	fileName string
	option   *options.UploadOptions
	id       primitive.ObjectID
	// sessionContext context.Context
}

func NewDownloadFile(dataBase *mongo.Database, id primitive.ObjectID) *DownloadFile {
	return &DownloadFile{dataBase: dataBase, id: id}
}

func (u *DownloadFile) Read(writer io.Writer) error {
	var bucket, err = gridfs.NewBucket(u.dataBase)
	if err != nil {
		return err
	}

	st, err := bucket.OpenDownloadStream(u.id)
	if err != nil {
		return err
	}

	defer func() { _ = st.Close() }()

	_, err = io.Copy(writer, st)
	if err != nil {
		return err
	}

	return nil
}
