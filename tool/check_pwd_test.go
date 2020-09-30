package tool

import (
	"testing"
)

func TestCheckPwd(t *testing.T) {
	var minLength int = 8
	var maxLength int = 32
	var minLevel int = 2

	pwds := [...]string{"123", "12345678", "abcdefgh", "QWERTYUI", "asdfZXCV"}
	for _, p := range pwds {
		CheckPwd(minLength, maxLength, minLevel, p)
		//err := CheckPwd(minLength, maxLength, minLevel, p)
		//fmt.Println(p, "password", "error:\n", err)
	}
}
