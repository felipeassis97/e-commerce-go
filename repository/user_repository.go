package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/model"
	"log"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur UserRepository) FindById(ID int) (*model.User, error) {
	query := "SELECT id, name FROM users WHERE id = $1"

	stmt, err := ur.connection.Prepare(query)

	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return nil, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	var userObj model.User
	err = stmt.QueryRow(ID).Scan(&userObj.ID, &userObj.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Printf("Error while executing query: %s", err)
		return nil, err
	}
	return &userObj, nil

}

func (ur UserRepository) CreateUser(user model.User) (*model.User, error) {
	query := "INSERT INTO users (name, password) VALUES ($1, $2)"
	stmt, err := ur.connection.Prepare(query)
	if err != nil {
		fmt.Printf("Error Prepare: %s", err)
		return nil, err
	}

	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("error closing connection: %v", err)
		}
	}()

	var userObj model.User
	err = stmt.QueryRow(user.Name, user.Password).Scan(&userObj.ID, &userObj.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Printf("Error while executing query: %s", err)
		return nil, err
	}

	return &userObj, nil
}
