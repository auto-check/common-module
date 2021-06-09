package model

import "time"

type Log struct {
	ID int64 `json:"id"`
	WorkID int64 `json:"work_id"`
	Date time.Time `json:"date"`
	Message string `json:"message"`
	Success bool `json:"success"`
}