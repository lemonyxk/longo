/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2021-12-07 02:52
**/

package main

import (
	"log"

	"github.com/lemoyxk/longo/model"
)

func main() {
	log.SetFlags(0)
	model.Do("a", "", "/Users/lemo/lemo-hub/longo/example/model/test.go")
}
