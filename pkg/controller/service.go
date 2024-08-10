package controller

import (
	"net/http"
	"strconv"

	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/repository"

	"github.com/Armageddon6026/zender/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ServiceController struct {
	serviceRepository repository.ServiceRepository
}

func NewServiceController(serviceRepository repository.ServiceRepository) Controller {
	return &ServiceController{
		serviceRepository: serviceRepository,
	}
}

func (s *ServiceController) CreateServiceByName(ctx *gin.Context) {
	reqBody := new(model.ServiceUpdate)
	if err := ctx.ShouldBind(reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if reqBody.Name == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "lose parameter"})
		return
	}

	key, err := util.GeneratePlainPassword(20)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}
	resBody := &model.ServiceInsert{
		Name:       reqBody.Name,
		Uuid:       uuid.New().String(),
		PrivateKey: key,
		GroupId:    reqBody.GroupId,
	}
	err = s.serviceRepository.InsertService(resBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resBody)
}

func (s *ServiceController) GetServices(ctx *gin.Context) {
	services, err := s.serviceRepository.SelectServices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func (s *ServiceController) GetServicesByGroup(ctx *gin.Context) {
	group, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	services, err := s.serviceRepository.SelectServicesByGroup(group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, services)
}

func (s *ServiceController) UpdateServiceByUuid(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: uuid + " is null"})
		return
	}
	reqBody := new(model.ServiceUpdate)
	err := ctx.ShouldBind(reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: err.Error()})
		return
	}
	if reqBody.Name == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Empty data in body"})
		return
	}

	err = s.serviceRepository.UpdaateServiceByUuid(uuid, reqBody)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (s *ServiceController) DeleteServiceByUuid(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse{Msg: "Uuid is null"})
		return
	}

	err := s.serviceRepository.DeleteServiceByUuid(uuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse{Msg: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}

func (s *ServiceController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/services", s.GetServices)
	api.GET("/services/:id", s.GetServicesByGroup)
	api.POST("/services", s.CreateServiceByName)
	api.PUT("/services/:uuid", s.UpdateServiceByUuid)
	api.DELETE("/services/:uuid", s.DeleteServiceByUuid)
}
