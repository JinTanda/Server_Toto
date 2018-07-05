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
}

func GetClasses(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	classes := TestModelOfClasses()
	log.Println(classes)

	w.Header().Set("Content-Type","appilication.json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classes); err != nil {
		panic(err)
	}	
}

func PostClass(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	var class models.Class

	body,err := ioutil.ReadAll(io.LimitReader(r.Body,1048576))
	if err != nil{
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &class); err != nil{
		w.WriteHeader(500)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	bytes,_ := json.Marshal(&class)
	log.Println(string(bytes))
}

func GetClassWithTime(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	time := ps.ByName("time")
	user_id := ps.ByName("user_id")
	log.Println(time,user_id)
	// classes := database.getClassWithTime(time,user_id)
	classes := TestModelOfClasses()

	w.Header().Set("Content-Type","appilication.json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(classes); err != nil {
		panic(err)
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	//Get all users data from data base.
	users := models.Users{
		models.User{
			ID: "1",
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
	bytes,_ := json.Marshal(&user)
	log.Println(string(bytes))
}


func TestModelOfClasses() models.Classes{
	classes := models.Classes{
		models.Class{
			ID: "1",
			Name: "統計的データ解析特論",
			Teacher: "烏山",
			Room: "0221",
			Time: "月1",
			University: "NIT",
		},
		models.Class{
			ID: "2",
			Name: "線形代数",
			Teacher: "橋本",
			Room: "5221",
			Time: "火2",
			University: "NIT",
		},
	}
	return classes
}
