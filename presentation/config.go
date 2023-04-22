package presentation

import (
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/MelvinKim/Hotel-Reservation-System/infrastructure/database"
	"github.com/MelvinKim/Hotel-Reservation-System/presentation/interactor"
	"github.com/MelvinKim/Hotel-Reservation-System/presentation/rest"
	"github.com/MelvinKim/Hotel-Reservation-System/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	serverTimeoutSeconds = 120
)

var allowedHeaders = []string{
	"Authorization", "Accept", "Accept-Charset", "Accept-Language",
	"Accept-Encoding", "Origin", "Host", "User-Agent", "Content-Length",
	"Content-Type", " X-Authorization", " Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Headers",
}

// Router sets up the gorilla Mux router
func Router(ctx context.Context) (*mux.Router, error) {
	create := database.NewPostgresDB()
	get := database.NewPostgresDB()
	update := database.NewPostgresDB()
	hotel := usecase.NewUseCase(create, get, update)

	// Initialize the interactor
	i, err := interactor.NewHotelInteractor(hotel)
	if err != nil {
		return nil, fmt.Errorf("can't instantiate service: %w", err)
	}

	h := rest.NewPresentationHandlers(i)

	r := mux.NewRouter()

	hotelRoutes := r.PathPrefix("/api/v1").Subrouter()
	hotelRoutes.Path("/guest").Methods(http.MethodPost).HandlerFunc(h.CreateGuest())
	hotelRoutes.Path("/reservation").Methods(http.MethodPost).HandlerFunc(h.CreateReservation())
	hotelRoutes.Path("/cancel-reservation").Methods(http.MethodPost).HandlerFunc(h.CancelReservation())

	return r, nil
}

// PrepareServer starts up a server
func PrepareServer(
	ctx context.Context,
	port int,
) *http.Server {
	// start up  the router
	r, err := Router(ctx)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Server startup error")
	}

	// start the server
	addr := fmt.Sprintf(":%d", port)
	h := handlers.CompressHandlerLevel(r, gzip.BestCompression)

	h = handlers.CORS(
		handlers.AllowedHeaders(allowedHeaders),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST"}),
	)(h)
	h = handlers.CombinedLoggingHandler(os.Stdout, h)
	h = handlers.ContentTypeHandler(
		h,
		"application/json",
		"application/x-www-form-urlencoded",
	)
	srv := &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: serverTimeoutSeconds * time.Second,
		ReadTimeout:  serverTimeoutSeconds * time.Second,
	}
	log.Infof("Server running at port %v", addr)
	return srv

}
