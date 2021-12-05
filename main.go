/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-06 00:56
**/

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/lemoyxk/longo/model"
)

func init() {
	log.SetFlags(0)
}

func main() {

	_ = flag.Bool("x", false, "fixed the rule")
	file := flag.String("f", "", "file")
	db := flag.String("db", "", "file")
	c := flag.String("c", "", "file")

	flag.Parse()

	var filePath = *file
	var dbName = *db
	var collectionName = *c

	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		println(err)
		return
	}

	f, err := os.Stat(absFilePath)
	if os.IsNotExist(err) {
		println(absFilePath, "not exists")
		return
	}

	if f.IsDir() {
		println(absFilePath, "is a dir")
		return
	}

	model.Do(dbName, collectionName, absFilePath)

	println("model build success")
}
