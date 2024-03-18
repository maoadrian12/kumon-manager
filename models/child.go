package models

type Child struct {
	Username string `gorm:"primaryKey"`
	Name     string
	Parent   string `gorm:"foreignKey:Username;references:Username", primaryKey`
	Pass     string
}
