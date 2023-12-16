package models

import "gorm.io/gorm"

type SensorData struct {
	gorm.Model
	Water  int    `json:"water"`
	Wind   int    `json:"wind"`
	Status string `json:"status"`
}
