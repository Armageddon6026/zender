package util

import (
	"strconv"
	"time"

	"github.com/Armageddon6026/zender/pkg/common"
	"github.com/Armageddon6026/zender/pkg/model"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Account string `json:"account"`
	Role    int    `json:"role"`
	jwt.RegisteredClaims
}

func GetJwtSecretByByte() []byte {
	return []byte(common.JWT_SECRET)
}

func CreateJwtToken(u *model.UserInfo) (string, error) {
	now := time.Now()
	jwtId := u.Account + strconv.FormatInt(now.Unix(), 10)
	claims := Claims{
		Account: u.Account,
		Role:    u.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    common.JWT_ISSUER,
			Subject:   u.Account,
			Audience:  []string{u.Account},
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(common.JWT_EXPIRE_TIME)),
			ID:        jwtId,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(GetJwtSecretByByte())
}
