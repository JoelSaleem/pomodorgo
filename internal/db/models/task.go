package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID		  int
	Name	  string
	Description string
}