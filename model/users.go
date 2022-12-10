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
	return q.gorm.Create(users).Error
}

func (q *Users) FindAll() ([]db.Users, error) {
	// add query to find all users here
	result := []db.Users{}
	err := q.gorm.Raw("select * from users").Scan(&result).Error
	return result, err
}

func (q *Users) FindByID(id string) (db.Users, error) {
	// add query to find user by id here
	result := db.Users{}
	err := q.gorm.Raw("select * from users where id = ?", id).Scan(&result).Error
	return result, err
}

func (q *Users) Update(id string, users any) error {
	err := q.gorm.Model(&db.Users{}).Where("id = ?", id).Updates(users).Error
	return err
}

func (q *Users) Delete(id string) error {
	// add query to delete user by id here
	err := q.gorm.Where("id = ?", id).Delete(&db.Users{}).Error
	return err
}
