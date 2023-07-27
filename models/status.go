package models

type Status struct {
	Status string `json:"status" binding:"required"`
	Value  string `json:"value,omitempty"`
	Error  string `json:"error,omitempty"`
}
