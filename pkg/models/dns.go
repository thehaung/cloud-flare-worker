package models

type DNS struct {
	IpAddress string `json:"ip_address" bson:"ip_address"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}
