package models

import (
	"gorm.io/gorm"
)

// be careful with low bar on the fields name
type Task struct {
	gorm.Model        //Convert to table
	Fecha      string `gorm:"type:date"`
	Titulo     string
	Autor      string
}
