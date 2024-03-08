package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteJSON(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("Failed to create file.")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("Failed to convert data to JSON")
	}

	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
