package models

type SignupDetail struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password" validate:"min=8,max=20"`
	Phone     string `json:"phone"`
}
type SignupDetailResponse struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
type TokenUser struct {
	Users        SignupDetailResponse
	AccessToken  string
	RefreshToken string
}
