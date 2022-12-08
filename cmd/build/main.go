package main

import (
	"log"
	"os"

	"doessqlitescale/html"
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

	f, err := os.Create("_site/index.html")
	if err != nil {
		return err
	}

	if err := html.HomePage().Render(f); err != nil {
		return err
	}

	log.Println("Built")

	return nil
}
