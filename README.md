```go
/**
* @program: mongo
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-26 19:10
**/

package main

import (
	"github.com/lemonyxk/longo"
)

func main() {

	var url = "mongodb://root:1354243@127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019"
	var mgo, err = longo.NewClient().Connect(&longo.Config{Url: url})
	if err != nil {
		panic(err)
	}

	err = mgo.RawClient().Ping(nil, longo.ReadPreference.Primary)
	if err != nil {
		panic(err)
	}
}

```