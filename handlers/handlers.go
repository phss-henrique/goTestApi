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

	err := h.Repo.Create(&user) 
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	
	
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	users, err := h.Repo.GetAll() 
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User
	
	err := h.Repo.Update(&user)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func(h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	//	private getUsers(UserRepository repo){ 
	// 		List<User> users = repo.getAllUsers();
	//  return users;}

	err := h.Repo.Delete(id)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}



	w.WriteHeader(http.StatusNoContent)
}