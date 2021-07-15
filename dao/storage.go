package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetStorageByExtId(obx *objectbox.ObjectBox, id string) (*models.Storage, models.ServerMessage) {

	var sm models.ServerMessage

	boxStorage := models.BoxForStorage(obx)

	queryStorage := boxStorage.Query(models.Storage_.ExtId.Equals(id, true))
	Storages, err := queryStorage.Find()
	queryStorage.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "Storage",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Storages) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "Storage",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Storages) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "Storage",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Storages[0], sm
}
