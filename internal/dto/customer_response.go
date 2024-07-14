package dto

type CustomerResponse struct {
	Name        string `json:"name"`
	DOB         string `json:"dob"`
	Telephone   string `json:"telephone"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	Family      []struct {
		Relation string `json:"relation"`
		Name     string `json:"name"`
		DOB      string `json:"dob"`
	} `json:"family"`
}