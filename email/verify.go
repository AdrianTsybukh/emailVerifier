package email

type EmailVerifier struct {
	verifiers []Verifier
}

type Verifier interface {
	Verify(email string) (bool, error)
}

func NewVerifier() *EmailVerifier {
	return &EmailVerifier{
		verifiers: []Verifier{
			SyntaxVerifier{},
			DomainVerifier{},
			TempVerifier{},
		},
	}
}

func (ev *EmailVerifier) Verify(email string) (bool, error) {
	for _, verifier := range ev.verifiers {
		valid, err := verifier.Verify(email)
		if err != nil || !valid {
			return false, err
		}
	}

	return true, nil
}