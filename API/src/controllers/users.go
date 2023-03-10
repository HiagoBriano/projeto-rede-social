package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"

	"net/http"
)

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}

	repository := repositories.NewUserRepository(db)
	repository.Create(user)
}

// SearchUsers busca todos os usuários no banco de dados
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários"))
}

// SearchUser busca um usuário no banco de dados
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// UpdateUser atualiza um usuário no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuário"))
}

// DeleteUser deleta um usuário no banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuário"))
}
