package main

import (
	"github.com/diogocezar/dctb-go-karatecas/data"
	"github.com/diogocezar/dctb-go-karatecas/server"
)

func main() {
	data.LoadAll()
	server := server.NewServer()
	server.Start()
}
