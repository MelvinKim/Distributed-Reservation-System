package database_test

import (
	"context"
	"testing"
	"time"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/MelvinKim/Hotel-Reservation-System/infrastructure/database"
	"github.com/brianvoe/gofakeit/v6"
)

func TestPostgresDB_CreateGuest(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Age:       uint(gofakeit.Uint16()),
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
			guest, err := p.CreateGuest(tt.args.ctx, tt.args.guest)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.CreateGuest() error = %v, wantErr %v", err, tt.wantErr)
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

func TestPostgresDB_GetGuests(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	type args struct {
		ctx context.Context
	}
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Age:       uint(gofakeit.Uint16()),
	}
	createdGuest, err := p.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("Can't create test guest profile: %v", err)
		return
	}
	if createdGuest == nil {
		t.Errorf("expected Guest profile but got nil: %v", createdGuest)
		return
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			guests, err := p.GetGuests(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.GetGuests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(guests) == 0 {
				t.Errorf("expected a couple of guest but got %v guests", len(guests))
			}
		})
	}
}

func TestPostgresDB_CreateHotel(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	type args struct {
		ctx   context.Context
		hotel *domain.Hotel
	}
	tests := []struct {
		name    string
		arg     args
		wantErr bool
	}{
		{
			name: "Happy Case",
			arg: args{
				ctx:   ctx,
				hotel: hotel,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hotel, err := p.CreateHotel(tt.arg.ctx, tt.arg.hotel)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.CreateGuest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && hotel != nil {
				if hotel.UUID == "" {
					t.Fatalf("expected guest profile to have a valid UUID.")
				}
				if hotel.CreatedAt == nil {
					t.Fatalf("expected guest profile to have a created at timestamp.")
				}
				if hotel.UpdatedAt == nil {
					t.Fatalf("expected guest profile to have an updated at timestamp.")
				}
			}
		})
	}
}

func TestPostgresDB_GetHotels(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hotels, err := p.GetHotels(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.GetGuests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(hotels) == 0 {
				t.Errorf("expected a couple of hotels but got %v hotels: ", len(hotels))
			}
		})
	}
}

func TestPostgresDB_CreateRoom(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}
	roomType := &domain.RoomType{
		HotelUUID: createdHotel.UUID,
		Inventory: 500,
		Reserved:  250,
	}
	createdRoomType, err := p.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("Can't create test roomType: %v", err)
		return
	}
	if createdRoomType == nil {
		t.Errorf("expected test roomType not to be nil, but got %v", createdHotel)
		return
	}
	room := &domain.Room{
		RoomTypeUUID: createdRoomType.UUID,
		HotelUUID:    createdHotel.UUID,
		Available:    true,
	}

	type args struct {
		ctx  context.Context
		room *domain.Room
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx:  ctx,
				room: room,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createdRoom, err := p.CreateRoom(tt.args.ctx, tt.args.room)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.CreateRoom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && createdRoom != nil {
				if createdRoom.UUID == "" {
					t.Fatalf("expected room to have a valid UUID.")
				}
				if createdRoom.CreatedAt == nil {
					t.Fatalf("expected room to have a created at timestamp.")
				}
				if createdRoom.UpdatedAt == nil {
					t.Fatalf("expected room to have an updated at timestamp.")
				}
			}

		})
	}
}

func TestPostgresDB_GetRooms(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}
	roomType := &domain.RoomType{
		HotelUUID: createdHotel.UUID,
		Inventory: 500,
		Reserved:  250,
	}
	createdRoomType, err := p.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("Can't create test roomType: %v", err)
		return
	}
	if createdRoomType == nil {
		t.Errorf("expected test roomType not to be nil, but got %v", createdRoomType)
		return
	}
	room := &domain.Room{
		RoomTypeUUID: createdRoomType.UUID,
		HotelUUID:    createdHotel.UUID,
		Available:    true,
	}
	createdRoom, err := p.CreateRoom(ctx, room)
	if err != nil {
		t.Errorf("Can't create test room: %v", err)
		return
	}
	if createdRoom == nil {
		t.Errorf("expected test room not to be nil, but got %v", createdRoom)
		return
	}
	type args struct {
		ctx  context.Context
		room *domain.Room
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx:  ctx,
				room: room,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rooms, err := p.GetRooms(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.GetRooms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(rooms) == 0 {
				t.Errorf("expected a couple of rooms but got: %v rooms", len(rooms))
			}
		})
	}
}

