package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestPower(t *testing.T) {
	gamma := gammaRate(testInput)
	require.Equal(t, int64(22), gamma)

	epsilon := epsilonRate(testInput)
	require.Equal(t, int64(9), epsilon)
}

func TestLife(t *testing.T) {
	oxygen := oxygenGeneratorRating(testInput)
	require.Equal(t, int64(23), oxygen)

	co2 := co2ScrubberRating(testInput)
	require.Equal(t, int64(10), co2)
}
