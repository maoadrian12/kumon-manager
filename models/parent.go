package models

type Parents struct {
	Username string `gorm:"primaryKey"`
	Name     string
	Pass     string
}
