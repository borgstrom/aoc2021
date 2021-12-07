package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = "16,1,2,0,4,2,7,1,2,14"

func TestAlign(t *testing.T) {
	crabs := loadCrabs(testInput)
	require.Equal(t, 10, len(crabs.positions))
	require.Equal(t, 0, crabs.min)
	require.Equal(t, 16, crabs.max)

	a := crabs.align()
	require.Equal(t, 2, a.position)
	require.Equal(t, 37, a.cost)

	fmt.Println("Took:", a.duration)
}
