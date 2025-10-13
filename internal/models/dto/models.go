package dto

type Request struct {
	FullURL string `json:"full_url" validate:"required"`
}
