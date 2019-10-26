package dto

import "github.com/maxzurawski/utilities/stringutils"

type SensorAttributeDTO struct {
	ID               uint                    `json:"id"`
	Version          uint                    `json:"version"`
	Symbol           string                  `json:"symbol"`
	Value            stringutils.MultiString `json:"value"`
	SensorRegisterID uint                    `json:"sensor_id"`
}
