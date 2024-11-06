package services

import (
	"errors"
	"sync"

	"github.com/quankori/go-manhattan-distance/server/internals/cinema/proto"
	"github.com/quankori/go-manhattan-distance/server/pkg/logger"
	"github.com/quankori/go-manhattan-distance/server/pkg/utils"
)

type Seat struct {
	Row        int
	Column     int
	IsReserved bool
}

type CinemaService interface {
	QueryAvailableSeats() []*proto.Seat
	ReserveSeat(row, column int) error
	CancelSeat(row, column int) error
	isSeatAvailable(row, column int) bool
	isDistanced(row, column int) bool
}

type cinemaService struct {
	Rows        int
	Columns     int
	MinDistance int
	Seats       [][]*Seat
	mu          sync.Mutex
}

func NewUserService(rows, columns, minDistance int) CinemaService {
	seats := make([][]*Seat, rows)
	for i := range seats {
		seats[i] = make([]*Seat, columns)
		for j := range seats[i] {
			seats[i][j] = &Seat{Row: i, Column: j, IsReserved: false}
		}
	}

	log := logger.NewLogger()
	defer log.Sync() // Ensures all logs are flushed before exit
	return &cinemaService{
		Rows:        rows,
		Columns:     columns,
		MinDistance: minDistance,
		Seats:       seats,
	}
}

// ReserveSeat implements CinemaService.
func (c *cinemaService) ReserveSeat(row int, column int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.isSeatAvailable(row, column) {
		return errors.New("seat is already reserved or invalid")
	}

	// Check if seat respects the minimum distancing rule
	if !c.isDistanced(row, column) {
		return errors.New("seat does not meet minimum distance requirements")
	}
	c.Seats[row][column].IsReserved = true
	return nil
}

// CancelSeat cancels a reservation for a given seat
func (c *cinemaService) CancelSeat(row, column int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if row < 0 || row >= c.Rows || column < 0 || column >= c.Columns {
		return errors.New("invalid seat coordinates")
	}
	if !c.Seats[row][column].IsReserved {
		return errors.New("seat is not reserved")
	}

	c.Seats[row][column].IsReserved = false
	return nil
}

// QueryAvailableSeats returns a list of available seats that can be reserved together
func (c *cinemaService) QueryAvailableSeats() []*proto.Seat {
	c.mu.Lock()
	defer c.mu.Unlock()

	var availableSeats []*proto.Seat
	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Columns; j++ {
			if !c.Seats[i][j].IsReserved && c.isDistanced(i, j) {
				availableSeats = append(availableSeats, &proto.Seat{Row: int32(i), Column: int32(j), IsReserved: false})
			}
		}
	}
	return availableSeats
}

// isSeatAvailable checks if a seat is within bounds and not reserved
func (c *cinemaService) isSeatAvailable(row, column int) bool {
	if row < 0 || row >= c.Rows || column < 0 || column >= c.Columns {
		return false
	}
	return !c.Seats[row][column].IsReserved
}

// isDistanced verifies if the minimum distance is respected for a seat
func (c *cinemaService) isDistanced(row, column int) bool {
	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Columns; j++ {
			if c.Seats[i][j].IsReserved && utils.ManhattanDistance(row, column, i, j) < c.MinDistance {
				return false
			}
		}
	}
	return true
}
