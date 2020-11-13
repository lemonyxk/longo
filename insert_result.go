/**
* @program: longo
*
* @description:
*
* @author: user
*
* @create: 2020-11-13 15:30
**/

package longo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type InsertManyResult struct {
	insertManyResult *mongo.InsertManyResult
	err              error
}

func (imr *InsertManyResult) Result() *mongo.InsertManyResult {
	return imr.insertManyResult
}

func (imr *InsertManyResult) Error() error {
	return imr.err
}

type InsertOneResult struct {
	insertOneResult *mongo.InsertOneResult
	err             error
}

func (ior *InsertOneResult) Result() *mongo.InsertOneResult {
	return ior.insertOneResult
}

func (ior *InsertOneResult) Error() error {
	return ior.err
}
