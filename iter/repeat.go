package iter

import "github.com/BooleanCat/go-functional/option"

// RepeatIter implements `Repeat`. See `Repeat`'s documentation.
type RepeatIter[T any] struct {
	item T
}

// Repeat instantiates a `RepeatIter` always yield the provided item.
//
// This iterator will never be exhausted.
func Repeat[T any](item T) *RepeatIter[T] {
	return &RepeatIter[T]{item}
}

// Next implements the Iterator interface for `RepeatIter`.
func (iter *RepeatIter[T]) Next() option.Option[T] {
	return option.Some(iter.item)
}

var _ Iterator[struct{}] = new(RepeatIter[struct{}])
