package controller

import (
	"net/http"
	"strconv"

	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	groupRepository repository.GroupRepository
}

func NewGroupController(groupRepository repository.GroupRepository) Controller {
	return &GroupController{
		groupRepository: groupRepository,
	}
}

func (g *GroupController) GetGroups(ctx *gin.Context) {
	groups, err := g.groupRepository.SelectGroups()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groups)
}

func (g *GroupController) GetGroupById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	group, err := g.groupRepository.SelectGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (g *GroupController) CreateGroup(ctx *gin.Context) {
	reqBody := new(model.GroupUpdate)
	err := ctx.ShouldBind(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if reqBody.Name == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "lose parameter"})
		return
	}

	err = g.groupRepository.InsertGroup(&model.GroupInfo{Name: reqBody.Name, Note: reqBody.Note})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}
	//get insert group data
	res, _ := g.groupRepository.SelectGroupByName(reqBody.Name)

	ctx.JSON(http.StatusCreated, res)
}

func (g *GroupController) UpdateGroupById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}

	reqBody := new(model.GroupUpdate)
	err = ctx.ShouldBind(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if reqBody.Name == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Empty data in body"})
		return
	}

	err = g.groupRepository.UpdateGroupById(id, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (g *GroupController) DeleteGroupById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	err = g.groupRepository.DeleteGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func (g *GroupController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/groups", g.GetGroups)
	api.GET("/groups/:id", g.GetGroupById)
	api.POST("/groups", g.CreateGroup)
	api.PUT("/groups/:id", g.UpdateGroupById)
	api.DELETE("/groups/:id", g.DeleteGroupById)
}
