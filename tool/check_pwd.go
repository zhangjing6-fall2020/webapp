package tool

import (
	"fmt"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("Password is shorter than %d characters, set a longer password", minLength)
		return fmt.Errorf("Password is shorter than %d characters\n set a longer password\n", minLength)
	}
	if len(pwd) > maxLength {
		log.Errorf("Password is longer than %d characters, set a shorter password", maxLength)
		return fmt.Errorf("Password is longer than %d characters\n set a shorter password\n", maxLength)
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
		log.Errorf("Password is too weak, set a complex one by using different characters")
		return fmt.Errorf("Password is too weak\n set a complex one by using different characters\n")
	}
	log.Tracef("Password level: %v", level)
	return nil
}
