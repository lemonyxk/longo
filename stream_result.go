/**
* @program: mongo
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-28 15:32
**/

package mongo

import (
	"context"
	
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

