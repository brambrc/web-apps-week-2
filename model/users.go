package model

import (
	"quoteapp/db"

	"gorm.io/gorm"
)

type Users struct {
	gorm *gorm.DB
}

func NewUsersModel(g *gorm.DB) *Users {
	return &Users{gorm: g}
}

func (q *Users) Create(users any) error {
	// add query to create user here
}

func (q *Users) FindAll() ([]db.Users, error) {
	// add query to find all users here
	return result, err
}

func (q *Users) FindByID(id string) (db.Users, error) {
	// add query to find user by id here
	return result, err
}

func (q *Users) Update(id string, users any) error {
	// add query to update user by id here
}

func (q *Users) Delete(id string) error {
	// add query to delete user by id here
}
