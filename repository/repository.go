package repository

import (
	"context"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
)

// CreateRepository defines create contract
type CreateRepository interface {
	CreateRoomType(
		ctx context.Context,
		roomType *domain.RoomType,
	) (*domain.RoomType, error)
	CreateRate(
		ctx context.Context,
		rate *domain.Rate,
	) (*domain.Rate, error)
	CreateGuest(
		ctx context.Context,
		guest *domain.Guest,
	) (*domain.Guest, error)
	CreateReservation(
		ctx context.Context,
		reservation *domain.Reservation,
	) (*domain.Reservation, error)
	CreateHotel(
		ctx context.Context,
		hotel *domain.Hotel,
	) (*domain.Hotel, error)
	CreateRoom(
		ctx context.Context,
		room *domain.Room,
	) (*domain.Room, error)
}

// GetRepository defines get/fetch contract
type GetRepository interface {
	GetReservations(
		ctx context.Context,
	) ([]domain.Reservation, error)
	GetHotels(
		ctx context.Context,
	) ([]domain.Hotel, error)
	GetRooms(
		ctx context.Context,
	) ([]domain.Room, error)
	GetGuests(
		ctx context.Context,
	) ([]domain.Guest, error)
	GetRoomTypes(
		ctx context.Context,
	) ([]domain.RoomType, error)
	GetRates(
		ctx context.Context,
	) ([]domain.Rate, error)
	GetRate(
		ctx context.Context,
		HotelUUID string,
		RoomTypeUUID string,
	) (*domain.Rate, error)
	GetRoom(
		ctx context.Context,
		RoomTypeUUID string,
		HotelUUID string,
	) (*domain.Room, error)
}

// UpdateRepository defined update/change contract
type UpdateRepository interface {
	CancelReservation(
		ctx context.Context,
		GuestUUID string,
		RoomTypeUUID string,
	) (*domain.Reservation, error)
}

// DeleteRepository defines deletion/inactivation contract
