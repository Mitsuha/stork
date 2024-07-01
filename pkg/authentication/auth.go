package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/mitsuha/stork/repository/model"
	"github.com/mitsuha/stork/repository/model/dao"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)

var PasswordIncorrect = errors.New("email or password is incorrect")

func Login(ctx context.Context, email, password string) (*model.User, error) {
	user, err := dao.User.WithContext(ctx).FindByEmail(email)
	if err != nil {
		fmt.Println(email, password, err)
		return nil, PasswordIncorrect
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		fmt.Println(user.Password, password, err)
		return nil, PasswordIncorrect
	}

	return user, nil
}
