package handlers

import (
	"goTestApi/models"
	"goTestApi/repositories"
	"encoding/json"
	"net/http"
)

type UserHandler struct{
	Repo *repositories.UserRepository
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := h.Repo.Create(&user) 
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	
	
}