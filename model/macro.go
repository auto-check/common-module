package model

type Macro struct {
	ID int64 `json:"id"`
	StudentID int64 `json:"student_id"`
	Birth string `json:"birth"`
	Password string `json:"password"`
	OnOff bool `json:"onoff"`
}