func TestPostgresDB_CreateReservation(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Age:       uint(gofakeit.Uint16()),
	}
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdGuest, err := p.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("Can't create test guest profile: %v", err)
		return
	}
	if createdGuest == nil {
		t.Errorf("expected Guest profile but got nil: %v", createdGuest)
		return
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}
	roomType := &domain.RoomType{
		HotelUUID: createdHotel.UUID,
		Inventory: 500,
		Reserved:  250,
	}
	createdRoomType, err := p.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("Can't create test roomType: %v", err)
		return
	}
	if createdRoomType == nil {
		t.Errorf("expected test roomType not to be nil, but got %v", createdHotel)
		return
	}
	reservation := &domain.Reservation{
		GuestUUID:    createdGuest.UUID,
		HotelUUID:    createdHotel.UUID,
		RoomTypeUUID: createdRoomType.UUID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(time.Hour * 72), // todo: look into how to add 2 days
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
			reservation, err := p.CreateReservation(tt.args.ctx, tt.args.reservation)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.CreateReservation() error = %v, wantErr %v", err, tt.wantErr)
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
					t.Fatalf("expected reservation to have an updated at timestamp.")
				}
			}
		})
	}
}

func TestPostgresDB_GetReservations(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Age:       uint(gofakeit.Uint16()),
	}
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdGuest, err := p.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("Can't create test guest profile: %v", err)
		return
	}
	if createdGuest == nil {
		t.Errorf("expected Guest profile but got nil: %v", createdGuest)
		return
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}
	roomType := &domain.RoomType{
		HotelUUID: createdHotel.UUID,
		Inventory: 500,
		Reserved:  250,
	}
	createdRoomType, err := p.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("Can't create test roomType: %v", err)
		return
	}
	if createdRoomType == nil {
		t.Errorf("expected test roomType not to be nil, but got %v", createdHotel)
		return
	}
	reservation := &domain.Reservation{
		GuestUUID:    createdGuest.UUID,
		HotelUUID:    createdHotel.UUID,
		RoomTypeUUID: createdRoomType.UUID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(time.Hour * 72), // todo: look into how to add 2 days
		Status:       string(domain.RESERVED),
	}
	createdReservation, err := p.CreateReservation(ctx, reservation)
	if err != nil {
		t.Errorf("Can't create test reservation: %v", err)
		return
	}
	if createdReservation == nil {
		t.Errorf("expected test reservation not to be nil, but got %v", createdReservation)
		return
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Happy Case",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reservations, err := p.GetReservations(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.GetReservations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(reservations) == 0 {
				t.Errorf("expected a couple of reservations but got: %v rooms", len(reservations))
			}
		})
	}
}

func TestPostgresDB_CancelReservation(t *testing.T) {
	ctx := context.Background()
	p := database.NewPostgresDB()
	guest := &domain.Guest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Age:       uint(gofakeit.Uint16()),
	}
	hotel := &domain.Hotel{
		Name:     gofakeit.Name(),
		Address:  gofakeit.Address().Address,
		Location: gofakeit.City(),
	}
	createdGuest, err := p.CreateGuest(ctx, guest)
	if err != nil {
		t.Errorf("Can't create test guest profile: %v", err)
		return
	}
	if createdGuest == nil {
		t.Errorf("expected Guest profile but got nil: %v", createdGuest)
		return
	}
	createdHotel, err := p.CreateHotel(ctx, hotel)
	if err != nil {
		t.Errorf("Can't create test hotel: %v", err)
		return
	}
	if createdHotel == nil {
		t.Errorf("expected test hotel not to be nil, but got %v", createdHotel)
		return
	}
	roomType := &domain.RoomType{
		HotelUUID: createdHotel.UUID,
		Inventory: 500,
		Reserved:  250,
	}
	createdRoomType, err := p.CreateRoomType(ctx, roomType)
	if err != nil {
		t.Errorf("Can't create test roomType: %v", err)
		return
	}
	if createdRoomType == nil {
		t.Errorf("expected test roomType not to be nil, but got %v", createdHotel)
		return
	}
	reservation := &domain.Reservation{
		GuestUUID:    createdGuest.UUID,
		HotelUUID:    createdHotel.UUID,
		RoomTypeUUID: createdRoomType.UUID,
		StartDate:    time.Now(),
		EndDate:      time.Now().Add(259200), // todo: look into how to add 2 days
		Status:       string(domain.RESERVED),
	}
	createdReservation, err := p.CreateReservation(ctx, reservation)
	if err != nil {
		t.Errorf("Can't create test reservation: %v", err)
		return
	}
	if createdReservation == nil {
		t.Errorf("expected test reservation not to be nil, but got %v", createdReservation)
		return
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
				GuestUUID:    createdGuest.UUID,
				RoomTypeUUID: createdRoomType.UUID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cancelledReservation, err := p.CancelReservation(tt.args.ctx, tt.args.GuestUUID, tt.args.RoomTypeUUID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresDB.CancelReservation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cancelledReservation != nil {
				if cancelledReservation.UUID == "" {
					t.Fatalf("expected reservation to have a valid UUID")
				}
				if cancelledReservation.Status == "RESERVED" {
					t.Fatalf("expected reservation status to be CANCELLED, but got: %v", cancelledReservation.Status)
				}
			}
		})
	}
}
