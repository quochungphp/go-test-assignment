package auth

import (
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/quochungphp/go-test-assignment/src/domain/user"
	"github.com/quochungphp/go-test-assignment/src/pkgs/token"
)

type AuthLoginAction struct {
	Db *pg.DB
}

func (Auth AuthLoginAction) Execute(Username string, Password string) (tokenDetail token.TokenDetail, err error) {
	user := user.Users{}
	err = Auth.Db.Model(&user).Where("username = ?", Username).Where("password =?", Password).Select()
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Unauthorized")
	}

	tokenDetail, err = token.CreateToken(user.ID)
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Generate token error")
	}
	return tokenDetail, nil
}
