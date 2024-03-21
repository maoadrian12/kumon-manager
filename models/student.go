package models

type Student struct {
	Name            string `gorm:"primaryKey"`
	Parent_username string `gorm:"references:Username"`
}
