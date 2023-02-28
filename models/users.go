package models

import (
	"web/db"
	"web/security"
)

type Usuario struct {
	Id    int
	Name  string
	Email string
}

func UsersCreate(name string, email string, password string) {

	connectdb := db.ConnectDB()
	query, err := connectdb.Prepare("INSERT INTO users(name, email, password, status) VALUES(?, ?, ?, ?)")

	if err != nil {
		panic("Error user create!")
	}

	securePassword, err := security.EncryptPassword(password)

	if err != nil {
		panic("Error encrypt password")
	}

	query.Exec(name, email, securePassword, 1)

	connectdb.Close()

}

func ReadUser() []Usuario {

	connectdb := db.ConnectDB()

	rows, err := connectdb.Query("SELECT  id, name, email FROM users WHERE status = 1")

	if err != nil {
		panic("Error read user!")
	}

	connectdb.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario

		rows.Scan(&usuario.Id, &usuario.Name, &usuario.Email)
		usuarios = append(usuarios, usuario)
	}

	return usuarios
}

func DeleteUser(id string) {

	connectdb := db.ConnectDB()

	query, err := connectdb.Prepare("UPDATE users SET status = ? WHERE id = ?")

	if err != nil {
		panic("Error delete user!")
	}

	query.Exec(0, id)

	connectdb.Close()

}

func UpdateUser(id, name, email, password string) {

}
