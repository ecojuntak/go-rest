package model

import (
	"time"
)

type User struct {
	// gorm.Model
	ID        uint64     `gorm:"AUTO_INCREMENT" json:"id" faker:"-"`
	Name      string     `gorm:"type:varchar(100);" json:"name" faker:"name"`
	Email     string     `gorm:"type:varchar(100);unique;" json:"email" faker:"email"`
	CreatedAt time.Time  `faker:"-"`
	UpdatedAt time.Time  `faker:"-"`
	DeletedAt *time.Time `faker:"-"`
}

type UserTransaction interface {
	GetAll() []User
	GetUser() User
	Update(User)
	Delete(uint64)
}

func (u *User) GetAll() []User {
	var users []User
	db.Find(&users)

	return users
}

func (u *User) Store(user *User) {
	db.Create(&user)
}
