package types

import "time"

type UrlRecord struct {
	Id      int64     `json:"id"`
	Url     string    `json:"url"`
	Code    int64     `json:"code"`
	Ttl     int32     `json:"ttl"`
	Created time.Time `json:"created"`
}

func (u UrlRecord) TableName() string {
	return "url_record"
}
