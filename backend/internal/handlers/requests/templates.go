package requests

// @Description Request used for creating a new template
type CreateTemplateRequest struct {
	Name    string         `extensions:"x-order=0" validate:"required" json:"name" example:"user"`
	Content map[string]any `extensions:"x-order=1" validate:"required" json:"content" example:"{\"key\":\"{{variable}}\"}"`
}

// @Description Request used for generating data from selected template
type GenerateDataRequest struct {
	Count int `validate:"required" json:"count" example:"5"`
}
