package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetVehicleByExtId(obx *objectbox.ObjectBox, id string) (*models.Vehicle, models.ServerMessage) {

	var sm models.ServerMessage

	boxVehicle := models.BoxForVehicle(obx)

	queryVehicle := boxVehicle.Query(models.Vehicle_.ExtId.Equals(id, true))
	Vehicles, err := queryVehicle.Find()
	queryVehicle.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "Vehicle",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Vehicles) == 0 {
		sm = models.ServerMessage{
			DataType: "Vehicle",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Vehicles) != 1 {
		sm = models.ServerMessage{
			DataType: "Vehicle",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Vehicles[0], sm
}
