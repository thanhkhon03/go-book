package jwtutils //Tao package custom

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4" // gọi thư viện jwt
)

//Secretkey để mã hoá và giải mã JWT

var secretKey = []byte("HS256")

// CreateToken tạo JWT mới
func CreateToken(userID string, duration time.Duration) (string, error) {
	// Định nghĩa payload
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(duration).Unix(), // Thoi gian het han
	}

	//Tao Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Ky token voi secret key
	return token.SignedString(secretKey)
}

// VerifyToken xac minh va giai ma JWT
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Kiem tra thuat toan ky:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return secretKey, nil

	})

	//Kiem tra loi
	if err != nil {
		return nil, err

	}

	//Tra ve ham claims neu ham hop le
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil

	}

	return nil, errors.New("invalid token")

}
