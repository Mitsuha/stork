package model

type Setting struct {
	Key   string `json:"key" gorm:"column:key;primary_key"`
	Value string `json:"value" gorm:"column:value"`
}
