package usecases

import (
	"context"

	"github.com/wisnuuakbr/booking-to-go-golang/internal/entities/repository"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/infrastructure/datastore"
)

type CustomerUseCase struct {
	customerRepo datastore.CustomerRepository
	familyListRepo datastore.FamilyListRepository
	nationalityRepo datastore.NationalityRepository
}

func NewCustomerUseCase(customerRepo datastore.CustomerRepository, familyListRepo datastore.FamilyListRepository, nationalityRepo datastore.NationalityRepository) *CustomerUseCase {
	return &CustomerUseCase{
        customerRepo:     customerRepo,
        familyListRepo:  familyListRepo,
        nationalityRepo: nationalityRepo,
    }
}

func (uc *CustomerUseCase) GetCustomer(ctx context.Context, cstID int) (*repository.Customer, []*repository.FamilyList, error) {
	
	customer, err := uc.customerRepo.GetCustomerByID(ctx, cstID)
	if err != nil {
		return nil, nil, err
	}

	familyList, err := uc.familyListRepo.GetFamilyListByCustomerID(ctx, cstID)
	if err != nil {
		return nil, nil, err
	}

	return customer, familyList, nil
}