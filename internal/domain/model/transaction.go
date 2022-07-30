package model

import "time"

type Transaction struct {
	ID                   string    `json:"id" gorm:"default:gen_random_uuid()"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	Amount               Money     `json:"amount"`
	CreatedAt            time.Time `json:"created_at" sql:"type:timestamp without time zone"`
}
