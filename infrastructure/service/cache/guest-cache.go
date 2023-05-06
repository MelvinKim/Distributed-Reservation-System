package cache

import "github.com/MelvinKim/Hotel-Reservation-System/domain"

type GuestCache interface {
	Set(key string, value *domain.Guest)
	Get(key string) *domain.Guest
}
