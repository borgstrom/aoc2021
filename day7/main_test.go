package main

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/require"
)

const testInput = "16,1,2,0,4,2,7,1,2,14"

func TestAlign(t *testing.T) {
	crabs := loadCrabs(testInput)
	require.Equal(t, 10, len(crabs.positions))
	require.Equal(t, 0, crabs.min)
	require.Equal(t, 16, crabs.max)

	a := crabs.align(true)
	require.Equal(t, 2, a.position)
	require.Equal(t, 37, a.cost)

	fmt.Println("Took:", a.duration)

	a = crabs.align(false)
	spew.Dump(a)
	require.Equal(t, 5, a.position)
	require.Equal(t, 168, a.cost)

	fmt.Println("Took:", a.duration)
}

func TestAlignTo(t *testing.T) {
	crabs := loadCrabs(testInput)
	require.Equal(t, 10, len(crabs.positions))
	require.Equal(t, 0, crabs.min)
	require.Equal(t, 16, crabs.max)

	cost := crabs.alignTo(5, false)
	require.Equal(t, 168, cost)

	cost = crabs.alignTo(2, false)
	require.Equal(t, 206, cost)
}
