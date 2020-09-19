package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	addr string
}

func InitServer(ipAddr string) *App {
	return &App{
		addr: ipAddr,
	}
}

func (app *App) Run() {
	route := Router()
	fmt.Println("Server listen on localhost:", app.addr)
	log.Println(http.ListenAndServe(app.addr, route))
}
