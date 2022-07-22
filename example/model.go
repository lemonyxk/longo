/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2022-07-22 22:11
**/

package main

type Person struct {
	Name    string `bson:"name" json:"name" index:"name,1"`
	Age     int    `bson:"age" json:"age" `
	Address string `bson:"address" json:"address" indexes:"address,-1,age,1"`
	Type    string `bson:"type" json:"type" `
	MongoID string `bson:"_id" json:"_id"`
}
