package usecase

import (
	"context"
	"log"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/MelvinKim/Hotel-Reservation-System/repository"
)

type UsecasesContract interface {
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
	CreateRoomType(
		ctx context.Context,
		roomType *domain.RoomType,
	) (*domain.RoomType, error)
	CreateRoom(
		ctx context.Context,
		room *domain.Room,
	) (*domain.Room, error)
	CreateRate(
		ctx context.Context,
		rate *domain.Rate,
	) (*domain.Rate, error)
	GetReservations(
		ctx context.Context,
	) ([]domain.Reservation, error)
	GetRoom(
		ctx context.Context,
		RoomTypeUUID string,
		HotelUUID string,
	) (*domain.Room, error)
	GetGuests(
		ctx context.Context,
	) ([]domain.Guest, error)
	GetRoomTypes(
		ctx context.Context,
	) ([]domain.RoomType, error)
	CancelReservation(
		ctx context.Context,
		GuestUUID string,
		RoomTypeUUID string,
	) (*domain.Reservation, error)
}

// Usecase represents the Application's business logic
type Usecase struct {
	Create repository.CreateRepository
	Get    repository.GetRepository
	Update repository.UpdateRepository
}

// Checkpreconditions asserts all pre-conditions are met
func (u *Usecase) Checkpreconditions() {
	if u.Create == nil {
		log.Panicf("hotel usecase has not initialized a create repository")
	}
	if u.Get == nil {
		log.Panicf("hotel usecase has not initialized a get repository")
	}
	if u.Update == nil {
		log.Panicf("hotel  usecase has not initialized a delete repository")
	}
}

// NewUseCase initializes  a new hotel usecase
func NewUseCase(
	create repository.CreateRepository,
	get repository.GetRepository,
	update repository.UpdateRepository,
) *Usecase {
	uc := &Usecase{
		Create: create,
		Get:    get,
		Update: update,
	}
	uc.Checkpreconditions()
	return uc
}

// CreateGuest creates a new guest
func (u *Usecase) CreateGuest(
	ctx context.Context,
	guest *domain.Guest,
) (*domain.Guest, error) {
	return u.Create.CreateGuest(ctx, guest)
}

// CreateReservation creates a new reservation
func (u *Usecase) CreateReservation(
	ctx context.Context,
	reservation *domain.Reservation,
) (*domain.Reservation, error) {
	return u.Create.CreateReservation(ctx, reservation)
}

// CreateHotel creates a new Hotel
func (u *Usecase) CreateHotel(
	ctx context.Context,
	hotel *domain.Hotel,
) (*domain.Hotel, error) {
	return u.Create.CreateHotel(ctx, hotel)
}

// CreateRoomType creates a new RoomType
func (u *Usecase) CreateRoomType(
	ctx context.Context,
	roomType *domain.RoomType,
) (*domain.RoomType, error) {
	return u.Create.CreateRoomType(ctx, roomType)
}

// CreateRoom creates a new Room
func (u *Usecase) CreateRoom(
	ctx context.Context,
	room *domain.Room,
) (*domain.Room, error) {
	return u.Create.CreateRoom(ctx, room)
}

// CreateRate creates a new Rate (money charged for per night for a reservation)
func (u *Usecase) CreateRate(
	ctx context.Context,
	rate *domain.Rate,
) (*domain.Rate, error) {
	return u.Create.CreateRate(ctx, rate)
}

// GetReservations gets all reservations
func (u *Usecase) GetReservations(
	ctx context.Context,
) ([]domain.Reservation, error) {
	return u.Get.GetReservations(ctx)
}

// GetRoom gets a specific Room by RoomTypeUUID and HotelUUID
func (u *Usecase) GetRoom(
	ctx context.Context,
	RoomTypeUUID string,
	HotelUUID string,
) (*domain.Room, error) {
	return u.Get.GetRoom(ctx, RoomTypeUUID, HotelUUID)
}

// GetGuests gets all Guests who have had a reservation with the Hotel
func (u *Usecase) GetGuests(
	ctx context.Context,
) ([]domain.Guest, error) {
	return u.Get.GetGuests(ctx)
}

// GetRoomTypes gets all RoomTypes
func (u *Usecase) GetRoomTypes(
	ctx context.Context,
) ([]domain.RoomType, error) {
	return u.Get.GetRoomTypes(ctx)
}

// CancelReservation cancels a reservation made earlier
func (u *Usecase) CancelReservation(
	ctx context.Context,
	GuestUUID string,
	RoomTypeUUID string,
) (*domain.Reservation, error) {
	return u.Update.CancelReservation(ctx, GuestUUID, RoomTypeUUID)
}
