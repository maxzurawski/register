package dbprovider

import (
	"errors"
	"fmt"

	"github.com/maxzurawski/register/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) DeleteSensor(uuid string) (uint, error) {
	if stringutils.IsZero(uuid) {
		return 0, errors.New("given uuid is empty")
	}
	if valid := stringutils.IsUuidValid(uuid); !valid {
		return 0, errors.New(fmt.Sprintf("given uuid is invalid [%s]", uuid))
	}
	_, err := Mgr.GetSensorByUuid(uuid)
	if err != nil {
		return 0, err
	}
	err = mgr.GetDb().Unscoped().Where("Uuid=?", uuid).Delete(model.SensorRegister{}).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}
