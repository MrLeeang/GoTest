package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// SecretKey 认证key
const (
	SecretKey = "CqOEIVNYGvJGSxnJKBTWzrmzCRAAMR5cbv4piVTZ+4A="
	Issuer    = "bugbug"
)

// jwtCustomClaims token签名信息
type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	UID interface{} // 用户id
}

// GenerateToken 生成token
func GenerateToken(uid interface{}) string {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 72).Unix()),
			Issuer:    Issuer,
		},
		uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return tokenString
}

// VerifyToken 验证token
func VerifyToken(tokenSrt string) jwt.MapClaims {
	//var verifyToken *jwt.Token
	verifyToken, err := jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(verifyToken)
	return verifyToken.Claims.(jwt.MapClaims)
}

func main() {
	// tokenStr := "AAABdB8938ABAAAAAAABkcY=.pOrGAs4/Cl/8sUEbwVCVNe5UOrSiYtbhWSgVhngp01s6gnkWLq6NNnZIHkRxpcwgdquKEMeGfzmFoveW/dW1Q0VCD6R2tvvjfa3L9kZgltg=.0Btf+tkD96bt1mVPGd+LTF4UFk8F0+T8ae9J0EOYJ5E="
	// verifyToken := VerifyToken(tokenStr)
	// fmt.Println(verifyToken)
	tokenStr := GenerateToken(12)
	// fmt.Println(tokenStr)
	// tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg1MTM3OTgsImlzcyI6ImJ1Z2J1ZyIsIlVJRCI6MTJ9.nBETsBhQ2IkzczF0sz8GMBYY0dbZeNNxZiWuc6W88Iw"
	// verifyToken := VerifyToken(tokenStr)

	fmt.Println(tokenStr)
}
