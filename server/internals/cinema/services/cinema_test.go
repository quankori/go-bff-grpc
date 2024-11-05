package services

import (
	"testing"

	"github.com/quankori/go-manhattan-distance/server/pkg/utils"
)

func TestCinemaReservation(t *testing.T) {
	cin := NewUserService(5, 5, 2)

	// Test reservation on an empty seat
	if err := cin.ReserveSeat(1, 1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test reservation on an already reserved seat
	if err := cin.ReserveSeat(1, 1); err == nil {
		t.Error("expected error for reserved seat, got none")
	}

	// Test reservation that violates distance
	if err := cin.ReserveSeat(1, 2); err == nil {
		t.Error("expected error for seat violating distance, got none")
	}

	// Test valid reservation respecting the distance
	if err := cin.ReserveSeat(3, 3); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCinemaCancellation(t *testing.T) {
	cin := NewUserService(5, 5, 2)

	// Reserve and then cancel
	_ = cin.ReserveSeat(1, 1)
	if err := cin.CancelSeat(1, 1); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Cancel on an already canceled seat
	if err := cin.CancelSeat(1, 1); err == nil {
		t.Error("expected error for canceling an unreserved seat, got none")
	}
}

func TestAvailableSeats(t *testing.T) {
	cin := NewUserService(5, 5, 2)
	_ = cin.ReserveSeat(1, 1)

	availableSeats := cin.QueryAvailableSeats()
	for _, seat := range availableSeats {
		row, col := seat.Row, seat.Column
		if (row == 1 && col == 1) || utils.ManhattanDistance(int(row), int(col), 1, 1) < cin.(*cinemaService).MinDistance {
			t.Errorf("expected seat to be unavailable, found available at (%d, %d)", row, col)
		}
	}
}
