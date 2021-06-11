package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetGoodsGroupByExtId(obx *objectbox.ObjectBox, id string) (*models.GoodsGroup, models.ServerMessage) {

	var sm models.ServerMessage

	boxGoodsGroup := models.BoxForGoodsGroup(obx)

	queryGoodsGroup := boxGoodsGroup.Query(models.GoodsGroup_.ExtId.Equals(id, true))
	GoodsGroups, err := queryGoodsGroup.Find()
	queryGoodsGroup.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "goods group",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(GoodsGroups) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "goods group",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(GoodsGroups) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "goods group",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return GoodsGroups[0], sm
}
