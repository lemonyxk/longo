/**
* @program: longo
*
* @description:
*
* @author: lemon
*
* @create: 2022-07-24 18:51
**/

package longo

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"strings"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// in array
func inMapArray(m map[string]bool, src string) bool {
	var arr = strings.Split(src, ",")
	if len(arr) == 0 {
		return false
	}
	return m[arr[0]]
}

// parse index
func parseIndex(indexArr []string) []mongo.IndexModel {

	var res []mongo.IndexModel

	if len(indexArr) == 0 {
		return res
	}

	for i := 0; i < len(indexArr); i++ {
		var indexConfig = strings.Split(indexArr[i], ",")
		var indexName = getIndexName(indexConfig)
		if len(indexName) == 0 {
			continue
		}

		var im = mongo.IndexModel{Keys: bson.M{indexName[0].name: indexName[0].sort}, Options: options.Index().SetUnique(isUnique(indexConfig))}

		res = append(res, im)
	}

	return res
}

func parseIndexes(indexesArr []string) []mongo.IndexModel {

	var res []mongo.IndexModel

	if len(indexesArr) == 0 {
		return res
	}

	for i := 0; i < len(indexesArr); i++ {
		var indexConfig = strings.Split(indexesArr[i], ",")
		var indexName = getIndexesName(indexConfig)
		if len(indexName) == 0 {
			continue
		}

		var im = mongo.IndexModel{Keys: bson.D{}, Options: options.Index().SetUnique(isUnique(indexConfig))}

		var keys = bson.D{}
		for j := 0; j < len(indexName); j++ {
			keys = append(keys, bson.E{Key: indexName[j].name, Value: indexName[j].sort})
		}

		im.Keys = keys

		res = append(res, im)
	}

	return res
}

// is unique
func isUnique(indexConfig []string) bool {
	for i := 0; i < len(indexConfig); i++ {
		if indexConfig[i] == "unique" {
			return true
		}
	}
	return false
}

func getIndexName(indexConfig []string) []bsonE {
	var res []bsonE
	var indexName = indexConfig[0]
	if strings.HasSuffix(indexName, "_-1") {
		res = append(res, bsonE{name: indexName[:len(indexName)-3], sort: -1})
	}
	if strings.HasSuffix(indexName, "_1") {
		res = append(res, bsonE{name: indexName[:len(indexName)-2], sort: 1})
	}
	return res
}

// get multiple indexes
func getIndexesName(indexConfig []string) []bsonE {
	var indexName = indexConfig[0]

	var res []bsonE
	var s = 0
	for i := 0; i < len(indexName); i++ {
		if indexName[i] == '_' {
			if i+1 < len(indexName) {
				if indexName[i+1] == '1' {
					res = append(res, bsonE{name: indexName[s:i], sort: 1})
					i += 2
					s = i + 1
				} else if indexName[i+1] == '-' {
					if i+2 < len(indexName) {
						if indexName[i+2] == '1' {
							res = append(res, bsonE{name: indexName[s:i], sort: -1})
							i += 3
							s = i + 1
						}
					}
				}
			}
		}
	}
	return res
}

type Index struct {
	Name string         `bson:"name"`
	Key  map[string]int `bson:"key"`
}

type bsonE struct {
	name string
	sort int
}
