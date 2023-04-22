package usecase_test

import (
	"context"
	"testing"
	"time"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/MelvinKim/Hotel-Reservation-System/infrastructure/database"
	"github.com/MelvinKim/Hotel-Reservation-System/repository/mock"
	hotel "github.com/MelvinKim/Hotel-Reservation-System/usecase"
	"github.com/brianvoe/gofakeit/v6"
)

var (
	mockCreate = mock.NewMockCreateRepository()
	mockGet    = mock.NewMockGetRepository()
	mockUpdate = mock.NewMockUpdateRepository()
)

// newTestUseCase initializes a new test Usecase
func newTestUseCase() *hotel.Usecase {
	create := database.NewPostgresDB()
	get := database.NewPostgresDB()
	update := database.NewPostgresDB()
	u := hotel.NewUseCase(create, get, update)
	return u
}

// newMockTestUseCase
func newMockTestUseCase() *hotel.Usecase {
	mockUsecase := hotel.NewUseCase(mockCreate, mockGet, mockUpdate)
	return mockUsecase
}

func TestUsecase_CreateGuest(t *testing.T) {
	u := newTestUseCase()
	ctx := context.Background()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Age:       uint(gofakeit.Uint8()),
		Email:     gofakeit.Email(),
	}

	type args struct {
		ctx   context.Context
		guest *domain.Guest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy case",
			args: args{
				ctx:   ctx,
				guest: guest,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guest, err := u.CreateGuest(tt.args.ctx, tt.args.guest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && guest != nil {
				if guest.UUID == "" {
					t.Fatalf("expected guest profile to have a valid UUID.")
				}
				if guest.CreatedAt == nil {
					t.Fatalf("expected guest profile to have a created at timestamp.")
				}
				if guest.UpdatedAt == nil {
					t.Fatalf("expected guest profile to have an updated at timestamp.")
				}
			}
		})
	}
}

func TestUsecase_CreateReservation(t *testing.T) {
	u := newTestUseCase()
	ctx := context.Background()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Age:       uint(gofakeit.Uint8()),
		Email:     gofakeit.Email(),
	}
	guest, err := u.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("failed to create test guest :%v", err)
	}
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	hotel, err = u.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("failed to create test hotel: %v", hotel)
	}
	roomType := &domain.RoomType{
		HotelUUID: hotel.UUID,
		Inventory: 40,
		Reserved:  20,
	}
	roomType, err = u.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("failed to create test room type: %v", err)
	}
	reservation := &domain.Reservation{
		GuestUUID:    guest.UUID,
		HotelUUID:    hotel.UUID,
		RoomTypeUUID: roomType.UUID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(time.Hour * 72),
		Status:       string(domain.RESERVED),
	}

	type args struct {
		ctx         context.Context
		reservation *domain.Reservation
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx:         ctx,
				reservation: reservation,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservation, err := u.CreateReservation(tt.args.ctx, tt.args.reservation)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && reservation != nil {
				if reservation.UUID == "" {
					t.Fatalf("expected reservation to have a valid UUID.")
				}
				if reservation.CreatedAt == nil {
					t.Fatalf("expected reservation to have a created at timestamp.")
				}
				if reservation.UpdatedAt == nil {
					t.Fatalf("expected reservation to have an updated at timestamp. ")
				}
			}
		})
	}
}

func TestUsecase_CancelReservation(t *testing.T) {
	u := newTestUseCase()
	ctx := context.Background()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Age:       uint(gofakeit.Uint8()),
		Email:     gofakeit.Email(),
	}
	guest, err := u.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("failed to create test guest :%v", err)
	}
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	hotel, err = u.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("failed to create test hotel: %v", hotel)
	}
	roomType := &domain.RoomType{
		HotelUUID: hotel.UUID,
		Inventory: 40,
		Reserved:  20,
	}
	roomType, err = u.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("failed to create test room type: %v", err)
	}
	reservation := &domain.Reservation{
		GuestUUID:    guest.UUID,
		HotelUUID:    hotel.UUID,
		RoomTypeUUID: roomType.UUID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(time.Hour * 72),
		Status:       string(domain.RESERVED),
	}
	_, err = u.CreateReservation(ctx, reservation)
	if err != nil {
		t.Errorf("failed to create test reservation: %v", err)
	}

	type args struct {
		ctx          context.Context
		GuestUUID    string
		RoomTypeUUID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx:          ctx,
				GuestUUID:    guest.UUID,
				RoomTypeUUID: roomType.UUID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservation, err := u.CancelReservation(ctx, tt.args.GuestUUID, tt.args.RoomTypeUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CancelReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && reservation != nil {
				if reservation.UUID == "" {
					t.Fatalf("expected reservation to have a valid UUID.")
				}
				if reservation.CreatedAt == nil {
					t.Fatalf("expected reservation to have a created at timestamp.")
				}
				if reservation.UpdatedAt == nil {
					t.Fatalf("expected reservation to have an updated at timestamp. ")
				}
			}
		})
	}
}
