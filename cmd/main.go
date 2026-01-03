package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	srv := server.NewServ(logger)

	if err := srv.Http.ListenAndServe(); err != nil {

		logger.Fatal(err)
	}
}
