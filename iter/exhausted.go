package iter

import "github.com/BooleanCat/go-functional/option"

// ExhaustedIter implements `Exhausted`. See `Exhausted`'s documentation.
type ExhaustedIter[T any] struct{}

// Exhausted instantiates an `ExhaustedIter` that will immediately be exhausted
// (`Next` will always return a `None` variant).
func Exhausted[T any]() *ExhaustedIter[T] {
	return new(ExhaustedIter[T])
}

// Next implements the Iterator interface for `ExhaustedIter`.
func (iter *ExhaustedIter[T]) Next() option.Option[T] {
	return option.None[T]()
}

var _ Iterator[struct{}] = new(ExhaustedIter[struct{}])

// Collect is an alternative way of invoking Collect(iter)
func (iter *ExhaustedIter[T]) Collect() []T {
	return Collect[T](iter)
}
