package models

import (
	"time"

	"gorm.io/gorm"
)

// be careful with low bar on the fields name
type Task struct {
	gorm.Model           //Convert to table
	Fecha      time.Time `gorm:"type:date"` // Utiliza time.Time en lugar de now.Now
	Titulo     string
	Autor      string
}
