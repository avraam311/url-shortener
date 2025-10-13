package dto

type FullURL struct {
	URL string `json:"url" validate:"required"`
}
