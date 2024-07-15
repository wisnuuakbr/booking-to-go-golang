package httphandler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/dto"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/usecases"
)

type CustomerHandler struct {
	CustomerUseCase *usecases.CustomerUseCase
}

func NewCustomerHandler(uc *usecases.CustomerUseCase) *CustomerHandler {
	return &CustomerHandler{
		CustomerUseCase: uc,
	}
}

func (h *CustomerHandler) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	ctx 		:= r.Context()
	params 		:= mux.Vars(r)
	cstID, err  := strconv.Atoi(params["cst_id"])

	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusInternalServerError)
		return
	}

	customer, familyList, err := h.CustomerUseCase.GetCustomer(ctx, cstID)
	if err != nil {
		log.Printf("Error fetching customer details: %v", err)
		http.Error(w, "Failed to fetch customer details", http.StatusInternalServerError)
		return
	}

	response := dto.CustomerResponse{
		Name: 		 customer.Name,
		DOB: 		 customer.DOB.Format("2006-01-02"),
		Telephone: 	 customer.PhoneNum,
		Email: 		 customer.Email,
		Nationality: customer.Nationality.NationalityName + " (" + customer.Nationality.NationalityCode + ")",
		Family: make([]struct {
			Relation string `json:"relation"`
			Name 	 string `json:"name"`
			DOB 	 string `json:"dob"`
		}, len(familyList)),
	}

	for i, family := range familyList {
		response.Family[i] = struct {
			Relation string `json:"relation"`
            Name 	 string `json:"name"`
            DOB 	 string `json:"dob"`
		}{
			Relation: family.Relation,
            Name: 	  family.Name,
            DOB: 	  family.DOB,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}