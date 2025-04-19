/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-28 15:28
**/

package longo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MultiResult struct {
	cursor *mongo.Cursor
	err    error
}

func (ag *MultiResult) All(sessionContext context.Context, result interface{}) error {
	if ag.err != nil {
		return ag.err
	}

	if ag.cursor == nil {
		return fmt.Errorf("cursor is nil")
	}

	if ag.cursor.Err() != nil {
		return ag.cursor.Err()
	}

	return ag.cursor.All(sessionContext, result)
}
