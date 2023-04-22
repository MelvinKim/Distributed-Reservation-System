package mock

import (
	"context"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/brianvoe/gofakeit/v6"
)

// MockCreateRepository mocks the database create repository
type MockCreateRepository struct {
	MockCreateRoomType func(
		ctx context.Context,
		roomType *domain.RoomType,
	) (*domain.RoomType, error)
	MockCreateRate func(
		ctx context.Context,
		rate *domain.Rate,
	) (*domain.Rate, error)
	MockCreateGuest func(
		ctx context.Context,
		guest *domain.Guest,
	) (*domain.Guest, error)
	MockCreateReservation func(
		ctx context.Context,
		reservation *domain.Reservation,
	) (*domain.Reservation, error)
	MockCreateHotel func(
		ctx context.Context,
		hotel *domain.Hotel,
	) (*domain.Hotel, error)
	MockCreateRoom func(
		ctx context.Context,
		hotel *domain.Room,
	) (*domain.Room, error)
}

// NewMockCreateRepository initializes
func NewMockCreateRepository() *MockCreateRepository {
	return &MockCreateRepository{
		MockCreateRoomType: func(ctx context.Context, roomType *domain.RoomType) (*domain.RoomType, error) {
			return &domain.RoomType{}, nil
		},
		MockCreateRate: func(ctx context.Context, rate *domain.Rate) (*domain.Rate, error) {
			return &domain.Rate{}, nil
		},
		MockCreateGuest: func(ctx context.Context, guest *domain.Guest) (*domain.Guest, error) {
			return &domain.Guest{}, nil
		},
		MockCreateReservation: func(ctx context.Context, reservation *domain.Reservation) (*domain.Reservation, error) {
			return &domain.Reservation{}, nil
		},
		MockCreateHotel: func(ctx context.Context, hotel *domain.Hotel) (*domain.Hotel, error) {
			return &domain.Hotel{}, nil
		},
		MockCreateRoom: func(ctx context.Context, hotel *domain.Room) (*domain.Room, error) {
			return &domain.Room{}, nil
		},
	}
}

// CreateRoomType mocks CreateRoomType
func (c *MockCreateRepository) CreateRoomType(
	ctx context.Context,
	roomType *domain.RoomType,
) (*domain.RoomType, error) {
	return c.MockCreateRoomType(ctx, roomType)
}

// CreateRate mocks CreateRate
func (c *MockCreateRepository) CreateRate(
	ctx context.Context,
	rate *domain.Rate,
) (*domain.Rate, error) {
	return c.MockCreateRate(ctx, rate)
}

// CreateGuest mocks CreateGuest
func (c *MockCreateRepository) CreateGuest(
	ctx context.Context,
	guest *domain.Guest,
) (*domain.Guest, error) {
	return c.MockCreateGuest(ctx, guest)
}

// CreateReservation mocks CreateReservation
func (c *MockCreateRepository) CreateReservation(
	ctx context.Context,
	reservation *domain.Reservation,
) (*domain.Reservation, error) {
	return c.MockCreateReservation(ctx, reservation)
}

// CreateHotel mocks CreateHotel
func (c *MockCreateRepository) CreateHotel(
	ctx context.Context,
	hotel *domain.Hotel,
) (*domain.Hotel, error) {
	return c.MockCreateHotel(ctx, hotel)
}

// CreateRoom mocks CreateRoom
func (c *MockCreateRepository) CreateRoom(
	ctx context.Context,
	room *domain.Room,
) (*domain.Room, error) {
	return c.MockCreateRoom(ctx, room)
}

// MockGetRepository mocks the database's get repository
type MockGetRepository struct {
	MockGetReservations func(
		ctx context.Context,
	) ([]domain.Reservation, error)
	MockGetHotels func(
		ctx context.Context,
	) ([]domain.Hotel, error)
	MockGetRooms func(
		ctx context.Context,
	) ([]domain.Room, error)
	MockGetGuests func(
		ctx context.Context,
	) ([]domain.Guest, error)
	MockGetRoomTypes func(
		ctx context.Context,
	) ([]domain.RoomType, error)
	MockGetRates func(
		ctx context.Context,
	) ([]domain.Rate, error)
	MockGetRate func(
		ctx context.Context,
		HotelUUID string,
		RoomTypeUUID string,
	) (*domain.Rate, error)
	MockGetRoom func(
		ctx context.Context,
		RoomTypeUUID string,
		HotelUUID string,
	) (*domain.Room, error)
}

