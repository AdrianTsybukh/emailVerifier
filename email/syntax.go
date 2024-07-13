package email

import "regexp"

type SyntaxVerifier struct {}

func (s SyntaxVerifier) Verify(email string) (bool, error) {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email), nil
}