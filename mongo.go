/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 15:55
**/

package mongo

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// ReadPreference
// primary
// primarypreferred
// secondary
// secondarypreferred
// nearest

// ReadConcern
// local
// majority
// linearizable
// available
// snapshot

// WriteConcern
// w:0|1|majority
// j:false|true
// wtimeout

type Config struct {
	User           string
	Auth           string
	Hosts          []string
	Url            string
	ReadPreference *readpref.ReadPref
	ReadConcern    *readconcern.ReadConcern
	WriteConcern   *writeconcern.WriteConcern
}

type Client struct {
	config Config
}

func (c *Client) SetReadPreference(readPreference string) {
	c.config.ReadPreference = NewReadPreference(readPreference)
}

func (c *Client) SetReadConcern(readConcern string) {
	c.config.ReadConcern = NewReadConcern(readConcern)
}

func (c *Client) SetWriteConcern(writeConcern WriteConcern) {
	c.config.WriteConcern = NewWriteConcern(writeConcern)
}

func (c *Client) SetUrl(url string) {
	c.config.Url = url
}

func (c *Client) init(config *Config) {

	if config != nil {
		c.config = *config
	}

	if c.config.Url == "" {
		if len(c.config.Hosts) == 0 {
			c.config.Hosts = []string{"127.0.0.1:27017"}
		}
		var hostsString = strings.Join(c.config.Hosts, ",")
		if c.config.User == "" || c.config.Auth == "" {
			c.config.Url = "mongodb://" + hostsString
		} else {
			c.config.Url = "mongodb://" + c.config.User + ":" + c.config.Auth + "@" + hostsString
		}
	}

	if c.config.ReadPreference == nil {
		c.config.ReadPreference = ReadPreference.Primary
	}

	if c.config.ReadConcern == nil {
		c.config.ReadConcern = ReadConcern.Local
	}

	if c.config.WriteConcern == nil {
		c.config.WriteConcern = NewWriteConcern(WriteConcern{W: 1, J: false, Wtimeout: 3 * time.Second})
	}
}

func NewReadPreference(readPreference string) *readpref.ReadPref {
	rf, _ := readpref.ModeFromString(readPreference)
	rp, _ := readpref.New(rf)
	return rp
}

func NewReadConcern(readConcern string) *readconcern.ReadConcern {
	return readconcern.New(readconcern.Level(readConcern))
}

func NewWriteConcern(writeConcern WriteConcern) *writeconcern.WriteConcern {
	var opts []writeconcern.Option
	opts = append(opts, writeconcern.W(writeConcern.W))
	opts = append(opts, writeconcern.J(writeConcern.J))
	opts = append(opts, writeconcern.WTimeout(writeConcern.Wtimeout))
	return writeconcern.New(opts...)
}

func (c *Client) Connect(config *Config) (*Mgo, error) {

	if config == nil {
		config = &Config{}
	}

	c.init(config)

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.config.Url), &options.ClientOptions{
		ReadPreference: c.config.ReadPreference, // default is Primary
		ReadConcern:    c.config.ReadConcern,    // default is local
		WriteConcern:   c.config.WriteConcern,   // default is w:1 j:false wTimeout:when w > 1
	})

	if err != nil {
		return nil, err
	}

	return &Mgo{client: client, config: c.config}, nil
}

func NewMongoClient() *Client {
	return &Client{}
}

type Mgo struct {
	client *mongo.Client
	config Config
}

func (m *Mgo) RawClient() *mongo.Client {
	return m.client
}

func (m *Mgo) Ping() error {
	return m.client.Ping(context.Background(), ReadPreference.Primary)
}

func (m *Mgo) DB(db string) *DB {
	return &DB{client: m.client, db: db, config: m.config}
}

type DB struct {
	client *mongo.Client
	config Config
	db     string
}

func (db *DB) Transaction(fn func(sessionContext mongo.SessionContext) error, opts ...*options.TransactionOptions) {
	_ = db.client.UseSession(context.Background(), func(sessionContext mongo.SessionContext) error {
		var err = sessionContext.StartTransaction(opts...)
		if err != nil {
			return err
		}

		err = fn(sessionContext)
		if err != nil {
			return sessionContext.AbortTransaction(sessionContext)
		}

		return sessionContext.CommitTransaction(sessionContext)
	})
}

func (db *DB) C(collection string) *Query {
	return &Query{client: db.client, db: db.db, config: db.config, collection: collection}
}

type Query struct {
	client     *mongo.Client
	config     Config
	db         string
	collection string
}

func (q *Query) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	return q.client.Database(q.db).Collection(q.collection).Clone(opts...)
}

func (q *Query) Name() string {
	return q.client.Database(q.db).Collection(q.collection).Name()
}

func (q *Query) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertOne(context.Background(), document, opts...)
}

func (q *Query) InsertMany(document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return q.client.Database(q.db).Collection(q.collection).InsertMany(context.Background(), document, opts...)
}

func (q *Query) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteOne(context.Background(), filter, opts...)
}

func (q *Query) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return q.client.Database(q.db).Collection(q.collection).DeleteMany(context.Background(), filter, opts...)
}

func (q *Query) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateOne(context.Background(), filter, update, opts...)
}

func (q *Query) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).UpdateMany(context.Background(), filter, update, opts...)
}

func (q *Query) ReplaceOne(filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return q.client.Database(q.db).Collection(q.collection).ReplaceOne(context.Background(), filter, update, opts...)
}

