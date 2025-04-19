/**
* @program: lottery-server
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-25 21:06
**/

package longo

import (
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// ReadPreference

var ReadPreference struct {
	Primary            *readpref.ReadPref
	PrimaryPreferred   *readpref.ReadPref
	Secondary          *readpref.ReadPref
	SecondaryPreferred *readpref.ReadPref
	Nearest            *readpref.ReadPref
}

func init() {
	ReadPreference.Primary = NewReadPreference("Primary")
	ReadPreference.PrimaryPreferred = NewReadPreference("PrimaryPreferred")
	ReadPreference.Secondary = NewReadPreference("Secondary")
	ReadPreference.SecondaryPreferred = NewReadPreference("SecondaryPreferred")
	ReadPreference.Nearest = NewReadPreference("Nearest")
}
