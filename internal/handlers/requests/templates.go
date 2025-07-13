package requests

// @Description Create Template Request
type CreateTemplateRequest struct {
	Name    string         `extensions:"x-order=0" validate:"required" json:"name" example:"user"`
	Content map[string]any `extensions:"x-order=1" validate:"required" json:"content" example:""`
}

// @Description Generate Data Request
type GenerateDataRequest struct {
	Count int `extensions:"x-order=0" validate:"required" json:"count" example:"5"`
}
