/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 21:07
**/

package lemongo

import (
	"time"
)

// WriteConcern

type WriteConcern struct {
	W        int
	J        bool
	Wtimeout time.Duration
}
