package model

import "time"

type SensorRegister struct {
	ID          *uint             `gorm:"primary_key;auto_increment"`
	Version     *uint             `gorm:"column:version"`
	Name        *string           `gorm:"column:name"`
	Uuid        *string           `gorm:"column:uuid;unique"`
	Type        *string           `gorm:"column:type;type:varchar(255);not null"`
	Description *string           `gorm:"column:description"`
	Attributes  []SensorAttribute `gorm:"foreignkey:SensorRegisterID;association_foreignkey:ID;association_autoupdate:true"`
	CreatedAt   *time.Time        `gorm:"column:created_at"`
	ModifiedAt  *time.Time        `gorm:"column:modified_at"`
}
