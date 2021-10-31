package token

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

type TokenDetail struct {
	AccessToken         string `sql:"text" json:"accessToken"`
	RefreshToken        string `sql:"text" json:"refreshToken"`
	RequestUserID       string `sql:"text" json:"requestUserId"`
	AccessTokenExpires  int64  `sql:"int" json:"accessTokenExpires"`
	RefreshTokenExpires int64  `sql:"int" json:"refreshTokenExpires"`
}

func CreateToken(UserID string) (tokenDetails TokenDetail, err error) {
	// Init token details
	tokenDetails.AccessTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RequestUserID = uuid.NewV4().String()
	// Access token
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["authorized"] = true
	accessTokenClaims["correlationId"] = tokenDetails.RequestUserID
	accessTokenClaims["userId"] = UserID
	accessTokenClaims["exp"] = tokenDetails.AccessTokenExpires
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenDetails.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return TokenDetail{}, errors.Wrapf(err, "Generate access token error")
	}
	// Creating Refresh Token
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["correlationId"] = tokenDetails.RequestUserID
	refreshTokenClaims["userId"] = UserID
	refreshTokenClaims["exp"] = tokenDetails.RefreshTokenExpires
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenDetails.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return TokenDetail{}, errors.Wrapf(err, "Generate refresh token error")
	}
	return tokenDetails, nil
}

func ValidToken(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return errors.Wrapf(err, "Verify token error")
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
