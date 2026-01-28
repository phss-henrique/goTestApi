package handlers

import (
	"goTestApi/models"
	"goTestApi/repositories"
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct{
	Repo *repositories.UserRepository
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	err := h.Repo.Create(&user){
		if err != nil{
			htt.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.writeHeader(http.StatusCreated)
	json.NewDecoder(w).Encode(user)
	}
	
}