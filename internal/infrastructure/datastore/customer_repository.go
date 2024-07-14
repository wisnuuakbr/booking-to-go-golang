package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wisnuuakbr/booking-to-go-golang/internal/entities/repository"
)

type CustomerRepository interface {
	GetCustomerByID(ctx context.Context, id int) (*repository.Customer, error)
}

type customerRepository struct {
    DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) *customerRepository {
    return &customerRepository{
		DB: db,
	}
}

func (r *customerRepository) GetCustomerByID(ctx context.Context, id int) (*repository.Customer, error) {
	query := `
	SELECT 
			a.cst_id, a.nationality_id, a.cst_name, a.cst_dob, a.cst_phoneNum, a.cst_email,
			b.nationality_id, b.nationality_name, b.nationality_code
	FROM customer a
	LEFT JOIN Nationality b ON a.nationality_id = b.nationality_id
	WHERE a.cst_id = $1`

	row := r.DB.QueryRowContext(ctx, query, id)

	var customer repository.Customer
	var nationality repository.Nationality

	err := row.Scan(
		&customer.CstID,
		&customer.Nationality.NationalityID,
		&customer.Name,
		&customer.DOB,
		&customer.PhoneNum,
		&customer.Email,
		&nationality.NationalityID,
		&nationality.NationalityName,
		&nationality.NationalityCode,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("data customer not found")
		}
		return nil, err
	}

	customer.Nationality = &nationality
	return &customer, nil
}