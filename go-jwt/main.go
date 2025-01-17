package main

import (
	"fmt"
	config "go-jwt/config"
	"go-jwt/driver"
	"go-jwt/handler"
	"net/http"
)

func main() {
	driver.ConnectMongoDB(config.DB_USER, config.DB_PASS)

	//userRepo := repoImpl.NewUserRepo(mongo.Client.Database(DB_NAME)) Unlock this to run check login

	/*user := models.User{
		Email:       "admin@example.com",
		Password:    "1234567890",
		DisplayName: "Admin",
	}
	err := userrepo.Insert(user)
	if err == nil { //Kiem tra neu khong loi se in ra chuoi String
		fmt.Println("Insert ok")
	}*/
	/*user, _ := userRepo.CheckLoginInfo("admin@example.com", "1234567890")  Cau lenh de check login info
	fmt.Println(user)*/
	http.HandleFunc("/login", handler.Login)       //Log in vao server bang du lieu vua dang ki
	http.HandleFunc("/register", handler.Register) // Dang ki tai khoan vao server
	http.HandleFunc("/user", handler.GetUser)      // Hien thi len cho nguoi dung

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}
