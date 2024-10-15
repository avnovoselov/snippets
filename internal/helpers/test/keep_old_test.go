package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/avnovoselov/snippets/internal/helpers/test"
)

func TestKeepOLD(t *testing.T) {
	plusOne := func(a int) int {
		return a + 1
	}
	plusTwo := func(a int) int {
		return a + 2
	}

	require.Equal(t, 2, plusOne(1))
	d := test.KeepOLD[func(a int) int](&plusOne, &plusTwo)

	require.Equal(t, 3, plusOne(1))
	d()
	require.Equal(t, 2, plusOne(1))

}