// NewMockGetRepository initializes a new mock Get Repository
func NewMockGetRepository() *MockGetRepository {
	rate := domain.Rate{
		HotelUUID:    gofakeit.UUID(),
		RoomTypeUUID: gofakeit.UUID(),
		Rate:         30,
		Date:         gofakeit.Date(),
	}
	room := domain.Room{
		RoomTypeUUID: gofakeit.UUID(),
		HotelUUID:    gofakeit.UUID(),
		Available:    true,
	}
	return &MockGetRepository{
		MockGetReservations: func(ctx context.Context) ([]domain.Reservation, error) {
			return []domain.Reservation{}, nil
		},
		MockGetHotels: func(ctx context.Context) ([]domain.Hotel, error) {
			return []domain.Hotel{}, nil
		},
		MockGetRooms: func(ctx context.Context) ([]domain.Room, error) {
			return []domain.Room{}, nil
		},
		MockGetGuests: func(ctx context.Context) ([]domain.Guest, error) {
			return []domain.Guest{}, nil
		},
		MockGetRoomTypes: func(ctx context.Context) ([]domain.RoomType, error) {
			return []domain.RoomType{}, nil
		},
		MockGetRates: func(ctx context.Context) ([]domain.Rate, error) {
			return []domain.Rate{}, nil
		},
		MockGetRate: func(ctx context.Context, HotelUUID, RoomTypeUUID string) (*domain.Rate, error) {
			return &rate, nil
		},
		MockGetRoom: func(ctx context.Context, RoomTypeUUID, HotelUUID string) (*domain.Room, error) {
			return &room, nil
		},
	}
}

// GetReservations mocks GetReservations
func (g *MockGetRepository) GetReservations(
	ctx context.Context,
) ([]domain.Reservation, error) {
	return g.MockGetReservations(ctx)
}

// GetHotels mocks GetHotels
func (g *MockGetRepository) GetHotels(
	ctx context.Context,
) ([]domain.Hotel, error) {
	return g.MockGetHotels(ctx)
}

// GetRooms mocks GetRooms
func (g *MockGetRepository) GetRooms(
	ctx context.Context,
) ([]domain.Room, error) {
	return g.MockGetRooms(ctx)
}

// GetGuests mocks GetGuests
func (g *MockGetRepository) GetGuests(
	ctx context.Context,
) ([]domain.Guest, error) {
	return g.MockGetGuests(ctx)
}

// GetRoomTypes mocks GetRoomTypes
func (g *MockGetRepository) GetRoomTypes(
	ctx context.Context,
) ([]domain.RoomType, error) {
	return g.MockGetRoomTypes(ctx)
}

// GetRates mocks GetRates
func (g *MockGetRepository) GetRates(
	ctx context.Context,
) ([]domain.Rate, error) {
	return g.MockGetRates(ctx)
}

// GetRate mocks GetRate
func (g *MockGetRepository) GetRate(
	ctx context.Context,
	HotelUUID string,
	RoomTypeUUID string,
) (*domain.Rate, error) {
	return g.MockGetRate(ctx, HotelUUID, RoomTypeUUID)
}

// GetRoom mocks GetRoom
func (g *MockGetRepository) GetRoom(
	ctx context.Context,
	RoomTypeUUID string,
	HotelUUID string,
) (*domain.Room, error) {
	return g.MockGetRoom(ctx, RoomTypeUUID, HotelUUID)
}

// MockUpdateRepository mocks the database's Update repository
type MockUpdateRepository struct {
	MockCancelReservation func(
		ctx context.Context,
		GuestUUID string,
		RoomTypeUUID string,
	) (*domain.Reservation, error)
}

// NewMockUpdateRepository initializes a new MockUpdate Repository
func NewMockUpdateRepository() *MockUpdateRepository {
	return &MockUpdateRepository{
		MockCancelReservation: func(ctx context.Context, GuestUUID, RoomTypeUUID string) (*domain.Reservation, error) {
			return &domain.Reservation{}, nil
		},
	}
}

// CancelReservation mocks CancelReservation
func (u *MockUpdateRepository) CancelReservation(
	ctx context.Context,
	GuestUUID string,
	RoomTypeUUID string,
) (*domain.Reservation, error) {
	return u.MockCancelReservation(ctx, GuestUUID, RoomTypeUUID)
}
