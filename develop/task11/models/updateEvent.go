package models

type UpdEvent struct {
	UserID  int    `json:"user_id"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Uid     int64  `json:"uid"`
	NewData Event  `json:"new_data"`
}