func (q *Query) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) *MultiResult {
	cursor, err := q.client.Database(q.db).Collection(q.collection).Aggregate(context.Background(), pipeline, opts...)
	return &MultiResult{cursor: cursor, err: err}
}

type MultiResult struct {
	cursor *mongo.Cursor
	err    error
}

func (ag *MultiResult) All(result interface{}) error {
	if ag.err != nil {
		return ag.err
	}
	return ag.cursor.All(context.Background(), result)
}

// func (ag *MultiResult) RefAll(result interface{}) error {
//
// 	if ag.err != nil {
// 		return ag.err
// 	}
//
// 	var ctx = context.Background()
// 	defer func() { _ = ag.cursor.Close(ctx) }()
//
// 	refResult := reflect.ValueOf(result)
// 	if refResult.Kind() != reflect.Ptr || refResult.Elem().Kind() != reflect.Slice {
// 		return errors.New("result argument must be a slice address")
// 	}
//
// 	refSlice := refResult.Elem()
// 	refSlice = refSlice.Slice(0, 0)
// 	refElm := refSlice.Type().Elem()
// 	for ag.cursor.Next(ctx) {
// 		refM := reflect.New(refElm)
// 		_ = ag.cursor.Decode(refM.Interface())
// 		refSlice = reflect.Append(refSlice, refM.Elem())
// 	}
//
// 	refResult.Elem().Set(refSlice.Slice(0, refSlice.Len()))
// 	return nil
// }

// func (ag *MultiResult) Err(result interface{}) error {
// 	return ag.cursor.Err()
// }
//
// func (ag *MultiResult) ID(result interface{}) int64 {
// 	return ag.cursor.ID()
// }

type SingleResult struct {
	singleResult *mongo.SingleResult
}

func (sg *SingleResult) Get(result interface{}) error {
	return sg.singleResult.Decode(result)
}

// func (sg *SingleResult) Err(result interface{}) error {
// 	return sg.singleResult.Err()
// }

type ChangeStreamResult struct {
	changeStream *mongo.ChangeStream
	err          error
}

func (ct *ChangeStreamResult) Get(result interface{}) error {
	if ct.err != nil {
		return ct.err
	}
	var ctx = context.Background()
	defer func() { _ = ct.changeStream.Close(ctx) }()
	for ct.changeStream.Next(ctx) {
		return ct.changeStream.Decode(result)
	}
	return nil
}

func (ct *ChangeStreamResult) ResumeToken(result interface{}) (bson.Raw, error) {
	if ct.err != nil {
		return nil, ct.err
	}
	return ct.changeStream.ResumeToken(), nil
}

func (q *Query) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).CountDocuments(context.Background(), filter, opts...)
}

func (q *Query) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return q.client.Database(q.db).Collection(q.collection).EstimatedDocumentCount(context.Background(), opts...)
}

func (q *Query) Distinct(fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return q.client.Database(q.db).Collection(q.collection).Distinct(context.Background(), fieldName, filter, opts...)
}

func (q *Query) Find(filter interface{}, opts ...*options.FindOptions) *MultiResult {
	cursor, err := q.client.Database(q.db).Collection(q.collection).Find(context.Background(), filter, opts...)
	return &MultiResult{cursor: cursor, err: err}
}

func (q *Query) FindOne(filter interface{}, opts ...*options.FindOneOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOne(context.Background(), filter, opts...)}
}

func (q *Query) FindOneAndDelete(filter interface{}, opts ...*options.FindOneAndDeleteOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndDelete(context.Background(), filter, opts...)}
}

func (q *Query) FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndReplace(context.Background(), filter, replacement, opts...)}
}

func (q *Query) FindOneAndUpdate(filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *SingleResult {
	return &SingleResult{singleResult: q.client.Database(q.db).Collection(q.collection).FindOneAndUpdate(context.Background(), filter, update, opts...)}
}

func (q *Query) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) *ChangeStreamResult {
	res, err := q.client.Database(q.db).Collection(q.collection).Watch(context.Background(), pipeline, opts...)
	return &ChangeStreamResult{changeStream: res, err: err}
}

func (q *Query) Indexes() *IndexView {
	return &IndexView{view: q.client.Database(q.db).Collection(q.collection).Indexes()}
}

func (q *Query) Drop() error {
	return q.client.Database(q.db).Collection(q.collection).Drop(context.Background())
}

type IndexView struct {
	view mongo.IndexView
}

func (iv *IndexView) List(opts ...*options.ListIndexesOptions) *MultiResult {
	cursor, err := iv.view.List(context.Background(), opts...)
	return &MultiResult{cursor: cursor, err: err}
}

func (iv *IndexView) DropAll(opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return iv.view.DropAll(context.Background(), opts...)
}

func (iv *IndexView) DropOne(name string, opts ...*options.DropIndexesOptions) (bson.Raw, error) {
	return iv.view.DropOne(context.Background(), name, opts...)
}

func (iv *IndexView) CreateMany(models []mongo.IndexModel, opts ...*options.CreateIndexesOptions) ([]string, error) {
	return iv.view.CreateMany(context.Background(), models, opts...)
}

func (iv *IndexView) CreateOne(model mongo.IndexModel, opts ...*options.CreateIndexesOptions) (string, error) {
	return iv.view.CreateOne(context.Background(), model, opts...)
}
