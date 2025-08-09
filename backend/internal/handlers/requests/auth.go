package requests

// @Description Register request
type RegisterRequest struct {
	Email    string `extensions:"x-order=0" validate:"required" json:"email" example:"fisenkomaksim.id@gmail.com"`
	Password string `extensions:"x-order=1" validate:"required" json:"password" example:"11111111"`
}

// @Description Login request
type LoginRequest struct {
	Email    string `extensions:"x-order=0" validate:"required" json:"email" example:"fisenkomaksim.id@gmail.com"`
	Password string `extensions:"x-order=1" validate:"required" json:"password" example:"11111111"`
}
