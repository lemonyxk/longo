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

type DeleteResult struct {
	deleteResult *mongo.DeleteResult
	err          error
}

func (dr *DeleteResult) Result() *mongo.DeleteResult {
	return dr.deleteResult
}

func (dr *DeleteResult) Error() error {
	return dr.err
}
