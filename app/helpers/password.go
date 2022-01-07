package helpers

import "golang.org/x/crypto/bcrypt"

func Generate(password []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func Compare(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err
}
