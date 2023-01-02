/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:35
**/

package longo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BucketFiles struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Length     int                `json:"length" bson:"length"`
	ChunkSize  int                `json:"chunkSize" bson:"chunkSize"`
	UploadDate time.Time          `json:"uploadDate" bson:"uploadDate"`
	Filename   string             `json:"filename" bson:"filename"`
	Metadata   interface{}        `json:"metadata" bson:"metadata"`
}

type BucketChunks struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	FilesID primitive.ObjectID `json:"files_id" bson:"files_id"`
	N       int                `json:"n" bson:"n"`
	Data    []byte             `json:"data" bson:"data"`
}

type Bucket struct {
	client *mongo.Client
	config Config
	db     string
	// context context.Context
	mgo *Mgo
}

func (b *Bucket) New() (*gridfs.Bucket, error) {
	return gridfs.NewBucket(b.client.Database(b.db))
}

func (b *Bucket) NewFilesModel() *Model[BucketFiles] {
	return NewModel[BucketFiles](b.db, "fs.files").SetHandler(b.mgo)
}

func (b *Bucket) NewChunksModel() *Model[BucketChunks] {
	return NewModel[BucketChunks](b.db, "fs.chunks").SetHandler(b.mgo)
}

// func (b *Bucket) Context(ctx context.Context) *Bucket {
// 	b.context = ctx
// 	return b
// }

func (b *Bucket) UploadFile(id primitive.ObjectID, fileName string) *UploadFile {
	return NewUploadFile(b.client.Database(b.db), id, fileName)
}

func (b *Bucket) DeleteFile(id primitive.ObjectID) error {
	return NewDeleteFile(b.client.Database(b.db), id)
}

func (b *Bucket) DownloadFile(id primitive.ObjectID) *DownloadFile {
	return NewDownloadFile(b.client.Database(b.db), id)
}

func (b *Bucket) FindOne(filter interface{}) *FindOne {
	var bucket, err = gridfs.NewBucket(b.client.Database(b.db))
	return &FindOne{bucket.GetFilesCollection(), &options.FindOneOptions{}, filter, context.Background(), err}
}

func (b *Bucket) FindByID(id primitive.ObjectID) *FindOne {
	var bucket, err = gridfs.NewBucket(b.client.Database(b.db))
	return &FindOne{bucket.GetFilesCollection(), &options.FindOneOptions{}, bson.M{"_id": id}, context.Background(), err}
}

func (b *Bucket) Find(filter interface{}) *Find {
	var bucket, err = gridfs.NewBucket(b.client.Database(b.db))
	return &Find{bucket.GetFilesCollection(), &options.FindOptions{}, filter, context.Background(), err}
}
