package models

type Child struct {
	Name   string `gorm:"primaryKey"`
	Parent string `gorm:"foreignKey:Username;references:Username", primaryKey`
}
