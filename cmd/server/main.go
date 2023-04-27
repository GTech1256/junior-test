package main

import (
	"log"

	"github.com/GTech1256/junior-test/internal/pkg/app"
)

func main() {
	a, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	a.StartMetrics()
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
