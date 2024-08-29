package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Armageddon6026/zender/pkg/common"
	"github.com/Armageddon6026/zender/pkg/controller"
	"github.com/Armageddon6026/zender/pkg/cron"
	"github.com/Armageddon6026/zender/pkg/gcontroller"
	"github.com/Armageddon6026/zender/pkg/middleware"
	"github.com/Armageddon6026/zender/pkg/notification"
	"github.com/Armageddon6026/zender/pkg/repository"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type server struct {
	http   *http.Server
	grpc   *grpc.Server
	cron   *cron.Cron
	config *common.SystemConfig
}

func New(config *common.SystemConfig) *server {
	repo := repository.New(&mysql.Config{
		User:                 config.Db.User,
		Passwd:               config.Db.Password,
		Addr:                 config.Db.Addr,
		DBName:               "zender",
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	})
	noti := notification.GetLoginServiceManager()

	authController := controller.NewAuthController(repo.User())
	userController := controller.NewUserController(repo.User())
	groupController := controller.NewGroupController(repo.Group())
	serviceController := controller.NewServiceController(repo.Service())
	eventController := controller.NewEventController(noti)
	grpcController := gcontroller.New(noti, repo.LoginService(), repo.Service())

	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(
		middleware.GrpcRequest(grpcController), // GRPC
		middleware.CORSMiddleware(),
		middleware.SecureMiddleware(),
		static.Serve("/", static.LocalFile(config.WebRoot, true)), // SPA
	)

	api := engine.Group(common.API_BASE_PATH, middleware.AuthJwtRequest())
	initRoute(
		api,
		authController,
		userController,
		groupController,
		serviceController,
	)
	sse := api.Group(common.API_SSE_PATH, middleware.SSERequest())
	initRoute(
		sse,
		eventController,
	)
	engine.NoRoute(controller.ApiNotFound, func(ctx *gin.Context) {
		// if request is not api, let Spa client decide to redirect the page
		ctx.Request.URL.Path = "/"
		engine.HandleContext(ctx)
	})
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: engine,
	}

	// Cron job
	cron := cron.New(noti, repo.LoginService())

	return &server{
		http:   srv,
		grpc:   grpcController,
		cron:   cron,
		config: config,
	}
}

func (s *server) Run() {
	defer s.close()

	s.cron.Task.Start()

	go func() {
		err := s.http.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Zender start")

	<-quit
}

func (s *server) close() {
	log.Println("Shutdown Server ...")
	ctx, cnacel := context.WithCancel(context.Background())
	defer cnacel()
	s.http.Shutdown(ctx)
	s.grpc.GracefulStop()
	s.cron.Task.Stop()
	time.Sleep(5 * time.Second)
}

func initRoute(api *gin.RouterGroup, controllers ...controller.Controller) {
	for _, c := range controllers {
		c.RegisterRoute(api)
	}
}
