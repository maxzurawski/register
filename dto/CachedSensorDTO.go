package dto

type CachedSensorDTO struct {
	Uuid        string `json:"uuid"`
	Active      bool   `json:"active"`
	Type        string `json:"type"`
	Max         string `json:"max"`
	Min         string `json:"min"`
	Nacta       int    `json:"nacta"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
