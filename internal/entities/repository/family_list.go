package repository

import "time"

type FamilyList struct {
	FlID     int
	CstID    int
	Relation string
	DOB      time.Time
}