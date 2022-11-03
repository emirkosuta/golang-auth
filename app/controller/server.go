package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {

}

func Start() {
	router = mux.NewRouter

	initHandlers()
	fmt.Printf("router intialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
