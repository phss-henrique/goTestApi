package repositories

import (
	"database/sql"
	"testeApiGo/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(user *models.User) error{
	query := `INSERT INTO teste (nome, numero) VALUES ($1, $2) RETURNING id`
	return r.DB.QueryRow(query, user.Nome, user.Numero).Scan(&user.ID)
}

func (r *UserRepository) GetAll() ([]models.User, error){

	rows, err := r.DB.Query(`SELECT id, nome, numero FROM teste`)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	

	return users, nil
}