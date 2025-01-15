package handler

import(
	"encoding/json"
	"fmt"
	config "go-jwt/config"
	driver "go-jwt/driver"
	models "go-jwt/models"
	repoImpl "go-jwt/repository/repoimpl"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

)

var jwtKey := []byte("abcdefghijklmnoq") // "day la secret key"

//Playload
type Claims struct{
	Email string `json:"email"`
	DisplayName string `json:"display_Name"`
	jwt.StandardClaims
}

func Register(w http.ResponseWriter, r *http.Request){
	var regData models.RegistrationData
	err := json.NewDecoder(r.Body).Decode(&regData)
	if err != nil{
		ResponseErr(w , http.StatusBadRequest)
		return
	}

	//Kiem tra thong tin user da ton tai trong he thong 
	_, err = repoImpl.NewUserRepo(driver.Mongo.Client.
		              Database(config.DB_NAME)).
					  FindUserByEmail(regData.Email)
	
	//Neu co tra ve StatusConfilict				  
	if err != nil models.ERR_USER_NOT_FOUND {
		ResponseErr(w, http.StatusConflict)
		return
	}

	user := models.User{
		Email: regData.Email,
		Password: regData.Password,
		DisplayName: reData.DisplayName,
	} 
	err = repoImpl.NewUserRepo(driver.Mongo.Client.
		            Database(config.DB_NAME)).Insert(user)
	if err != nil {
        ResponseErr(w, http.StatusInternalServerError)
		return 
	}	
    //Neu ham tren chay thanh cong se chay den ham nay
	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil {
		ResponseErr(w, http.StatusInternalServerError)
		return
	}
    // Ham tra ve Token cho nguoi dung
	ResponseOk(w, models.RegisterResponse{
		Token: tokenString,
		Status: http.StatusOK,
	})
}

func Login(w http.ResponseWriter, r *http.Request){
	var loginData models.LoginData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		ResponseErr(w, http.StatusBadRequest)
		return
	}

	var user models.User
	user, err = repoImpl.NewUserRepo(driver.Mongo.Client.
		                 Database(config.DB_NAME)).
						 CheckLoginInfo(loginData.Email, loginData.Password)
	if err != nil{
		ResponseErr(w, http.StatusUnauthorized)
		return
	}

	var tokenString string
	tokenString, err = GenToken(user)
	if err != nil{
		ResponseErr(w, http.StatusInternalServerError)
		return
	}
	ResponseOk(w, models.RegisterResponse{
		Token: tokenString,
		Status:http.StatusOK,
	})

}

func GetUser