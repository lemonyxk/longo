/**
* @program: mongo
*
* @description:
*
* @author: lemon
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

type BucketFilesList []*BucketFiles

type BucketFiles struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Length     int                `json:"length" bson:"length"`
	ChunkSize  int                `json:"chunkSize" bson:"chunkSize"`
	UploadDate time.Time          `json:"uploadDate" bson:"uploadDate"`
	Filename   string             `json:"filename" bson:"filename"`
	Metadata   interface{}        `json:"metadata" bson:"metadata"`
}

type BucketChunksList []*BucketChunks

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
	mgo             *Mgo
	databaseOptions []*options.DatabaseOptions
}

func (b *Bucket) New() (*gridfs.Bucket, error) {
	return gridfs.NewBucket(b.client.Database(b.db, b.databaseOptions...))
}

func (b *Bucket) NewFilesModel(ctx context.Context, opt ...*options.CollectionOptions) *Model[BucketFilesList, BucketFiles] {
	return NewModel[BucketFilesList](ctx, b.mgo).DB(b.db, b.databaseOptions...).C("fs.files", opt...)
}

func (b *Bucket) NewChunksModel(ctx context.Context, opt ...*options.CollectionOptions) *Model[BucketChunksList, BucketChunks] {
	return NewModel[BucketChunksList](ctx, b.mgo).DB(b.db, b.databaseOptions...).C("fs.chunks", opt...)
}

// func (b *Bucket) Context(ctx context.Context) *Bucket {
// 	b.context = ctx
// 	return b
// }

func (b *Bucket) UploadFile(id primitive.ObjectID, fileName string) *UploadFile {
	return NewUploadFile(b.client.Database(b.db, b.databaseOptions...), id, fileName)
}

func (b *Bucket) DeleteFile(id primitive.ObjectID) error {
	return NewDeleteFile(b.client.Database(b.db, b.databaseOptions...), id)
}

func (b *Bucket) DownloadFile(id primitive.ObjectID) *DownloadFile {
	return NewDownloadFile(b.client.Database(b.db, b.databaseOptions...), id)
}

func (b *Bucket) FindOne(filter interface{}) *FindOne {
	var bucket, err = gridfs.NewBucket(b.client.Database(b.db, b.databaseOptions...))
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
