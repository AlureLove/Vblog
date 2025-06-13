package utils

type PageRequest struct {
	PageSize   uint `json:"page_size"`
	PageNumber uint `json:"page_number"`
}

func NewPageRequest() *PageRequest {
	return &PageRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

func (p *PageRequest) Offset() uint {
	return (p.PageNumber - 1) * p.PageSize
}
