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
	Name        string `bson:"name" json:"name" index:"name_1,unique"`
	Age         int    `bson:"age" json:"age" index:"age_1_name_-1"`
	Address     string `bson:"address" json:"address" indexes:"address_-1_age_1,unique"`
	Type        string `bson:"type" json:"type"`
	MongoID     string `bson:"_id" json:"_id"`
	ProxyUserID string `bson:"proxy_user_id" json:"proxy_user_id" indexes:"proxy_user_id_1_type_-1,unique"`
}
