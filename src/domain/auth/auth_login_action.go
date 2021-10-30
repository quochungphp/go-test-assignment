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

func (self AuthLoginAction) Execute(Username string, Password string) (tokenDetail token.TokenDetail, err error) {
	userData := &user.User{ID: Username, Password: Password}
	err = self.Db.Model(userData).WherePK().Select()
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Unauthorized")
	}

	tokenDetail, err = token.CreateToken(userData.ID)
	if err != nil {
		return token.TokenDetail{}, errors.Wrapf(err, "Generate token error")
	}
	return tokenDetail, nil
}
