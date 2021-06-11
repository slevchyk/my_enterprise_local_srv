package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetGoodsByExtId(obx *objectbox.ObjectBox, id string) (*models.Goods, models.ServerMessage) {

	var sm models.ServerMessage

	boxGoods := models.BoxForGoods(obx)

	queryGoods := boxGoods.Query(models.Goods_.ExtId.Equals(id, true))
	Goodss, err := queryGoods.Find()
	queryGoods.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "goods",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Goodss) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "goods",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Goodss) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "goods",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Goodss[0], sm
}
