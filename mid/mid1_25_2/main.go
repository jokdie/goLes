package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"mid1_25_2/internal/handler"
	"mid1_25_2/internal/repository"
	"mid1_25_2/internal/service"
	"mid1_25_2/mid1_25_2/api/user"
)

func main() {
	userRep := repository.NewUserRepository()

	s := service.NewUserService(userRep)

	server := grpc.NewServer()

	userServer := handler.NewUserGrpcServer(s)

	user.RegisterUserServiceServer(
		server,
		userServer,
	)

	listener, err := net.Listen(
		"tcp",
		":8080",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("grpc server :8080")

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
