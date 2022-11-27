package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Placa   string
	Chassi  string
	Renavam int
	Modelo  string
	Marca   string
	Ano     int
}
