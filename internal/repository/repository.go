package repository

type Pagination struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
	Count  uint `json:"count"`
}
