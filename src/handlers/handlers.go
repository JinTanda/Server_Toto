package handlers

import (
	"fmt"
	"net/http"
	"../models"
	"log"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w,"Welcome!")
	fmt.Fprintf(w,r.Method)
}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	if r.Method == "GET"{
		//Get all users data from data base.
		users := models.Users{
			models.User{
				Name: "test",
				University: "NIT",
				Department: "technology",
				Major: "Computer Science",
				},
			models.User{Name: "test2"},
		}
		log.Println(users)

		//Change response type from text/plain to appilication/json.
		w.Header().Set("Content-Type","appilication.json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(users); err != nil {
			panic(err)
		}
	}
}

func PostUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var user models.User

	body,err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil{
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &user); err != nil{
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil{
			panic(err)
		}
		return 
	}
	log.Println(user)
	


}
