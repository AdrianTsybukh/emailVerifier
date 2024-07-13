package email

import (
	"net"
	"strings"
)

type DomainVerifier struct{}

func (d DomainVerifier) Verify(email string) (bool, error) {
	domain := strings.Split(email, "@")[1]
	_, err := net.LookupMX(domain)
	return err == nil, err
}