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
	router.GET("/classes", logger.Logging(handlers.GetClasses,"GetClasses"))
	router.GET("/classes/:user_id/:time", logger.Logging(handlers.GetClassWithTime, "GetClassWithTime"))
	router.POST("/classes", logger.Logging(handlers.PostClass,"PostClass"))
	return router
}

func main() {
	router := SettingRouter()

	log.Fatal(http.ListenAndServe(":8000",router))
}

