package tool

import (
	"testing"
)

func TestCheckUsername(t *testing.T) {
	usernames := [...]string{"123", "abc", "123@gmail.com"}
	for _, v := range usernames {
		CheckUsername(&v)
		/*if !CheckUsername(v) {
			fmt.Println(v, "is not correct email addr!")
		}*/
	}

}
