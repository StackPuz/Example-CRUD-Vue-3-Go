package models

type Product struct {
	Id    int `gorm:"primaryKey;autoIncrement"`
	Name  string
	Price float64
}
