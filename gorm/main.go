package main

import (
	"fmt"
	"goExample/gorm/model"
	"log"
)

func main() {
	defer model.CloseDB()

	//addUser()
	user := findUser()
	Articles := make([]*model.Article, 2)
	Articles[0] = &model.Article{
		Content: "测试一",
	}
	Articles[1] = &model.Article{
		Content: "测试二",
	}
	user.Articles = Articles
	user.Save()

	fmt.Println(user)
}

func findUser() *model.User {
	user := &model.User{
		Username: 123456,
	}

	if user.FirstByUsername(); user.ID == 0 {
		log.Println("该用户未查询成功")
		return nil
	}

	log.Println(user)

	return user
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

//
//func userArticle(user *model.User) {
//	user.Related()
//}
