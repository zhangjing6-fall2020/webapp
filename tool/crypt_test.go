package tool

import "testing"

func TestVerifyPasswd(t *testing.T) {
	pwd := "qwerasdf1234"
	VerifyPasswd(BcryptAndSalt(pwd), pwd)
}
