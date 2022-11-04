package controller

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/emirkosuta/golang-auth/controller/user"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("api/users", controller.GetAllUsers).Methods("GET")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router intialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
