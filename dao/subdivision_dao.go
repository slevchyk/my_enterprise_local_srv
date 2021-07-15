package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetSubdivisionByExtId(obx *objectbox.ObjectBox, id string) (*models.Subdivision, models.ServerMessage) {

	var sm models.ServerMessage

	boxGoodsGroup := models.BoxForSubdivision(obx)

	queryGoodsGroup := boxGoodsGroup.Query(models.GoodsGroup_.ExtId.Equals(id, true))
	GoodsGroups, err := queryGoodsGroup.Find()
	queryGoodsGroup.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "Subdivision",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(GoodsGroups) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "Subdivision",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(GoodsGroups) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "Subdivision",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return GoodsGroups[0], sm
}
