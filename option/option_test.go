package option_test

import (
	"fmt"
	"testing"

	"github.com/BooleanCat/go-functional/internal/assert"
	"github.com/BooleanCat/go-functional/option"
)

func TestSomeStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", option.Some("foo")), "Some(foo)")
	assert.Equal(t, fmt.Sprintf("%s", option.Some(42)), "Some(42)")
}

func TestNoneStringer(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s", option.None[string]()), "None")
}

func TestSomeUnwrap(t *testing.T) {
	assert.Equal(t, option.Some(42).Unwrap(), 42)
}

func TestNoneUnwrap(t *testing.T) {
	defer func() {
		assert.Equal(t, fmt.Sprint(recover()), "called `Option.Unwrap()` on a `None` value")
	}()

	option.None[string]().Unwrap()
	t.Error("did not panic")
}

func TestSomeUnwrapOr(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOr(3), 42)
}

func TestNoneUnwrapOr(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOr(3), 3)
}

func TestSomeUnwrapOrElse(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOrElse(func() int { return 41 }), 42)
}

func TestNoneUnwrapOrElse(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOrElse(func() int { return 41 }), 41)
}

func TestSomeUnwrapOrZero(t *testing.T) {
	assert.Equal(t, option.Some(42).UnwrapOrZero(), 42)
}

func TestNoneUnwrapOrZero(t *testing.T) {
	assert.Equal(t, option.None[int]().UnwrapOrZero(), 0)
}

func TestIsSome(t *testing.T) {
	assert.True(t, option.Some(42).IsSome())
	assert.False(t, option.None[int]().IsSome())
}

func TestIsNone(t *testing.T) {
	assert.False(t, option.Some(42).IsNone())
	assert.True(t, option.None[int]().IsNone())
}

func TestSomeValue(t *testing.T) {
	value, ok := option.Some(42).Value()
	assert.Equal(t, value, 42)
	assert.True(t, ok)
}

func TestNoneValue(t *testing.T) {
	value, ok := option.None[int]().Value()
	assert.Equal(t, value, 0)
	assert.False(t, ok)
}