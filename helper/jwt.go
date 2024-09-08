package helper

import (
	"blogging-platform/dto"
	"blogging-platform/model"
	"fmt"
	"os"
	"time"

	"github.com/kataras/jwt"
)

var SharedKey = []byte(os.Getenv("JWT_SHARED_KEY"))

type Claims struct {
	AccessClaims  dto.AccessClaim
	DefaultClaims dto.DefaultClaim
}

func GenerateDefaultClaims(username string) dto.DefaultClaim {
	timenow := time.Now()

	return dto.DefaultClaim{
		Expired:   int(timenow.Add(24 * time.Hour).UnixMilli()),
		NotBefore: int(timenow.UnixMilli()),
		IssuedAt:  int(timenow.UnixMilli()),
		Issuer:    "blogging-platform",
		Audience:  "blogging-platform",
		JTI:       username,
		Typ:       "",
	}
}

func GenerateToken(user model.User) (dto.Token, error) {
	if user.Username == "" {
		return dto.Token{}, fmt.Errorf("username is required")
	}

	defaultClaim := GenerateDefaultClaims(user.Username)
	defaultClaim.Typ = "id_token"

	accessClaim := dto.AccessClaim{
		ID:       int(user.ID),
		Username: user.Username,
	}

	userClaims := Claims{
		AccessClaims:  accessClaim,
		DefaultClaims: defaultClaim,
	}

	// Generate JWT
	IDToken, err := jwt.Sign(jwt.HS256, SharedKey, userClaims)
	if err != nil {
		return dto.Token{}, err
	}

	return dto.Token{AccessToken: string(IDToken)}, nil
}
