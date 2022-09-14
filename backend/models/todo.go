package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Todo `json:"items"`
}
