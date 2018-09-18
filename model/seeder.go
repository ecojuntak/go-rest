package model

import (
	"fmt"

	"github.com/bxcodec/faker"
)

var userModel *User

func RunSeeder() (err error) {
	user := User{}
	err = faker.FakeData(&user)
	if err != nil {
		fmt.Println(err)
	}

	userModel.Store(user)

	return
}
