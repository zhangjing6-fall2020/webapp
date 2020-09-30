package tool

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func BcryptAndSalt(pwd string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("encrypt password error: ", err)
	}
	return string(hashPwd)
}

func VerifyPasswd(hashPwd string, toVerfiyPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(toVerfiyPwd))
	if err != nil {
		fmt.Println("verify passwd error: ", err)
		return false
	}
	return true
}
