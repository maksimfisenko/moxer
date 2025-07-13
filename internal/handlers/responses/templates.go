package responses

import "time"

// @Description Template response
type Template struct {
	Id        string         `extensions:"x-order=0" validate:"required" json:"id" example:"7975a7ec-bfda-42ad-831d-0b250277e402"`
	Name      string         `extensions:"x-order=1" validate:"required" json:"name" example:"user"`
	Content   map[string]any `extensions:"x-order=2" validate:"required" json:"content"`
	CreatedAt time.Time      `extensions:"x-order=3" validate:"required" json:"created_at" example:"2025-07-09T18:43:23.239168298+03:00"`
	UpdatedAt time.Time      `extensions:"x-order=4" validate:"required" json:"updated_at" example:"2025-07-09T18:43:23.239171581+03:00"`
}

// @Description Generated Data response
type GeneratedData struct {
	Data []map[string]any `extensions:"x-order=0" validate:"required" json:"data"`
}
