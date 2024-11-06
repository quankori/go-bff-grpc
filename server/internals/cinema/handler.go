package cinema

import (
	"context"

	"github.com/quankori/go-manhattan-distance/server/internals/cinema/proto"
	"github.com/quankori/go-manhattan-distance/server/internals/cinema/services"
	"github.com/quankori/go-manhattan-distance/server/pkg/logger"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cinemaServer struct {
	proto.UnimplementedCinemaServiceServer
	service services.CinemaService
	logger  *logger.Logger
}

func NewCinemaServer(svc services.CinemaService, logger *logger.Logger) *cinemaServer {
	return &cinemaServer{service: svc, logger: logger}
}

func (s *cinemaServer) QueryAvailableSeats(ctx context.Context, req *proto.QueryAvailableSeatsRequest) (*proto.QueryAvailableSeatsResponse, error) {
	seats := s.service.QueryAvailableSeats()
	s.logger.LogJSON("Seat", seats)
	return &proto.QueryAvailableSeatsResponse{Seats: seats}, nil
}

func (s *cinemaServer) ReserveSeat(ctx context.Context, req *proto.ReserveSeatRequest) (*proto.ReserveSeatResponse, error) {
	err := s.service.ReserveSeat(int(req.Row), int(req.Column))
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &proto.ReserveSeatResponse{
		Success: true,
		Message: "Seat reserved successfully",
	}, nil
}

func (s *cinemaServer) CancelSeat(ctx context.Context, req *proto.CancelSeatRequest) (*emptypb.Empty, error) {
	err := s.service.CancelSeat(int(req.Row), int(req.Column))
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
