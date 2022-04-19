package main

import (
	"fmt"
	"loaddados/internal/clients/database/postgres"
	"loaddados/internal/clients/datahandler"
	"loaddados/internal/clients/file"
	"loaddados/services"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("error reading .env file. %s", envErr)
	}
}

func main() {

	ini := time.Now()

	// File service
	fileClient := file.NewFileClient("base_teste.txt")
	fileService := services.NewFileService(fileClient)
	lines := fileService.Run()
	// Por algum motivo este close está dando erro, e estou com preguica de entender o pq hahaha
	defer fileService.Close()

	// Handler service
	dataHandlerClient := datahandler.NewDataHandlerClient()
	dataHandlerService := services.NewDataHandlerService(dataHandlerClient)
	data := dataHandlerService.GetData(lines)

	// Database Service
	databaseClient := postgres.NewConnection(
		"postgres",
		os.Getenv("ENV_PORT"),
		os.Getenv("ENV_USER"),
		os.Getenv("ENV_PASSWORD"),
		os.Getenv("ENV_DBNAME"),
	)

	storageService := services.NewStorageService(databaseClient)
	if err := storageService.SendToInsert(data); err != nil {
		log.Println(err)
	}
	defer storageService.Close()

	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}
