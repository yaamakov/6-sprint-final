package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "SERVER: ", log.Ldate|log.Ltime|log.Lshortfile)

	server := server.NewServer(logger)
	err := server.Start()
	if err != nil {
		logger.Fatalf("Server not starting with error: %v", err)
	}

	defer server.Shutdown().Error()

}
