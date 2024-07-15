package repository

import "time"

type Customer struct {
	CstID         int
	Nationality   *Nationality
	Name          string
	DOB           time.Time
	PhoneNum      string
	Email         string
}