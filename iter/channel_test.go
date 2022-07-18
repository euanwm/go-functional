package iter_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/iter"
)

func ExampleFromChannel() {
	ch := make(chan int, 2)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	fmt.Println(iter.Collect[int](iter.FromChannel(ch)))

	// Output: [1 2]
}

func TestFromChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	assert.Equal(t, iter.FromChannel(ch).Next().Unwrap(), 1)
	assert.Equal(t, iter.FromChannel(ch).Next().Unwrap(), 2)
	assert.Equal(t, iter.FromChannel(ch).Next().Unwrap(), 3)
	assert.True(t, iter.FromChannel(ch).Next().IsNone())
}

func TestFromChannelEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	assert.True(t, iter.FromChannel(ch).Next().IsNone())
}
