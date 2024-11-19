package authenticate

import (
	"errors"

	"github.com/gin-gonic/gin"

	"server/internal/types/model"
)

const (
	currUserCtxField = "current_user"
)

func SetCurrentUser(ctx *gin.Context, user *model.User) {
	ctx.Set(currUserCtxField, user)
}

func GetCurrentUser(ctx *gin.Context) (*model.User, error) {
	u, exist := ctx.Get(currUserCtxField)
	if !exist {
		return nil, errors.New("no current user in gin context")
	}
	user, ok := u.(*model.User)
	if !ok {
		return nil, errors.New("current user`s type is wrong")
	}
	if user == nil {
		return nil, errors.New("current user is nil")
	}
	return user, nil
}
