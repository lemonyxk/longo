/**
* @program: lottery-server
*
* @description:
*
* @author: lemo
*
* @create: 2019-10-25 21:07
**/

package longo

import (
	"time"
)

// WriteConcern
// W write nodes
// J write logs
// WTimeout write wait timeout, just useful when w > 1
type WriteConcern struct {
	W        int
	J        bool
	WTimeout time.Duration
}
