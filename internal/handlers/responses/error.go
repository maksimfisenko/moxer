package responses

// @Description Error response
type ErrorResponse struct {
	Error string `json:"error" validate:"required" example:"error message"`
}
