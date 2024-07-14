package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wisnuuakbr/booking-to-go-golang/internal/entities/repository"
)

type NationalityRepository interface {
	GetNationalityByID(ctx context.Context, id int) (*repository.Nationality, error)
}

type nationalityRepository struct {
	DB *sql.DB
}

func NewNationalityRepository(db *sql.DB) NationalityRepository {
	return &nationalityRepository{DB: db}
}

func (r *nationalityRepository) GetNationalityByID(ctx context.Context, id int) (*repository.Nationality, error) {
	query := `
	SELECT 
		nationality_id, nationality_name, nationality_code
	FROM Nationality
	WHERE nationality_id = $1`

	row := r.DB.QueryRowContext(ctx, query, id)

	var nationality repository.Nationality
	err := row.Scan(&nationality.NationalityID, &nationality.NationalityName, &nationality.NationalityCode)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nationality not found")
		}
		return nil, err
	}

	return &nationality, nil
}
