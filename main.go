package main

import (
	router "github.com/diogocezar/dctb-go-karatecas/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8888"))
}
