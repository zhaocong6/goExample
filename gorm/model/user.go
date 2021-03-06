package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username uint      `gorm:"not null;unique"`
	Password string    `gorm:"type:varchar(255);not null"`
	Articles []*Article `gorm:"ForeignKey:UserId"`
}

func (u *User) Add() *gorm.DB {
	return db.Create(u)
}

func (u *User) FirstByUsername() {
	db.Where("username = ?", u.Username).First(&u)
}

func (*User) TableName() string {
	return "users"
}

func (u *User) Related(s string) {
	db.Model(&u).Related(s)
}

func (u *User) Save() *gorm.DB {
	return db.Save(&u)
}
