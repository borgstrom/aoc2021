package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const testInput = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1\n\n22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19\n\n 3 15  0  2 22\n 9 18 13 17  5\n19  8  7 25 23\n20 11 10 24  4\n14 21 16 12  6\n\n14 21 17 24  4\n10 16 15  9 19\n18  8 23 26 20\n22 11 13  6  5\n 2  0 12  3  7\n"

func TestParseBoards(t *testing.T) {
	raw := strings.Split(testInput, "\n")
	boards := parseBoards(raw[2:])
	require.Equal(t, [5]int{22, 13, 17, 11, 0}, boards[0].data[0])
	require.Equal(t, [5]int{1, 12, 20, 15, 19}, boards[0].data[4])

	require.Equal(t, [5]int{14, 21, 17, 24, 4}, boards[2].data[0])
	require.Equal(t, [5]int{2, 0, 12, 3, 7}, boards[2].data[4])
}

func TestBoardMatch(t *testing.T) {
	raw := strings.Split(testInput, "\n")
	numbers := parseNumbers(raw[0])
	boards := parseBoards(raw[2:])

	require.Equal(t, []int{7, 4, 9, 5, 11}, numbers[0:5])

	// All 5 first numbers should match
	for i := 0; i < 5; i++ {
		for _, board := range boards {
			require.True(t, board.match(numbers[i]))
		}
	}

	require.Equal(t, [][2]int{{2, 4}, {1, 3}, {2, 1}, {3, 4}, {0, 3}}, boards[0].matches)
	require.Equal(t, [][2]int{{2, 2}, {3, 4}, {1, 0}, {1, 4}, {3, 1}}, boards[1].matches)
	require.Equal(t, [][2]int{{4, 4}, {0, 4}, {1, 3}, {3, 4}, {3, 1}}, boards[2].matches)

	require.False(t, boards[0].winner())
	require.False(t, boards[1].winner())
	require.False(t, boards[2].winner())

	// Draw next 6 numbers
	for i := 5; i < 11; i++ {
		for _, board := range boards {
			board.match(numbers[i])
		}
	}

	// Still no winners
	require.False(t, boards[0].winner())
	require.False(t, boards[1].winner())
	require.False(t, boards[2].winner())

	// Finally, on the 12th number and the third board should win
	for _, board := range boards {
		board.match(numbers[11])
	}

	require.False(t, boards[0].winner())
	require.False(t, boards[1].winner())
	require.True(t, boards[2].winner())

	// Calculate the score
	require.Equal(t, 188, boards[2].score())
}
