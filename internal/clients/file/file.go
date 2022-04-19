package file

import (
	"bufio"
	"os"
)

type fileClient struct {
	fileName string
	file     *os.File
}

func NewFileClient(filename string) *fileClient {
	return &fileClient{
		fileName: filename,
	}
}

func (f *fileClient) ReadFile() (*os.File, error) {
	readFile, err := os.Open(f.fileName)
	if err != nil {
		return nil, err
	}
	return readFile, nil
}

func (f *fileClient) FileScanner(readFile *os.File) []string {
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func (f *fileClient) Close() {
	err := f.file.Close()
	if err != nil {
		return
	}
}
