package grpc

import (
	"log"
	"net"

	"github.com/quankori/go-manhattan-distance/server/internals/services"
	"github.com/quankori/go-manhattan-distance/server/pkg/cinema"
	"google.golang.org/grpc"
)

type server struct {
	cinema.CinemaServiceServer
	svc *services.Cinema
}

func StartGrpc() {
	svc := services.NewCinema(5, 5, 2)

	lis, err := net.Listen("tcp", ":8100")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cinema.RegisterCinemaServiceServer(grpcServer, &server{svc: svc})

	log.Println("gRPC server listening on port 8100")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
