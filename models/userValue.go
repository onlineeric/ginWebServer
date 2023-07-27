package models

type UserValue struct {
	Value string `json:"value" binding:"required"`
}
