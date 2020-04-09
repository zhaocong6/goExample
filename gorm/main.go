package main

import (
	"goExample/gorm/model"
)

func main() {
	defer model.CloseDB()

}
