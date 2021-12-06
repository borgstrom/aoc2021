package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = "3,4,3,1,2"

func TestTick(t *testing.T) {
	s := loadInput(testInput)
	require.Equal(t, school{0: 0, 1: 1, 2: 1, 3: 2, 4: 1, 5: 0, 6: 0, 7: 0, 8: 0}, s)

	s.tick()
	require.Equal(t, school{0: 1, 1: 1, 2: 2, 3: 1, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}, s)

	s.tick()
	require.Equal(t, school{0: 1, 1: 2, 2: 1, 3: 0, 4: 0, 5: 0, 6: 1, 7: 0, 8: 1}, s)

	s.tick()
	require.Equal(t, school{0: 2, 1: 1, 2: 0, 3: 0, 4: 0, 5: 1, 6: 1, 7: 1, 8: 1}, s)

	s.tick()
	require.Equal(t, school{0: 1, 1: 0, 2: 0, 3: 0, 4: 1, 5: 1, 6: 3, 7: 1, 8: 2}, s)

	for x := 4; x < 80; x++ {
		s.tick()
	}
	require.Equal(t, 5934, s.count())
}
