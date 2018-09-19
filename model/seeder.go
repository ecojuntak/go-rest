package model

import (
	"fmt"

	"github.com/bxcodec/faker"
)

var userModel *User

func RunSeeder() (err error) {
	var user = User{}

	for i := 0; i < 10; i++ {
		err = faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}

		userModel.Store(&user)
	}
	return
}
