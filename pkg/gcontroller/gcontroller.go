package gcontroller

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "github.com/Armageddon6026/zender/pkg/gcontroller/protobuf"
	"github.com/Armageddon6026/zender/pkg/model"
	"github.com/Armageddon6026/zender/pkg/notification"
	"github.com/Armageddon6026/zender/pkg/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Gcontrollers struct {
	loginServiceManager    *notification.LoginServiceManager
	loginServiceRepository repository.LoginServiceRepository
	pb.UnimplementedServicesFuncServer
}

func New(loginServiceManager *notification.LoginServiceManager,
	loginServiceRepository repository.LoginServiceRepository,
	serviceRepository repository.ServiceRepository,
) *grpc.Server {
	interceptors := newGInterceptor(loginServiceRepository, serviceRepository)
	g := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.authLoginServiceUnaryInterceptor),
		grpc.StreamInterceptor(interceptors.authLoginServiceStreamInterceptor),
	)
	pb.RegisterServicesFuncServer(g, &Gcontrollers{
		loginServiceManager:    loginServiceManager,
		loginServiceRepository: loginServiceRepository,
	})
	return g
}

func (g *Gcontrollers) InsertLoginServiceInfo(ctx context.Context, loginServiceInfo *pb.LoginServiceInfo) (*pb.Response, error) {
	//Deep copy from grpc object to sse object
	err := g.loginServiceRepository.InsertLoginService(&model.LoginService{
		Uuid: loginServiceInfo.Uuid,
		Name: loginServiceInfo.Name,
		Ip:   loginServiceInfo.Ip,
	})
	if err != nil {
		return &pb.Response{Code: http.StatusInternalServerError, Message: loginServiceInfo.Uuid + " has been Login"}, nil
	}

	g.loginServiceManager.SendMessage(g.loginServiceRepository.SelectLoginServices())
	fmt.Println(loginServiceInfo.Uuid, " Login success")

	return &pb.Response{Code: http.StatusOK, Message: loginServiceInfo.Uuid + " Login success"}, nil
}

func (g *Gcontrollers) UpdateLoginServiceInfo(stream pb.ServicesFunc_UpdateLoginServiceInfoServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	appUuid, _, _ := getGrpcAuth(md)
	for {
		in, err := stream.Recv()
		if err != nil {
			fmt.Println(appUuid, " disconnected")
			break
		}
		if err == io.EOF {
			log.Println(appUuid, " end stream")
			break
		}
		//Deep copy from grpc object to sse object
		g.loginServiceRepository.UpdateLoginService(&model.LoginService{
			Uuid:             appUuid,
			ScanFiles:        copyScanFile2Map(in.ScanFiles),
			ScanApplications: copyScanApp2Map(in.ScanApplications),
		})
	}
	// disconnect grpc stream
	g.loginServiceRepository.DeleteLoginServiceByUuid(appUuid)
	g.loginServiceManager.SendMessage(g.loginServiceRepository.SelectLoginServices())

	return stream.SendAndClose(&pb.Response{Code: http.StatusOK, Message: "Logout success"})
}

func copyScanFile2Map(scanfiles []*pb.ScanFile) map[string]model.ScanFile {
	scanFilesMap := make(map[string]model.ScanFile, len(scanfiles))
	for _, v := range scanfiles {
		scanFilesMap[v.Path] = model.ScanFile{Path: v.Path, DataCount: v.DataCount, LastDataTime: v.LastTime}
	}
	return scanFilesMap
}

func copyScanApp2Map(scanApps []*pb.ScanApplication) map[string]model.ScanApplication {
	scanAppsMap := make(map[string]model.ScanApplication, len(scanApps))
	for _, v := range scanApps {
		scanAppsMap[v.Name] = model.ScanApplication{Name: v.Name, LastDataTime: v.LastTime}
	}
	return scanAppsMap
}
