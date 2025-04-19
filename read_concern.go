/**
* @program: lottery-server
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-25 21:04
**/

package longo

import (
	"go.mongodb.org/mongo-driver/v2/mongo/readconcern"
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
	ReadConcern.Local = NewReadConcern("Local")
	ReadConcern.Majority = NewReadConcern("Majority")
	ReadConcern.Linearizable = NewReadConcern("Linearizable")
	ReadConcern.Available = NewReadConcern("Available")
	ReadConcern.Snapshot = NewReadConcern("Snapshot")
}
