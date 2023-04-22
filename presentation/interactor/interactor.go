// Package interactor represents reusable chunks of code that abstract
// logic from presenters while simplifying your app and making future changes effortless
package interactor

import (
	hotel "github.com/MelvinKim/Hotel-Reservation-System/usecase"
)

// Interactor represents an assemble of all use cases into a single object that can be instantiated anywhere
type Interactor struct {
	Hotel hotel.UsecasesContract
}

// NewHotelInteractor returns a new hotel interactor
func NewHotelInteractor(
	hotel hotel.UsecasesContract,
) (*Interactor, error) {
	return &Interactor{
		Hotel: hotel,
	}, nil
}
