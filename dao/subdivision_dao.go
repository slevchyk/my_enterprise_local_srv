package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetSubdivisionByExtId(obx *objectbox.ObjectBox, id string) (*models.Subdivision, models.ServerMessage) {

	var sm models.ServerMessage

	if id == "" {
		sm = models.ServerMessage{
			DataType: "Subdivision",
			DataId: id,
			Action:   "query",
			Message:  "an empty id",
		}
		return nil, sm
	}

	boxSubdivision := models.BoxForSubdivision(obx)

	querySubdivision := boxSubdivision.Query(models.Subdivision_.ExtId.Equals(id, true))
	Subdivisions, err := querySubdivision.Find()
	querySubdivision.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "Subdivision",
			DataId: id,
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Subdivisions) == 0 {
		sm = models.ServerMessage{
			DataType: "Subdivision",
			DataId: id,
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Subdivisions) != 1 {
		sm = models.ServerMessage{
			DataType: "Subdivision",
			DataId: id,
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Subdivisions[0], sm
}
