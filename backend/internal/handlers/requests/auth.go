package requests

// @Description Request containing user credentials used for registration and logging in
type CredentialsRequest struct {
	Email    string `extensions:"x-order=0" validate:"required" json:"email" example:"email@example.com"`
	Password string `extensions:"x-order=1" validate:"required" json:"password" example:"Str0ngPassWoRD"`
}
