package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func load() ([]string, error) {
	_, caller, _, ok := runtime.Caller(2)
	if !ok {
		return nil, fmt.Errorf("failed to get caller")
	}

	exeDir := filepath.Dir(caller)
	inputFile := filepath.Join(exeDir, "input")
	fd, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Load() ([]string, error) {
	return load()
}

func MustLoad() []string {
	input, err := load()
	if err != nil {
		panic(fmt.Errorf("failed to load: %w", err))
	}
	return input
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return i
}

func CommaInts(s string) []int {
	ints := strings.Split(s, ",")
	out := make([]int, len(ints))
	for x, i := range ints {
		out[x] = MustAtoi(i)
	}
	return out
}
