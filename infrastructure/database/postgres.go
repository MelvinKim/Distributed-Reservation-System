package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDB sets up a Postgresql database layer within the service
type PostgresDB struct {
	DB *gorm.DB
}

// Checkpreconditions assert all conditions required to run the service are met
func (p *PostgresDB) Checkpreconditions() {
	if p.DB == nil {
		log.Panic("postgresql database ORM has not been initialized")
	}
}

// NewPostgresDB initializes a new postgres db instance
func NewPostgresDB() *PostgresDB {
	db := PostgresDB{
		DB: Init(),
	}
	db.Checkpreconditions()
	return &db
}

// Migrate runs the databas's migrations
func Migrate(db *gorm.DB) {
	tables := []interface{}{
		&domain.Guest{},
		&domain.Hotel{},
		&domain.RoomType{},
		&domain.Room{},
		&domain.Rate{},
		&domain.Reservation{},
	}
	for _, table := range tables {
		if err := db.AutoMigrate(table); err != nil {
			log.Panicf("can't run migrations on table %v: err: %v", table, err)
		}
	}
}

// Init initializes a new gorm DB instance by connecting to the database specified
func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Nairobi",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("can't open connection to postgres database: %v", err)
	}
	log.Info("Database connected successfully.")
	Migrate(db)
	log.Info("Database migrations ran successfully.")
	return db
}

// GetReservations fetches all reservations from the database
func (p *PostgresDB) GetReservations(
	ctx context.Context,
) ([]domain.Reservation, error) {
	var reservations []domain.Reservation
	if err := p.DB.Where(&domain.Reservation{}).Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

// GetHotels fetches all Hotels from the database
func (p *PostgresDB) GetHotels(
	ctx context.Context,
) ([]domain.Hotel, error) {
	var hotels []domain.Hotel
	if err := p.DB.Where(&domain.Hotel{}).Find(&hotels).Error; err != nil {
		return nil, err
	}
	return hotels, nil
}

// GetRooms fetches all Rooms from the database
func (p *PostgresDB) GetRooms(
	ctx context.Context,
) ([]domain.Room, error) {
	var rooms []domain.Room
	if err := p.DB.Where(&domain.Room{}).Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

// GetGuests fetches all Guests from the database
func (p *PostgresDB) GetGuests(
	ctx context.Context,
) ([]domain.Guest, error) {
	var guests []domain.Guest
	if err := p.DB.Where(&domain.Guest{}).Find(&guests).Error; err != nil {
		return nil, err
	}
	return guests, nil
}

// GetRoomTypes fetches alll room types from the DB
func (p *PostgresDB) GetRoomTypes(
	ctx context.Context,
) ([]domain.RoomType, error) {
	var roomTypes []domain.RoomType
	if err := p.DB.Where(&domain.RoomType{}).Find(&roomTypes).Error; err != nil {
		return nil, err
	}
	return roomTypes, nil
}

// GetRoomTypes fetches rates from the DB
func (p *PostgresDB) GetRates(
	ctx context.Context,
) ([]domain.Rate, error) {
	var rates []domain.Rate
	if err := p.DB.Where(&domain.Rate{}).Find(&rates).Error; err != nil {
		return nil, err
	}
	return rates, nil
}

// GetRate fetches a rate for a specific room type
func (p *PostgresDB) GetRate(
	ctx context.Context,
	HotelUUID string,
	RoomTypeUUID string,
) (*domain.Rate, error) {
	var rate domain.Rate
	if err := p.DB.Where(&domain.Rate{
		HotelUUID:    HotelUUID,
		RoomTypeUUID: RoomTypeUUID,
	}).Find(&rate).Error; err != nil {
		return nil, err
	}
	if rate.UUID == "" {
		return nil, nil
	}

	return &rate, nil
}

// GetRoom fetches a Room within a Given hotel, belonging to a specific room type
func (p *PostgresDB) GetRoom(
	ctx context.Context,
	RoomTypeUUID string,
	HotelUUID string,
) (*domain.Room, error) {
	var room domain.Room
	if err := p.DB.Where(&domain.Room{
		HotelUUID:    HotelUUID,
		RoomTypeUUID: RoomTypeUUID,
	}).Find(&room).Error; err != nil {
		return nil, err
	}
	if room.UUID == "" {
		return nil, nil
	}
	return &room, nil
}

// CreateRoomType creates a RoomType
func (p *PostgresDB) CreateRoomType(
	ctx context.Context,
	roomType *domain.RoomType,
) (*domain.RoomType, error) {
	if err := p.DB.Create(roomType).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new room type: %v", err)
	}
	return roomType, nil
}

// CreateRate creates a rate per night for every room
func (p *PostgresDB) CreateRate(
	ctx context.Context,
	rate *domain.Rate,
) (*domain.Rate, error) {
	if err := p.DB.Create(rate).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new rate: %v", err)
	}
	return rate, nil
}

// Creates a Guest for a particular reservation
func (p *PostgresDB) CreateGuest(
	ctx context.Context,
	guest *domain.Guest,
) (*domain.Guest, error) {
	if err := p.DB.Create(guest).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new guest: %v", err)
	}
	return guest, nil
}

// CreateReservation creates a new Reservation
func (p *PostgresDB) CreateReservation(
	ctx context.Context,
	reservation *domain.Reservation,
) (*domain.Reservation, error) {
	if err := p.DB.Create(reservation).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new reservation: %v", err)
	}
	return reservation, nil
}

// CreateHotel creates a new hotel
func (p *PostgresDB) CreateHotel(
	ctx context.Context,
	hotel *domain.Hotel,
) (*domain.Hotel, error) {
	if err := p.DB.Create(hotel).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new hotel: %v", err)
	}
	return hotel, nil
}

// CreateRoom creates a new Room for a specific hotel
func (p *PostgresDB) CreateRoom(
	ctx context.Context,
	room *domain.Room,
) (*domain.Room, error) {
	if err := p.DB.Create(room).Error; err != nil {
		return nil, fmt.Errorf("infrastructure: can't create a new room: %v", err)
	}
	return room, nil
}

// CancelReservation cancels a reservation that a guest had created
func (p *PostgresDB) CancelReservation(
	ctx context.Context,
	GuestUUID string,
	RoomTypeUUID string,
) (*domain.Reservation, error) {
	var reservation domain.Reservation
	if err := p.DB.Where(&domain.Reservation{
		Status:       string(domain.RESERVED),
		GuestUUID:    GuestUUID,
		RoomTypeUUID: RoomTypeUUID,
	}).First(&reservation).Error; err != nil {
		return nil, err
	}

	reservation.Status = string(domain.CANCELLED)
	now := time.Now()
	reservation.UpdatedAt = &now
	if err := p.DB.Save(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

/*
- missing methods:
1. getting the reservations of a particular hotel
2. getting the rooms of a particular hotel
3. etc
*/
