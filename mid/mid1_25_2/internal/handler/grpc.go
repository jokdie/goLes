package handler

import (
	"context"

	"mid1_25_2/internal/service"
	"mid1_25_2/mid1_25_2/api/user"
)

type UserGrpcServer struct {
	user.UnimplementedUserServiceServer

	service *service.UserService
}

func NewUserGrpcServer(
	s *service.UserService,
) *UserGrpcServer {
	return &UserGrpcServer{
		service: s,
	}
}

func (g *UserGrpcServer) GetUser(
	ctx context.Context,
	req *user.GetUserRequest,
) (*user.User, error) {
	u, err := g.service.GetUser(
		ctx,
		int(req.Id),
	)

	if err != nil {
		return nil, err
	}

	return &user.User{
		Id: int32(u.ID),
	}, nil
}
