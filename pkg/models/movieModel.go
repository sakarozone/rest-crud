package model


type MovieTable struct {
	ID        uint           `gorm:"primaryKey"`
	Name string
	Year uint
	Director string
	Rating uint 
  }

