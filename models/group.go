package models

type Group struct {
	Id        string `json:"id" yaml:"id" sql:"primary_key;column:id"`
	Name      string `json:"name" yaml:"name" sql:"column:name"`
	CreatedAt int64  `json:"created_at" yaml:"-" sql:"column:created_at"`
}
