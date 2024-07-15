package datastore

import (
	"context"
	"database/sql"

	"github.com/wisnuuakbr/booking-to-go-golang/internal/entities/repository"
)

type FamilyListRepository interface {
	GetFamilyListByCustomerID(ctx context.Context, cstID int) ([]*repository.FamilyList, error)
}

type familyListRepository struct {
	DB *sql.DB
}

func NewFamilyListRepository(db *sql.DB) FamilyListRepository {
	return &familyListRepository{DB: db}
}

func (r *familyListRepository) GetFamilyListByCustomerID(ctx context.Context, cstID int) ([]*repository.FamilyList, error) {
	query := `
	SELECT 
		fl_id, cst_id, fl_name, fl_relation, fl_dob
	FROM family_list
	WHERE cst_id = $1`

	rows, err := r.DB.QueryContext(ctx, query, cstID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var familyMembers []*repository.FamilyList
	for rows.Next() {
		var familyMember repository.FamilyList
		if err := rows.Scan(&familyMember.FlID, &familyMember.CstID, &familyMember.Name, &familyMember.Relation, &familyMember.DOB); err != nil {
			return nil, err
		}
		familyMembers = append(familyMembers, &familyMember)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return familyMembers, nil
}
