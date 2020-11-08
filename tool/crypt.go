package tool

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func BcryptAndSalt(pwd string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Errorf("encrypt password error: %v\n", err)
		fmt.Printf("encrypt password error: %v\n", err)
	}
	return string(hashPwd)
}

func VerifyPasswd(hashPwd string, toVerfiyPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(toVerfiyPwd))
	if err != nil {
		log.Errorf("verify password error: %v\n", err)
		fmt.Printf("verify password error: %v\n", err)
		return false
	}
	return true
}
