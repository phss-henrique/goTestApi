package repositories

import (
	"database/sql"
	"goTestApi/models"
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

	for rows.Next(){
		var u models.User
		if err := rows.Scan(&u.ID, &u.Nome, &u.Numero ); err != nil{
			return nil, err
		} 
		users = append(users,u)
	}

	return users, nil
}

func (r *UserRepository) FindById(id int) (*models.User, error){
	var u models.User

	err := r.DB.QueryRow(`SELECT id,nome,numero FROM teste WHERE id = $1 `, id).Scan(&u.ID, &u.Nome, &u.Numero)

	if err != nil{
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) Update(user *models.User) error{
	
	_, error := r.DB.Exec(
		`UPDATE teste SET nome = $1, numero = $2 id = $3`, user.Nome, user.Numero, user.ID, 
	)

	return error

}


func (r *UserRepository) Delete(id int) error{
		_, error := r.DB.Exec(
			`DELETE FROM teste WHERE id = $1`,id,
		)
	return error
}