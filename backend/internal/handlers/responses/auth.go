package responses

import "time"

// @Description Response containing information about user (without password)
type UserResponse struct {
	Id        string    `extensions:"x-order=0" validate:"required" json:"id" example:"7975a7ec-bfda-42ad-831d-0b250277e402"`
	Email     string    `extensions:"x-order=1" validate:"required" json:"email" example:"email@example.com"`
	CreatedAt time.Time `extensions:"x-order=2" validate:"required" json:"created_at" example:"2025-07-09T18:43:23.239168298+03:00"`
	UpdatedAt time.Time `extensions:"x-order=3" validate:"required" json:"updated_at" example:"2025-07-09T18:43:23.239171581+03:00"`
}

// @Description Response containing a unique generated JWT token
type Token struct {
	Token string `validate:"required" json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"`
}
