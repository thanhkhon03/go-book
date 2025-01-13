package main

import (
	"fmt"
	"go-jwt/jwtutils"
	"log"
	"time"

)

func main() {
	//Taọ token mới
	token, err := jwtutils.CreateToken("12345", time.Hour*2) // Token het han sau 2 gio
	if err != nil {
		log.Fatalf("Failed to create token: %v", err)

	}
	fmt.Println("Generated token:", token)

	//Xac minh token
	claims, err := jwtutils.VerifyToken(token)
	if err != nil {
		log.Fatalf("Failed to verify token: %v", err)

	}

	fmt.Printf("Token claims: %+v\n", claims)

}
