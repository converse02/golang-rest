package main

import (
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"example.com/restapi/internal/user"
	"example.com/restapi/pkg/logging"
)

func main() {

	logger := logging.GetLogger()

	logger.Info("Create router...")

	router := httprouter.New()

	logger.Info("Register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {

	logger := logging.GetLogger()

	logger.Info("start application")
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		ReadTimeout: 15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	logger.Info("server is listening port :1234")
	logger.Fatal(server.Serve(listener))
}