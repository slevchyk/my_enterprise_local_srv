package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetHarvestStatusByExtId(obx *objectbox.ObjectBox, id string) (*models.HarvestStatus, models.ServerMessage) {

	var sm models.ServerMessage

	boxHarvestStatus := models.BoxForHarvestStatus(obx)

	queryHarvestStatus := boxHarvestStatus.Query(models.HarvestStatus_.ExtId.Equals(id, true))
	HarvestStatuss, err := queryHarvestStatus.Find()
	queryHarvestStatus.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "HarvestStatus",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(HarvestStatuss) == 0 {
		sm = models.ServerMessage{
			DataType: "HarvestStatus",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(HarvestStatuss) != 1 {
		sm = models.ServerMessage{
			DataType: "HarvestStatus",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return HarvestStatuss[0], sm
}
