package di

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/adapters/httphandler"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/usecases"
)

func SetupRouter(uc *usecases.CustomerUseCase) http.Handler {
	r := mux.NewRouter()
	handler := httphandler.NewCustomerHandler(uc)

	r.HandleFunc("/customer/{cst_id}", handler.GetCustomerHandler).Methods("GET")

	return r
}