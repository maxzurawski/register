package dto

type SensorRegisterDTO struct {
	ID          uint                 `json:"id"`
	Version     uint                 `json:"version"`
	Name        string               `json:"name"`
	Uuid        string               `json:"uuid"`
	Type        string               `json:"type"`
	Description string               `json:"description"`
	Attributes  []SensorAttributeDTO `json:"attributes"`
}
