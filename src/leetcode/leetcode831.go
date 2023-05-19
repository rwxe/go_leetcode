package leetcode

import (
	"regexp"
	"strings"
)

func maskPII(s string) string {
	s2 := strings.ToLower(s)
	mailRe := regexp.MustCompile("[a-z]+@[a-z]+\\.[a-z]+")
	if mailRe.MatchString(s2) {
		return maskMail(s2)
	} else {
		return maskNum(s2)
	}

}

func maskMail(mail string) string {
	usernameRe := regexp.MustCompile("[a-z]+@")
	domainRe := regexp.MustCompile("@[a-z]+\\.[a-z]+")

	usernameMatch := usernameRe.FindString(mail)
	domainMatch := domainRe.FindString(mail)

	newUsername := []byte("*******")
	newUsername[0] = usernameMatch[0]
	newUsername[len(newUsername)-1] = usernameMatch[len(usernameMatch)-2] //no @

	return string(newUsername) + domainMatch
}

func maskNum(num string) string {
	seps := []string{"+", "-", "(", ")", " "}
	for _, sep := range seps {
		num = strings.ReplaceAll(num, sep, "")
	}
	tail := num[len(num)-4 : len(num)]
	localNum := "***-***-" + tail

	var numPrefix string
	for p := len(num) - 10; p != 0; p-- {
		if p == len(num)-10 {
			numPrefix = "+"
		}
		numPrefix += "*"
		if p == 1 {
			numPrefix += "-"
		}
	}
	return numPrefix + localNum
}
