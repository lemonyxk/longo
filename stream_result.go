/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:32
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ChangeStreamResult struct {
	changeStream *mongo.ChangeStream
	err          error
}

func (ct *ChangeStreamResult) Exec(result interface{}) error {
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

func (ct *ChangeStreamResult) ResumeToken() (bson.Raw, error) {
	if ct.err != nil {
		return nil, ct.err
	}
	return ct.changeStream.ResumeToken(), nil
}
