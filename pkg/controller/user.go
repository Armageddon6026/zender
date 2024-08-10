package controller

import (
	"net/http"

	"github.com/Armageddon6026/zender/pkg/repository"

	"github.com/Armageddon6026/zender/pkg/model"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) Controller {
	return &UserController{
		userRepository: userRepository,
	}
}

func (u *UserController) GetUsersByRole(ctx *gin.Context) {
	// Get role from `role` header which setted by JWT middleware
	role := ctx.GetInt("role")
	if role < 1 {
		ctx.JSON(http.StatusForbidden, errorResponse{Msg: "Forbidden"})
		return
	}
	user, err := u.userRepository.SelectUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) GetUserByJwt(ctx *gin.Context) {
	// Get account from `account` header which setted by JWT middleware
	account := ctx.GetString("account")
	user, err := u.userRepository.SelectUserByAccount(account)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (u *UserController) PutUserByJwt(ctx *gin.Context) {
	reqBody := new(model.UserUpdate)
	err := ctx.ShouldBind(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if (*reqBody == model.UserUpdate{}) {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Empty data in body"})
		return
	}

	account := ctx.GetString("account")
	err = u.userRepository.UpdateUserByAccount(account, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	if reqBody.Password != "" {
		ctx.JSON(http.StatusUnauthorized, errorResponse{Msg: "Require authorization "})
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}

func (u *UserController) PutUserAuthByRole(ctx *gin.Context) {
	role := ctx.GetInt("role")
	if role < 1 {
		ctx.JSON(http.StatusForbidden, errorResponse{Msg: "Forbidden"})
		return
	}

	reqBody := new(model.UserAuthUpdate)
	err := ctx.ShouldBind(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if (*reqBody == model.UserAuthUpdate{}) {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Empty data in body"})
		return
	}

	account := ctx.Param("account")
	err = u.userRepository.UpdateUserAuthByAccount(account, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/user", u.GetUserByJwt)
	api.PUT("/user", u.PutUserByJwt)
	api.PUT("/users/:account", u.PutUserAuthByRole)
	api.GET("/users", u.GetUsersByRole)
}
