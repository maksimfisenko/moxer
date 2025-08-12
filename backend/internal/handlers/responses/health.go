package responses

// @Description Response used for the healthcheck requests
type HealthCheckResponse struct {
	Status string `validate:"required" json:"status" example:"ok"`
}
