/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2021-10-06 16:18
**/

package longo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Command struct {
	command interface{}
	option  options.Lister[options.RunCmdOptions]
	db      *mongo.Database
}

func (c *Command) All(result interface{}) error {
	cursor, err := c.db.RunCommandCursor(context.Background(), c.command, c.option)
	var res = &MultiResult{cursor: cursor, err: err}
	return res.All(context.Background(), result)
}

func (c *Command) One(result interface{}) error {
	cursor := c.db.RunCommand(context.Background(), c.command, c.option)
	var res = &SingleResult{singleResult: cursor}
	return res.Get(result)
}
