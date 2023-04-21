package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationStatus string

const (
	RESERVED  ReservationStatus = "RESERVED"
	CANCELLED ReservationStatus = "CANCELLED"
)

// AbstractBase is an abstract struct that can be embedded in other structs
type AbstractBase struct {
	UUID      string `gorm:"primaryKey"`
	Active    bool   `gorm:"default:true"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate ensures a UUID and createdAt data is inserted
func (ab *AbstractBase) BeforeCreate(tx *gorm.DB) (err error) {
	ab.UUID = uuid.New().String()
	return
}

// Guest
type Guest struct {
	AbstractBase `gorm:"embedded"`
	FirstName    string `json:"first_name" gorm:"index"`
	LastName     string `json:"last_name" gorm:"index"`
	Email        string `json:"email" gorm:"unique"`
	Age          uint   `json:"age"`
}

// Hotel
type Hotel struct {
	AbstractBase `gorm:"embedded"`
	Name         string `json:"name" gorm:"unique, index"`
	Address      string `json:"address" gorm:"unique, index"`
	Location     string `json:"location" gorm:"index"`
}

// RoomType
type RoomType struct {
	AbstractBase `gorm:"embedded"`
	HotelUUID    string `json:"hotel_uuid"`
	Hotel        Hotel  `json:"hotel,omitempty" gorm:"foreignKey:HotelUUID"`
	Inventory    int64  `json:"inventory"`
	Reserved     int64  `json:"reserved"`
}

// Room
type Room struct {
	AbstractBase `gorm:"embedded"`
	RoomTypeUUID string   `json:"roomtype_uuid"`
	RoomType     RoomType `json:"room_type,omitempty" gorm:"foreignKey:RoomTypeUUID"`
	HotelUUID    string   `json:"hotel_uuid"`
	Hotel        Hotel    `json:"hotel,omitempty" gorm:"foreignKey:HotelUUID"`
	Available    bool     `json:"available" gorm:"default:false"`
}

// Rate represents the amount of money we we will charge for a particular room during a given data
type Rate struct {
	AbstractBase `gorm:"embedded"`
	HotelUUID    string    `json:"hotel_uuid"`
	Hotel        Hotel     `json:"hotel,omitempty" gorm:"foreignKey:HotelUUID"`
	RoomTypeUUID string    `json:"roomtype_uuid"`
	RoomType     RoomType  `json:"room_type,omitempty" gorm:"foreignKey:RoomTypeUUID"`
	Rate         int       `json:"rate"`
	Date         time.Time `json:"date"`
}

// Reservation
type Reservation struct {
	AbstractBase `gorm:"embedded"`
	GuestUUID    string    `json:"guest_uuid"`
	Guest        Guest     `json:"guest,omitempty" gorm:"foreignKey:GuestUUID"`
	HotelUUID    string    `json:"hotel_uuid"`
	Hotel        Hotel     `json:"hotel,omitempty" gorm:"foreignKey:HotelUUID"`
	RoomTypeUUID string    `json:"roomtype_uuid"`
	RoomType     RoomType  `json:"room_type,omitempty" gorm:"foreignKey:RoomTypeUUID"`
	StartDate    time.Time `json:"start_date" gorm:"not null"`
	EndDate      time.Time `json:"end_date" gorm:"not null"`
	Status       string    `json:"status"`
}

/*
Reservation status:
1. RESERVED
2. CANCELLED
*/
