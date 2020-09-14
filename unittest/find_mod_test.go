package unittest

import (
	"testing"
)

func Test_findMod(t *testing.T) {
	for i := 0; i < 100; i++ {
		result := findMod(i)
		// assert.Equal(t, result, result)
		t.Error(result)
	}

}
