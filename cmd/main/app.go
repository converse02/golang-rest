package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"example.com/restapi/internal/user"
)

func main() {

	log.Println("Create router...")

	router := httprouter.New()

	log.Println("Register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("server is listening port :1234")
	log.Fatal(server.Serve(listener))
}