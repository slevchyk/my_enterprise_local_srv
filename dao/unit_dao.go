package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetUnitByExtId(obx *objectbox.ObjectBox, id string) (*models.Unit, models.ServerMessage) {

	var sm models.ServerMessage

	boxUnit := models.BoxForUnit(obx)

	queryUnit := boxUnit.Query(models.Unit_.ExtId.Equals(id, true))
	units, err := queryUnit.Find()
	queryUnit.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "unit",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(units) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "unit",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(units) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "unit",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return units[0], sm
}
