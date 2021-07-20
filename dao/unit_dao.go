package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetUnitByExtId(obx *objectbox.ObjectBox, id string) (*models.Unit, models.ServerMessage) {

	var sm models.ServerMessage

	if id == "" {
		sm = models.ServerMessage{
			DataType: "unit",
			DataId:   id,
			Action:   "query",
			Message:  "an empty id",
		}
		return nil, sm
	}

	boxUnit := models.BoxForUnit(obx)

	queryUnit := boxUnit.Query(models.Unit_.ExtId.Equals(id, true))
	units, err := queryUnit.Find()
	queryUnit.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "unit",
			DataId:   id,
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(units) == 0 {
		sm = models.ServerMessage{
			DataType: "unit",
			DataId:   id,
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(units) != 1 {
		sm = models.ServerMessage{
			DataType: "unit",
			DataId:   id,
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return units[0], sm
}
