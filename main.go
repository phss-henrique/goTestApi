package main

import (
	"goTestApi/db"
	"goTestApi/handlers"
	"goTestApi/repositories"
	"log"
	"net/http"
)

func main(){
	database := db.Connect()

	repo := &repositories.UserRepository{DB: database}
	handler := &handlers.UserHandler{Repo: repo}
	
	http.HandleFunc("/teste", func(w http.ResponseWriter, r *http.Request){

		switch r.Method {
		case http.MethodGet:
			//handler.GetAll(w, r)
		
		case http.MethodPost:
			handler.Create(w, r)

		case http.MethodPut:
			//handler.Update(w, r)

		case http.MethodDelete:
			//handler.Delete(w, r)			
		}
	})
	//http.HandleFunc("/teste/id", handler.FindById)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}