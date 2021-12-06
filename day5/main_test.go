package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = "0,9 -> 5,9\n8,0 -> 0,8\n9,4 -> 3,4\n2,2 -> 2,1\n7,0 -> 7,4\n6,4 -> 2,0\n0,9 -> 2,9\n3,4 -> 1,4\n0,0 -> 8,8\n5,5 -> 8,2"

func TestParseLine(t *testing.T) {
	l := parseLine("0,9 -> 5,9")
	require.Equal(t, 0, l.x1)
	require.Equal(t, 9, l.y1)
	require.Equal(t, 5, l.x2)
	require.Equal(t, 9, l.y2)
	require.False(t, l.horizontal())
	require.True(t, l.vertical())
}

func TestCountOverlap(t *testing.T) {
	lines := parseLines(strings.Split(testInput, "\n"))
	count := countOverlap(lines, false)
	require.Equal(t, 5, count)

	count = countOverlap(lines, true)
	require.Equal(t, 12, count)
}
