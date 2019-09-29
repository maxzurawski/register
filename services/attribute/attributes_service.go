package attribute

import (
	"github.com/xdevices/register/dbprovider"
	"github.com/xdevices/register/dto"
)

type AttributesService interface {
	GetAll() ([]dto.AttributeDTO, error)
}

var Service AttributesService

type service struct {
	mgr dbprovider.DBManager
}

func Init() {
	s := service{}
	s.mgr = dbprovider.Mgr
	Service = &s
}