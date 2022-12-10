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
	if err := q.gorm.Create(users).Error; err != nil {
		return err
	}
	return nil

}

func (q *Users) FindAll() ([]db.Users, error) {
	// add query to find all users here
	var result []db.Users
	if err := q.gorm.Model(&db.Users{}).Find(&result).Error; err != nil {
		return []db.Users{}, err
	}
	return result, nil
}

func (q *Users) FindByID(id string) (db.Users, error) {
	// add query to find user by id here
	var result db.Users
	if err := q.gorm.Model(&db.Users{}).Where("id = ?", id).Find(&result).Error; err != nil {
		return db.Users{}, err
	}
	return result, nil
}

func (q *Users) Update(id string, users any) error {
	// add query to update user by id here
	if err := q.gorm.Model(&db.Users{}).Where("id = ?", id).Updates(users).Error; err != nil {
		return err
	}
	return nil
}

func (q *Users) Delete(id string) error {
	// add query to delete user by id here
	if err := q.gorm.Where("id = ?", id).Delete(&db.Users{}).Error; err != nil {
		return err
	}
	return nil
}
