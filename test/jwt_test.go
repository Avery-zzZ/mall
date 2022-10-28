package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

var myKey = []byte("mall")

func TestGenToken(t *testing.T){
	UserClaim := &UserClaims{
		Name: "kiki",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1*time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString,err := token.SignedString(myKey)
	if err != nil{
		t.Fatal(err)
	}
	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T){
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoia2lraSIsImV4cCI6MTY2NjkyNDU2OH0._rTntRkqqHiEFkA4BE5CM0CsE723LLLxOqF8Ncz87fQ"
	UserClaim := new(UserClaims)
	claims, err:=jwt.ParseWithClaims(tokenString, UserClaim, func(token *jwt.Token)(interface{},error){
		return myKey, nil
	})
	if err!=nil {
		t.Fatal(err)
	}
	if claims.Valid{
		fmt.Println(UserClaim)
	}
}