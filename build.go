package main

import (
	"log"
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

	log.Println("Building")

	if err := os.Mkdir("_site", 0755); err != nil {
		return err
	}

	f, err := os.Create("_site/index.html")
	if err != nil {
		return err
	}

	if _, err := f.WriteString("<html><head><title>doessqlitescale.com</title></head><body><h1>Does SQLite scale?</h1></body></html>"); err != nil {
		return err
	}

	return nil
}
