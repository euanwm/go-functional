package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleZip() {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	odds := iter.Exclude[int](iter.Count(), isEven)

	fmt.Println(iter.Take[iter.Tuple[int, int]](iter.Zip[int, int](evens, odds), 3).Collect())
	// Output: [{0 1} {2 3} {4 5}]
}

func TestZip(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	odds := iter.Exclude[int](iter.Count(), isEven)

	zipped := iter.Collect[iter.Tuple[int, int]](
		iter.Take[iter.Tuple[int, int]](iter.Zip[int, int](evens, odds), 3),
	)

	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}, {4, 5}})
}

func TestZipFirstExhausted(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Take[int](iter.Filter[int](iter.Count(), isEven), 2)
	odds := iter.Exclude[int](iter.Count(), isEven)

	zipped := iter.Collect[iter.Tuple[int, int]](iter.Zip[int, int](evens, odds))

	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}

func TestZipSecondExhausted(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	odds := iter.Take[int](iter.Exclude[int](iter.Count(), isEven), 2)

	zipped := iter.Collect[iter.Tuple[int, int]](iter.Zip[int, int](evens, odds))

	assert.SliceEqual(t, zipped, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}

func TestZipCollect(t *testing.T) {
	isEven := func(a int) bool { return a%2 == 0 }
	evens := iter.Filter[int](iter.Count(), isEven)
	odds := iter.Take[int](iter.Exclude[int](iter.Count(), isEven), 2)

	items := iter.Zip[int, int](evens, odds).Collect()

	assert.SliceEqual(t, items, []iter.Tuple[int, int]{{0, 1}, {2, 3}})
}
