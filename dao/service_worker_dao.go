package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetServiceWorkerByExtId(obx *objectbox.ObjectBox, id string) (*models.ServiceWorker, models.ServerMessage) {

	var sm models.ServerMessage

	boxServiceWorker := models.BoxForServiceWorker(obx)

	queryServiceWorker := boxServiceWorker.Query(models.ServiceWorker_.ExtId.Equals(id, true))
	ServiceWorkers, err := queryServiceWorker.Find()
	queryServiceWorker.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "ServiceWorker",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(ServiceWorkers) == 0 {
		sm = models.ServerMessage{
			DataType: "ServiceWorker",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(ServiceWorkers) != 1 {
		sm = models.ServerMessage{
			DataType: "ServiceWorker",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return ServiceWorkers[0], sm
}
