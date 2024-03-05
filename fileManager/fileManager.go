package fileManager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Failed to open file ")
	}
	scanner := bufio.NewScanner(file) // bufio.new scanner monitoring my file

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // adding to our slice line by line
	}

	err = scanner.Err() // this method helps us to handle errors

	if err != nil {
		file.Close()
		return nil, errors.New("Failed to read file ")
	}
	file.Close()
	return lines, nil
}
