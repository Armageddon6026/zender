package controller

import (
	"net/http"

	"github.com/Armageddon6026/zender/pkg/common"
	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository"
	"github.com/Armageddon6026/zender/pkg/util"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userRepository repository.UserRepository
}

func NewAuthController(userRepository repository.UserRepository) Controller {
	return &AuthController{
		userRepository: userRepository,
	}
}

// 驗證登入資訊，並回傳token
func (a *AuthController) Login(ctx *gin.Context) {
	account, password, ok := ctx.Request.BasicAuth()
	if !ok || account == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Wrong account or password !!"})
		return
	}

	user, err := a.userRepository.SelectUserByAuth(account, password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse{Msg: err.Error()})
		return
	}
	jwtToken, err := util.CreateJwtToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.SetCookie(common.JWT_HEADER, jwtToken, int(common.JWT_EXPIRE_TIME), "/", "", true, true)
	ctx.JSON(http.StatusOK, model.Jwt{Token: jwtToken})
}

func (a *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie(common.JWT_HEADER, "", -1, "/", "", true, false)
	ctx.JSON(http.StatusOK, model.Jwt{Token: ""})
}

func (a *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/login", a.Login)
	api.POST("/logout", a.Logout)
}
