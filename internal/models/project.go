package models

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
}
