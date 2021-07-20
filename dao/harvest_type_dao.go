package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetHarvestTypeByExtId(obx *objectbox.ObjectBox, id string) (*models.HarvestType, models.ServerMessage) {

	var sm models.ServerMessage

	boxHarvestType := models.BoxForHarvestType(obx)

	queryHarvestType := boxHarvestType.Query(models.HarvestType_.ExtId.Equals(id, true))
	HarvestTypes, err := queryHarvestType.Find()
	queryHarvestType.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "harvest type",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(HarvestTypes) == 0 {
		sm = models.ServerMessage{
			DataType: "harvest type",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(HarvestTypes) != 1 {
		sm = models.ServerMessage{
			DataType: "harvest type",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return HarvestTypes[0], sm
}
