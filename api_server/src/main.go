package main

import (
	"net/http"
	"log"
	logger "./logger"
	handler "./handler"
	"github.com/julienschmidt/httprouter"
)

func SettingRouter() *httprouter.Router{
	router := httprouter.New()
	router.GET("/", logger.Logging(handler.Index,"Index"))
	router.GET("/users", logger.Logging(handler.GetUsers, "GetUsers"))
	router.POST("/users", logger.Logging(handler.PostUser, "PostUser"))
	router.GET("/classes", logger.Logging(handler.GetClasses,"GetClasses"))
	router.GET("/classes/time/:user_id/:time", logger.Logging(handler.GetClassWithTime, "GetClassWithTime"))
	router.GET("/classes/id/:class_id", logger.Logging(handler.GetClassByID, "GetClassByID"))
	router.POST("/classes", logger.Logging(handler.PostClass,"PostClass"))
	return router
}

func main() {
	router := SettingRouter()

	log.Fatal(http.ListenAndServe(":8000",router))
}

