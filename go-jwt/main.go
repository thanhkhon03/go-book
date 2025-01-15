package main

import (
	"fmt"
	//"net/http"
	//"go-jwt/handler"
	"go-jwt/driver"
	repoImpl "go-jwt/repository/repoimpl"
)

func main() {
	mongo := driver.ConnectMongoDB(DB_USER, DB_PASS)

	userRepo := repoImpl.NewUserRepo(mongo.Client.Database(DB_NAME))

	/*user := models.User{
		Email:       "admin@example.com",
		Password:    "1234567890",
		DisplayName: "Admin",
	}
	err := userrepo.Insert(user)
	if err == nil { //Kiem tra neu khong loi se in ra chuoi String
		fmt.Println("Insert ok")
	}*/
	user, _ := userRepo.CheckLoginInfo("admin@example.com", "1234567890")
	fmt.Println(user)
}
