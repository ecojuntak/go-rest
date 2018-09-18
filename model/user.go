package model

type User struct {
	ID    uint64 `gorm:"type:int;column:id" json:"id" `
	Name  string `gorm:"type:varchar(100);column:name" json:"name" faker:"name"`
	Email string `gorm:"type:varchar(100);unique;column:email" json:"email" faker:"email"`
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

func (u *User) Store(user User) {
	db.Save(user)
}
