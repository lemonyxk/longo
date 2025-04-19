/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2023-01-02 10:18
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"io"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UploadFile struct {
	dataBase *mongo.Database
	fileName string
	option   options.Lister[options.GridFSUploadOptions]
	id       bson.ObjectID
}

func NewUploadFile(dataBase *mongo.Database, id bson.ObjectID, fileName string) *UploadFile {
	return &UploadFile{fileName: fileName, dataBase: dataBase, option: options.GridFSUpload(), id: id}
}

func (u *UploadFile) Option(opt options.Lister[options.GridFSUploadOptions]) *UploadFile {
	u.option = opt
	return u
}

func (u *UploadFile) Write(reader io.Reader) error {
	var bucket = u.dataBase.GridFSBucket()

	st, err := bucket.OpenUploadStreamWithID(context.Background(), u.id, u.fileName, u.option)
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

func NewDeleteFile(dataBase *mongo.Database, id bson.ObjectID) error {
	var bucket = dataBase.GridFSBucket()
	return bucket.Delete(context.Background(), id)
}

type DownloadFile struct {
	dataBase *mongo.Database
	fileName string
	option   *options.GridFSUploadOptions
	id       bson.ObjectID
}

func NewDownloadFile(dataBase *mongo.Database, id bson.ObjectID) *DownloadFile {
	return &DownloadFile{dataBase: dataBase, id: id}
}

func (u *DownloadFile) Read(writer io.Writer) error {
	var bucket = u.dataBase.GridFSBucket()

	st, err := bucket.OpenDownloadStream(context.Background(), u.id)
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
