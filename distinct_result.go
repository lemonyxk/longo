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

type DistinctResult struct {
	distinctResult *mongo.DistinctResult
}

func (sg *DistinctResult) Get(result interface{}) error {
	var err = sg.distinctResult.Decode(result)
	return err
}
