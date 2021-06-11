package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetAppUserByExtId(obx *objectbox.ObjectBox, id string) (*models.AppUser, models.ServerMessage) {

	var sm models.ServerMessage

	boxAppUser := models.BoxForAppUser(obx)

	queryAppUser := boxAppUser.Query(models.AppUser_.ExtId.Equals(id, true))
	appUsers, err := queryAppUser.Find()
	queryAppUser.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "app user",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(appUsers) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "app user",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(appUsers) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "app user",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return appUsers[0], sm
}
