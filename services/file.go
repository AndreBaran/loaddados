package services

import (
	"loaddados/domains"
	"log"
)

type fileService struct {
	fileClient domains.FileClient
}

func NewFileService(fileClient domains.FileClient) *fileService {
	return &fileService{
		fileClient: fileClient,
	}
}

func (f *fileService) Run() []string {
	readFile, err := f.fileClient.ReadFile()
	if err != nil {
		log.Fatal(err)
	}

	lines := f.fileClient.FileScanner(readFile)

	return lines
}

func (f *fileService) Close() {
	f.Close()
}
