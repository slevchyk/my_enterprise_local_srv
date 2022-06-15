package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetLocalityByExtId(obx *objectbox.ObjectBox, id string) (*models.Locality, models.ServerMessage) {

	var sm models.ServerMessage

	boxLocality := models.BoxForLocality(obx)

	queryLocality := boxLocality.Query(models.Locality_.ExtId.Equals(id, true))
	Localitys, err := queryLocality.Find()
	queryLocality.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "Locality",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Localitys) == 0 {
		sm = models.ServerMessage{
			DataType: "Locality",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Localitys) != 1 {
		sm = models.ServerMessage{
			DataType: "Locality",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Localitys[0], sm
}
