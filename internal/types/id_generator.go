package types

import "time"

type IdGenerator struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Current  int64     `json:"current"`
	Modified time.Time `json:"modified"`
}

func (i IdGenerator) TableName() string {
	return "id_generator"
}
