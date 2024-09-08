package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(p string) (string, error) {
	cost := 10
	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, cost)

	if err != nil {
		return string(hash), err
	}

	return string(hash), nil
}

func ComparePass(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
