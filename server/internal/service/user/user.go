package user

import (
	"context"
	"errors"
	"fmt"

	"server/internal/dal/db/query"
	"server/internal/types/model"
)

var (
	ErrPasswordNotMatch = errors.New("password not match")
)

func CreateUser(ctx context.Context, name, email, password string) (*model.User, error) {
	pwdDigest, err := encodePassword(password)
	if err != nil {
		return nil, fmt.Errorf("encode password fail: %w", err)
	}
	user := &model.User{
		Name:           name,
		Email:          email,
		PasswordDigest: pwdDigest,
	}
	err = query.User.WithContext(ctx).Create(user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return user, nil
}

func LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := FindUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("FindByEmail: %w", err)
	}
	if !checkPassword(user.PasswordDigest, password) {
		return nil, ErrPasswordNotMatch
	}
	return user, nil
}

func FindUser(ctx context.Context, id uint) (*model.User, error) {
	return query.User.WithContext(ctx).Where(query.User.ID.Eq(id)).First()
}

func FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return query.User.WithContext(ctx).FindByEmail(email)
}
