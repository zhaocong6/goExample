package main

import (
	"goExample/gorm/model"
	"log"
)

func main() {
	defer model.CloseDB()

	//addUser()
	findUser()
}

func findUser() {
	user := &model.User{
		Username: 123456,
	}

	if user.FirstByUsername(); user.ID == 0 {
		log.Println("该用户未查询成功")
		return
	}

	log.Println(user)
}

func addUser() {
	res := (&model.User{
		Username: 123456,
		Password: "password",
	}).Add()

	if res.Error != nil {
		log.Panicln(res.Error)
	}
}
