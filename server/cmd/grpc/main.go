package grpc

import (
	"log"
	"net"

	"github.com/quankori/go-manhattan-distance/server/internals"
	"github.com/quankori/go-manhattan-distance/server/internals/cinema"
	"github.com/quankori/go-manhattan-distance/server/internals/cinema/proto"
	"github.com/quankori/go-manhattan-distance/server/pkg/logger"
	"google.golang.org/grpc"
)

func StartGrpc() {
	container := internals.GetContainer()

	lis, err := net.Listen("tcp", ":8100")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log := logger.NewLogger()
	defer log.Sync() // Ensures all logs are flushed before exit

	grpcServer := grpc.NewServer()
	cinemaServer := cinema.NewCinemaServer(container.CinemaService, log)
	proto.RegisterCinemaServiceServer(grpcServer, cinemaServer)

	log.Info("gRPC server listening on port 8100")
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve: 8100")
	}
}
