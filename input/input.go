package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func Load() ([]string, error) {
	_, caller, _, ok := runtime.Caller(1)
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
