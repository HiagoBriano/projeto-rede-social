package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// NewUserRepository cria um repositório de usuários
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

// Create insere um usuário no banco de dados
func (u users) Create(user models.User) (int64, error) {
	return 0, nil
}
