package model

import "time"

type SensorAttribute struct {
	ID               *uint      `gorm:"primary_key;auto_increment"`
	Version          *uint      `gorm:"column:version"`
	Attribute        Attribute  `gorm:"foreignkey:RefSymbol;association_autoupdate:false"`
	RefSymbol        *string    `gorm:"column:ref_symbol"`
	Value            *string    `gorm:"column:value;type:varchar(255)"`
	SensorRegisterID *uint      `gorm:"column:sensor_register_id"`
	CreateAt         *time.Time `gorm:"column:created_at"`
	ModifiedAt       *time.Time `gorm:"column:modified_at"`
}
