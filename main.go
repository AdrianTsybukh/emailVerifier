package main

import (
	"fmt"
	"log"
	"verifier/email"
)

func main() {
	var emailToVerify string
	fmt.Scanln(&emailToVerify)
	verifier := email.NewVerifier()
	valid, err := verifier.Verify(emailToVerify)
	if err != nil {
		log.Fatalln("Verification error:", err)
	}

	if valid {
		fmt.Println("Email valid.")
	} else {
		fmt.Println("Email invalid.")
	}
}