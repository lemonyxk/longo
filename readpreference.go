/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 21:06
**/

package lemongo

import (
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ReadPreference

var ReadPreference struct {
	Primary            *readpref.ReadPref
	Primarypreferred   *readpref.ReadPref
	Secondary          *readpref.ReadPref
	Secondarypreferred *readpref.ReadPref
	Nearest            *readpref.ReadPref
}

func init() {
	ReadPreference.Primary = NewReadPreference("primary")
	ReadPreference.Primarypreferred = NewReadPreference("primarypreferred")
	ReadPreference.Secondary = NewReadPreference("secondary")
	ReadPreference.Secondarypreferred = NewReadPreference("secondarypreferred")
	ReadPreference.Nearest = NewReadPreference("nearest")
}
