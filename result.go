/**
* @program: longo
*
* @description:
*
* @author: lemo
*
* @create: 2022-05-17 03:25
**/

package longo

// Result is the result from a Gen operation.
type Result[T any] struct {
	res T
	err error
}

func (p Result[T]) Result() T {
	return p.res
}

func (p Result[T]) Error() error {
	return p.err
}

type Index struct {
	Name string         `bson:"name"`
	Key  map[string]int `bson:"key"`
}
