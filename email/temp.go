package email

import "strings"

type TempVerifier struct{}

var tempEmailProviders = []string{
	"10minutemail.com",
	"tempmail.com",
	"yopmail.com",
	// Add more temporary email providers here
}

func (t TempVerifier) Verify(email string) (bool, error) {
	domain := strings.Split(email, "@")[1]
	for _, provider := range tempEmailProviders {
		if strings.Contains(domain, provider) {
			return false, nil
		}
	}
	return true, nil
}