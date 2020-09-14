package unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMod(t *testing.T) {
	result := findMod(6)
	assert.Equal(t, "Buzz", result)

}
func Test_findMod2(t *testing.T) {
	result := findMod(15)
	assert.Equal(t, "FizzBuzz", result)

}
