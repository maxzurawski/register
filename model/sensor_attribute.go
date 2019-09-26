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

type SensorAttributes []SensorAttribute
type SensorAttributeFilter func(*SensorAttribute, string) bool

func (s SensorAttributes) FilterBy(fn SensorAttributeFilter, value string) *SensorAttribute {

	if s == nil || len(s) == 0 {
		return nil
	}

	for _, item := range s {
		if fn(&item, value) {
			return &item
		}
	}

	return nil
}
