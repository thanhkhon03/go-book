package main

import (
	"fmt"
	//"net/http"
	//"go-jwt/handler"
	"go-jwt/driver"
	models "go-jwt/model"
	repoImpl "go-jwt/repository/repoimpl"

)

func main() {
	mongo := driver.ConnectMongoDB(DB_USER, DB_PASS)

	userrepo := repoImpl.NewUserRepo(mongo.Client.Database(DB_NAME))

	user := models.User{
		Email:       "admin@example.com",
		Password:    "1234567890",
		DisplayName: "Admin",
	}
	err := userrepo.Insert(user)
	if err == nil {
		fmt.Println("Insert ok")
	}
}
