package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MelvinKim/Hotel-Reservation-System/application/common/dto"
	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/MelvinKim/Hotel-Reservation-System/presentation/interactor"
)

// PresentationHandlers represents all the REST API logic
type PresentationHandlers interface {
	CreateGuest() http.HandlerFunc
	CreateReservation() http.HandlerFunc
	CancelReservation() http.HandlerFunc
}

// PresentationHandlersImpl represents the usecase implementation object
type PresentationHandlersImpl struct {
	interactor *interactor.Interactor
}

// NewPresentationHandlers initializes a new REST handlers usecase
func NewPresentationHandlers(
	i *interactor.Interactor,
) PresentationHandlers {
	return &PresentationHandlersImpl{i}
}

// CreateGuest creates a new Guest profile
func (p PresentationHandlersImpl) CreateGuest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.GuestPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request boy to struct: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		guest := domain.Guest{
			FirstName: payload.FirstName,
			LastName:  payload.LastName,
			Email:     payload.Email,
			Age:       payload.Age,
		}
		createdGuest, err := p.interactor.Hotel.CreateGuest(ctx, &guest)
		if err != nil {
			msg := fmt.Sprintf("error creating guest: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdGuest)
	}
}

// CreateReservation creates a new Reservation
func (p PresentationHandlersImpl) CreateReservation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.ReservationPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request body to struct: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		reservation := &domain.Reservation{
			GuestUUID:    payload.GuestUUID,
			HotelUUID:    payload.HotelUUID,
			RoomTypeUUID: payload.RoomTypeUUID,
			StartDate:    time.Now(),
			EndDate:      time.Now().Add(time.Hour * 72),
			Status:       string(domain.RESERVED),
		}
		createdReservation, err := p.interactor.Hotel.CreateReservation(ctx, reservation)
		if err != nil {
			msg := fmt.Sprintf("error creating reservation: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdReservation)
	}
}

// CancelReservation cancels an existing Reservation
func (p PresentationHandlersImpl) CancelReservation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		payload := &dto.CancelReservationPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			msg := fmt.Sprintf("error unmarshalling request body to struct: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		cancelledReservation, err := p.interactor.Hotel.CancelReservation(ctx, payload.GuestUUID, payload.RoomTypeUUID)
		if err != nil {
			msg := fmt.Sprintf("error cancelling reservation: %v", err)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cancelledReservation)
	}
}
