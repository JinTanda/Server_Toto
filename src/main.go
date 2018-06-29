package main

import (
	"net/http"
	"log"
	"./logger"
	"./handlers"

	"github.com/julienschmidt/httprouter"
)

func SettingRouter() *httprouter.Router{
	router := httprouter.New()
	router.GET("/", logger.Logging(handlers.Index,"Index"))
	router.GET("/users", logger.Logging(handlers.GetUsers, "GetUsers"))
	router.POST("/users", logger.Logging(handlers.PostUser, "PostUser"))
	return router
}

func main() {
	router := SettingRouter()

	log.Fatal(http.ListenAndServe(":8000",router))
}

