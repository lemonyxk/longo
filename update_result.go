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

type UpdateResult struct {
	updateResult *mongo.UpdateResult
	err          error
}

func (rs *UpdateResult) Result() *mongo.UpdateResult {
	return rs.updateResult
}

func (rs *UpdateResult) Error() error {
	return rs.err
}
