package services

import (
	"errors"
	"sync"

	"github.com/quankori/go-manhattan-distance/server/pkg/utils"
)

type Seat struct {
	Row        int
	Column     int
	IsReserved bool
}

type Cinema struct {
	Rows        int
	Columns     int
	MinDistance int
	Seats       [][]*Seat
	mu          sync.Mutex
}

// NewCinema initializes a new cinema with the given dimensions and minimum distance
func NewCinema(rows, columns, minDistance int) *Cinema {
	seats := make([][]*Seat, rows)
	for i := range seats {
		seats[i] = make([]*Seat, columns)
		for j := range seats[i] {
			seats[i][j] = &Seat{Row: i, Column: j, IsReserved: false}
		}
	}
	return &Cinema{
		Rows:        rows,
		Columns:     columns,
		MinDistance: minDistance,
		Seats:       seats,
	}
}

// ReserveSeat reserves a seat if it meets the minimum distancing rule
func (c *Cinema) ReserveSeat(row, column int) error {
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
func (c *Cinema) CancelSeat(row, column int) error {
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
func (c *Cinema) QueryAvailableSeats() [][]int {
	c.mu.Lock()
	defer c.mu.Unlock()

	var availableSeats [][]int
	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Columns; j++ {
			if !c.Seats[i][j].IsReserved && c.isDistanced(i, j) {
				availableSeats = append(availableSeats, []int{i, j})
			}
		}
	}
	return availableSeats
}

// isSeatAvailable checks if a seat is within bounds and not reserved
func (c *Cinema) isSeatAvailable(row, column int) bool {
	if row < 0 || row >= c.Rows || column < 0 || column >= c.Columns {
		return false
	}
	return !c.Seats[row][column].IsReserved
}

// isDistanced verifies if the minimum distance is respected for a seat
func (c *Cinema) isDistanced(row, column int) bool {
	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Columns; j++ {
			if c.Seats[i][j].IsReserved && utils.ManhattanDistance(row, column, i, j) < c.MinDistance {
				return false
			}
		}
	}
	return true
}
