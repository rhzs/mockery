package example_project

import (
	"testing"

	"gitlab.com/incubus8/gotest/assert"
)

func Foo(s Stringer) string {
	return s.String()
}

func TestString(t *testing.T) {
	mockStringer := NewMockStringer(t)
	mockStringer.EXPECT().String().Return("mockery")
	assert.Equal(t, "mockery", Foo(mockStringer))
}
