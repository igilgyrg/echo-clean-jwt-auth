package main

import (
	"github.com/igilgyrg/todo-echo/internal/app"
	"log"
)

func main() {
	taskApplication := app.NewApp()

	err := taskApplication.Start()
	if err != nil {
		log.Fatal("error of starting application")
	}
}
