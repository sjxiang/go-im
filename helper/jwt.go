package helper

import (
	"time"
	"errors"
	
	"github.com/dgrijalva/jwt-go"
)


const ExpireDuation = time.Hour * 720  // 1 月

var JWTKey = "im"


// custom 载荷
type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`

	jwt.StandardClaims
}


// 生成 JWT
func GenerateToken(identity, email string) (string, error) {
	uc := UserClaims {
		Identity: identity,
		Email: email,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireDuation).Unix(),  
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(JWTKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


// 解析 JWT
func AnalyzeToken(token string) (*UserClaims, error) {
	uc := new(UserClaims)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, nil
}
