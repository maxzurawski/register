package dto

type SensorAttributeDTO struct {
	ID               uint   `json:"id"`
	Version          uint   `json:"version"`
	Symbol           string `json:"symbol"`
	Value            string `json:"value"`
	SensorRegisterID uint   `json:"sensor_id"`
}
