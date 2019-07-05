package main

import (
    ghandlers "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", handlers.HomeHandler)
}