package responses

// @Description Healthcheck response
type HealthcheckResponse struct {
	Status string `json:"status" validate:"required" example:"ok"`
}
