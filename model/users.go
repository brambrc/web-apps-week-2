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
	if err := q.gorm.Create(users).Error; err != nil {
		return err
	}
	return nil
	// add query to create user here

}

func (q *Users) FindAll() ([]db.Users, error) {
	result := []db.Users{}
	if err := q.gorm.Find(&result).Error; err != nil {
		return nil, err
	}
	// add query to find all users here

	return result, nil
}

func (q *Users) FindByID(id string) (db.Users, error) {
	var result db.Users
	err := q.gorm.Where("id = ?", id).First(&result)
	if err.Error != nil {
		return db.Users{}, err.Error
	}
	// add query to find user by id here
	return db.Users{}, nil
}

func (q *Users) Update(id string, users any) error {
	err := q.gorm.Where("id = ?", id).Updates(users).Error
	if err != nil {
		return err
	}
	return nil
	// add query to update user by id here
}

func (q *Users) Delete(id string) error {
	err := q.gorm.Where("id = ?", id).Delete(&db.Users{}).Error
	if err != nil {
		return err
	}
	return nil

	// add query to delete user by id here
}
