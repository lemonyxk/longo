/**
* @program: lottery-server
*
* @description:
*
* @author: lemon
*
* @create: 2019-10-25 21:07
**/

package longo

import (
	"time"
)

type W int

const (
	Majority W = -1
)

// WriteConcern
// W write nodes
// J write logs
// WTimeout write wait timeout, just useful when w > 1
type WriteConcern struct {
	W        W
	J        bool
	WTimeout time.Duration
}
