package tool

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Name string `json:"name"`
	Grade int `json:"grade"`
	jwt.RegisteredClaims
}

var myKey = []byte("8848TiPhone")

func GenToken(name string, grade int)(string,error){
	UserClaim := UserClaims{
		Name: name,
		Grade: grade,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString,err := token.SignedString(myKey)
	if err != nil{
		return "",err
	}
	return tokenString,nil
}

func ParseToken(tokenString string)(*UserClaims, error){
	userClaim := new(UserClaims)
	claims, err:=jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token)(interface{},error){
		return myKey, nil
	})
	if err!=nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

func GetMd5(s string) string {
	return fmt.Sprintf("%x",md5.Sum([]byte(s)))
}