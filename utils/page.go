package utils

type PageRequest struct {
	PageSize   uint `json:"page_size"`
	PageNumber uint `json:"page_number"`
}
