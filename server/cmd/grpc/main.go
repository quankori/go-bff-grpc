package grpc

import (
	"log"
	"net"

	"github.com/quankori/go-manhattan-distance/server/internals"
	"github.com/quankori/go-manhattan-distance/server/internals/cinema"
	"github.com/quankori/go-manhattan-distance/server/internals/cinema/proto"
	"google.golang.org/grpc"
)

func StartGrpc() {
	container := internals.GetContainer()

	lis, err := net.Listen("tcp", ":8100")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cinemaServer := cinema.NewCinemaServer(container.CinemaService)
	proto.RegisterCinemaServiceServer(grpcServer, cinemaServer)

	log.Println("gRPC server listening on port 8100")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
