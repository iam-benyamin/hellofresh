package userserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"

	"github.com/iam-benyamin/hellofresh/contract/goproto/userproto"
	"github.com/iam-benyamin/hellofresh/param/userparam"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
	"github.com/iam-benyamin/hellofresh/service/userservice"
	"google.golang.org/grpc"
)

type UserServer struct {
	userproto.UnimplementedUserServiceServer
	svc userservice.Service
}

func New(svc userservice.Service) UserServer {
	return UserServer{
		UnimplementedUserServiceServer: userproto.UnimplementedUserServiceServer{},
		svc:                            svc,
	}
}

func (s UserServer) Profile(ctx context.Context, req *userproto.ProfileRequest) (*userproto.ProfileResponse, error) {
	const op = "userservice.Profile"

	profile, err := s.svc.Profile(ctx, userparam.ProfileRequest{UserID: req.UserId})
	if err != nil {
		// TODO: implement middleware logger
		e := richerror.New(op).WithErr(err)

		if e.Kind() == richerror.KindNotFound {
			return nil, status.Errorf(codes.NotFound, e.Message())
		}

		return nil, status.Errorf(codes.NotFound, e.Message())
	}

	return &userproto.ProfileResponse{
		Id:        profile.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
	}, nil
}

func (s UserServer) Start() {
	address := fmt.Sprintf(":%d", 8086)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	userproto.RegisterUserServiceServer(grpcServer, s)

	fmt.Println("Starting grpc server on " + address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("couldn't server presence grpc server")
	}
}
