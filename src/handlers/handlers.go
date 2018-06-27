package handlers

import (
	"fmt"
	"net/http"
	"../models"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w,"Welcome!")
	fmt.Fprintf(w,r.Method)
}

func Users(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	users := Users{
		User{Name: "test"},
		User{Name: "test2"},
	}

}