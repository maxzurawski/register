package model

import (
	"strconv"
	"strings"
	"time"

	"github.com/maxzurawski/utilities/symbols"
)

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

type SensorsArray []SensorRegister
type SensorRegisterFilter func(*SensorRegister, string) bool

func (sa SensorsArray) FilterBy(fn SensorRegisterFilter, value string) *SensorRegister {

	if sa == nil || len(sa) == 0 {
		return nil
	}

	for _, item := range sa {
		if fn(&item, value) {
			return &item
		}
	}
	return nil
}

func (r *SensorRegister) IsActive() bool {
	for _, item := range r.Attributes {
		if strings.ToUpper(*item.RefSymbol) == strings.ToUpper(symbols.Active.String()) {
			if value, err := strconv.ParseBool(*item.Value); err != nil {
				return false
			} else {
				return value
			}
		}
	}
	return true
}

func (r *SensorRegister) GetAttributeAsString(symbol string) string {
	for _, item := range r.Attributes {
		if strings.ToUpper(*item.RefSymbol) == strings.ToUpper(symbol) {
			return *item.Value
		}
	}
	return ""
}

func (r *SensorRegister) GetAttributeAsInt(symbol string) int {
	for _, item := range r.Attributes {
		if strings.ToUpper(*item.RefSymbol) == strings.ToUpper(symbol) {
			if value, err := strconv.Atoi(*item.Value); err != nil {
				return -1
			} else {
				return value
			}
		}
	}
	return -1
}
