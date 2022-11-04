package controller

import (
	"fmt"
	"log"
	"net/http"

	user "github.com/emirkosuta/golang-auth/controller/user"
	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/users", user.GetAllUsers).Methods("GET")
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret area")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("router intialized and listening on 3200\n")
	log.Fatal(http.ListenAndServe(":3200", router))
}
