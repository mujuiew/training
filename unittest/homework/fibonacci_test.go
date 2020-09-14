package homework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fibonacciRecursion(t *testing.T) {
	result := fibonacciRecursion(6)
	assert.Equal(t, 8, result)
	// fibonacciRecursion(6)
	// t.Error(result)
}
