package model

import "time"

type Invoice struct {
	ID            int64     `json:"id"`
	UserBillingId int64     `json:"userBillingId,omitempty"`
	Number        string    `json:"number,omitempty"`
	FilePath      string    `json:"filePath"`
	CreatedAt     time.Time `json:"createdAt"`
}
