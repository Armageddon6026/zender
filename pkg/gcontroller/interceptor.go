package gcontroller

import (
	"context"
	"net/http"
	"strings"

	pb "github.com/Armageddon6026/zender/pkg/gcontroller/protobuf"
	"github.com/Armageddon6026/zender/pkg/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type gInterceptor struct {
	loginServiceRepository repository.LoginServiceRepository
	serviceRepository      repository.ServiceRepository
}

func newGInterceptor(loginServiceRepository repository.LoginServiceRepository, serviceRepository repository.ServiceRepository) *gInterceptor {
	return &gInterceptor{
		loginServiceRepository: loginServiceRepository,
		serviceRepository:      serviceRepository,
	}
}

func (g *gInterceptor) authLoginServiceUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(http.StatusBadRequest, "metadata didn't set up")
	}
	appUuid, appPrivatekey, err := getGrpcAuth(md)
	if err != nil {
		return nil, status.Errorf(http.StatusBadRequest, err.Error())
	}

	// Auth from sql server when connecting grpc first time
	if strings.HasSuffix(info.FullMethod, "InsertLoginServiceInfo") {
		res, err := g.serviceRepository.SelectServiceByAuth(appUuid, appPrivatekey)
		if err != nil {
			return nil, status.Errorf(http.StatusUnauthorized, err.Error())
		}
		req.(*pb.LoginServiceInfo).Name = res.Name
		// continue grpc request
		return handler(ctx, req)
	}

	// Auth from cache
	if _, err := g.loginServiceRepository.SelectLoginServiceByUuid(appUuid); err != nil {
		return nil, status.Errorf(http.StatusUnauthorized, err.Error())
	}
	// continue grpc request
	return handler(ctx, req)
}

func (g *gInterceptor) authLoginServiceStreamInterceptor(srv any, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(http.StatusBadRequest, "metadata didn't set up")
	}
	appUuid, _, err := getGrpcAuth(md)
	if err != nil {
		return status.Errorf(http.StatusBadRequest, err.Error())
	}
	// Auth from cache
	_, err = g.loginServiceRepository.SelectLoginServiceByUuid(appUuid)
	if err != nil {
		return status.Errorf(http.StatusUnauthorized, err.Error())
	}

	return handler(srv, stream)
}

func getGrpcAuth(md metadata.MD) (uuid, privateKey string, err error) {
	if val, ok := md["uuid"]; ok && len(val) == 1 {
		uuid = val[0]
	} else {
		return "", "", status.Errorf(http.StatusBadRequest, "uuid didn't set up")
	}
	if val, ok := md["privatekey"]; ok && len(val) == 1 {
		privateKey = val[0]
	} else {
		return "", "", status.Errorf(http.StatusBadRequest, "privateKey didn't set up")
	}
	return
}
