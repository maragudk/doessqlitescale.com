package main

import (
	"errors"
	"log"
	"net/http"
	"os"
)

func main() {
	log := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	if err := start(); err != nil {
		log.Println("Error:", err)
		os.Exit(1)
	}
}

func start() error {
	log := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile|log.LUTC)

	log.Println("Starting on http://localhost:8080")

	if err := http.ListenAndServe("localhost:8080", http.FileServer(http.Dir("_site"))); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}
