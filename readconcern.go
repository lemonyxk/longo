/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 21:04
**/

package mongo

import (
	"go.mongodb.org/mongo-driver/mongo/readconcern"
)

// ReadConcern

var ReadConcern struct {
	Local        *readconcern.ReadConcern
	Majority     *readconcern.ReadConcern
	Linearizable *readconcern.ReadConcern
	Available    *readconcern.ReadConcern
	Snapshot     *readconcern.ReadConcern
}

func init() {
	ReadConcern.Local = NewReadConcern("local")
	ReadConcern.Majority = NewReadConcern("majority")
	ReadConcern.Linearizable = NewReadConcern("linearizable")
	ReadConcern.Available = NewReadConcern("available")
	ReadConcern.Snapshot = NewReadConcern("snapshot")
}
