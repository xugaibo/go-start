package models

import "time"

type Base struct {
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	CreatedByName string    `json:"created_by_name"`
	UpdatedByName string    `json:"updated_by_name"`
	IsDelete      int       `json:"is_delete"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
