package main

import (
	"ScrapingQuinielamx/handler"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Root)

	return router
}
