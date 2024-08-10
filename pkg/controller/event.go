package controller

import (
	"fmt"

	"github.com/Armageddon6026/zender/pkg/notification"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	loginServiceManager *notification.LoginServiceManager
}

func NewEventController(loginServiceManager *notification.LoginServiceManager) Controller {
	return &EventController{
		loginServiceManager: loginServiceManager,
	}
}

func (e *EventController) ListenLoginServicesNotification(ctx *gin.Context) {
	theClientChan := make(chan string, 0)
	e.loginServiceManager.RegisterClient(theClientChan)
	closeNotify := ctx.Writer.CloseNotify()
	for {
		select {
		case message, ok := <-theClientChan:
			if !ok {
				return
			}
			ctx.Writer.WriteString("data:" + message + "\n\n")
			ctx.Writer.Flush()
		case <-closeNotify:
			e.loginServiceManager.UnRegisterClient(theClientChan)
			fmt.Println("stop SSE")
			return
		}
	}
}

func (e *EventController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/services-status", e.ListenLoginServicesNotification)
}
