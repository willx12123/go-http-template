package handler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	auth "server/internal/service/authenticate"
	userSvc "server/internal/service/user"
	"server/internal/types/dto"
	"server/internal/types/dto/helper"
)

type UserHandler struct {
}

func (c *UserHandler) Register(ctx *gin.Context) {
	var params dto.UserRegisterParams
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		helper.RenderBadRequest(ctx, err)
		return
	}
	user, err := userSvc.FindUserByEmail(ctx, params.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		helper.RenderInternalServerError(ctx, fmt.Errorf("FindUserByEmail: %w", err), "创建用户失败")
		return
	}
	if user != nil {
		helper.RenderOK(ctx, &dto.UserRegisterResp{EmailBeenUsed: true})
		return
	}

	user, err = userSvc.CreateUser(ctx, params.Name, params.Email, params.Password)
	if err != nil {
		helper.RenderInternalServerError(ctx, fmt.Errorf("CreateUser: %w", err), "创建用户失败")
		return
	}

	token, err := auth.EncodeToken(user.ID)
	if err != nil {
		helper.RenderInternalServerError(ctx, fmt.Errorf("EncodeToken: %w", err), "创建用户失败")
		return
	}
	helper.RenderOK(ctx, &dto.UserRegisterResp{
		Token: token,
	})
}

func (c *UserHandler) Login(ctx *gin.Context) {
	var params dto.UserLoginParams
	err := ctx.ShouldBindBodyWithJSON(&params)
	if err != nil {
		helper.RenderInternalServerError(ctx, err, "参数错误")
		return
	}
	user, err := userSvc.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		if errors.Is(err, userSvc.ErrPasswordNotMatch) {
			helper.RenderOK(ctx, &dto.UserLoginResp{EmailOrPasswordWrong: true})
			return
		}
		helper.RenderInternalServerError(ctx, fmt.Errorf("LoginUser: %w", err), "登录失败")
		return
	}

	token, err := auth.EncodeToken(user.ID)
	if err != nil {
		helper.RenderInternalServerError(ctx, fmt.Errorf("EncodeToken: %w", err), "登录失败")
		return
	}
	helper.RenderOK(ctx, &dto.UserLoginResp{
		Token: token,
	})
}
