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
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BucketFilesList []*BucketFiles

type BucketFiles struct {
	ID         bson.ObjectID `json:"_id" bson:"_id"`
	Length     int           `json:"length" bson:"length"`
	ChunkSize  int           `json:"chunkSize" bson:"chunkSize"`
	UploadDate time.Time     `json:"uploadDate" bson:"uploadDate"`
	Filename   string        `json:"filename" bson:"filename"`
	Metadata   interface{}   `json:"metadata" bson:"metadata"`
}

type BucketChunksList []*BucketChunks

type BucketChunks struct {
	ID      bson.ObjectID `json:"_id" bson:"_id"`
	N       int           `json:"n" bson:"n"`
	Data    []byte        `json:"data" bson:"data"`
	FilesID bson.ObjectID `json:"files_id" bson:"files_id"`
}

type Bucket struct {
	client          *mongo.Client
	config          Config
	db              string
	mgo             *Mgo
	databaseOptions []options.Lister[options.DatabaseOptions]
}

func (b *Bucket) New(opts ...options.Lister[options.BucketOptions]) *mongo.GridFSBucket {
	return b.client.Database(b.db, b.databaseOptions...).GridFSBucket(opts...)
}

func (b *Bucket) NewFilesModel(ctx context.Context, opt ...options.Lister[options.CollectionOptions]) *Model[BucketFilesList, BucketFiles] {
	return NewModel[BucketFilesList](ctx, b.mgo).DB(b.db, b.databaseOptions...).C("fs.files", opt...)
}

func (b *Bucket) NewChunksModel(ctx context.Context, opt ...options.Lister[options.CollectionOptions]) *Model[BucketChunksList, BucketChunks] {
	return NewModel[BucketChunksList](ctx, b.mgo).DB(b.db, b.databaseOptions...).C("fs.chunks", opt...)
}

func (b *Bucket) UploadFile(id bson.ObjectID, fileName string) *UploadFile {
	return NewUploadFile(b.client.Database(b.db, b.databaseOptions...), id, fileName)
}

func (b *Bucket) DeleteFile(id bson.ObjectID) error {
	return NewDeleteFile(b.client.Database(b.db, b.databaseOptions...), id)
}

func (b *Bucket) DownloadFile(id bson.ObjectID) *DownloadFile {
	return NewDownloadFile(b.client.Database(b.db, b.databaseOptions...), id)
}

func (b *Bucket) FindOne(filter interface{}) *FindOne {
	var bucket = b.client.Database(b.db, b.databaseOptions...).GridFSBucket()
	return &FindOne{bucket.GetFilesCollection(), options.FindOne(), filter, context.Background()}
}

func (b *Bucket) FindByID(id bson.ObjectID) *FindOne {
	var bucket = b.client.Database(b.db, b.databaseOptions...).GridFSBucket()
	return &FindOne{bucket.GetFilesCollection(), options.FindOne(), bson.M{"_id": id}, context.Background()}
}

func (b *Bucket) Find(filter interface{}) *Find {
	var bucket = b.client.Database(b.db, b.databaseOptions...).GridFSBucket()
	return &Find{bucket.GetFilesCollection(), options.Find(), filter, context.Background()}
}
