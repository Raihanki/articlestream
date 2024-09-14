package repositories

import (
	"database/sql"

	"github.com/Raihanki/articlestream/internal/entities"
)

type UserRepository struct {
	DB *sql.DB
}

func (repo *UserRepository) CreateUser(request entities.CreateUser) (int, error) {
	query := "INSERT INTO users (name) VALUES ($1) RETURNING id"
	var lastId int
	err := repo.DB.QueryRow(query, request.Name).Scan(&lastId)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (repo *UserRepository) GetUserById(id int) (entities.User, error) {
	query := "SELECT id, name, created_at, updated_at FROM users WHERE id = $1"
	row := repo.DB.QueryRow(query, id)

	user := entities.User{}
	err := row.Scan(&user.Id, &user.Name, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
