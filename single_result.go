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
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SingleResult struct {
	singleResult *mongo.SingleResult
}

func (sg *SingleResult) Get(result interface{}) error {
	var err = sg.singleResult.Decode(result)
	return err
}
