package tool

import (
	"fmt"
	"regexp"
)

const (
	levelD = iota
	LevelC
	LevelB
	LevelA
	LevelS
)

func CheckPwd(minLength, maxLength, minLevel int, pwd string) error {
	if len(pwd) < minLength {
		return fmt.Errorf("Password is shorter than %d characters, set a longer password!", minLength)
	}
	if len(pwd) > maxLength {
		return fmt.Errorf("Password is longer than %d characters, set a shorter password!", maxLength)
	}

	var level int = levelD
	patternList := []string{`[0-9]+`, `[a-zA-Z]+`, `[~!@#$%^&*?_-]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}

	if level < minLevel {
		return fmt.Errorf("Password is too weak, set a complex one by using different characters!")
	}
	return nil
}
