package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetPersonByExtId(obx *objectbox.ObjectBox, id string) (*models.Person, models.ServerMessage) {

	var sm models.ServerMessage

	boxPerson := models.BoxForPerson(obx)

	queryPerson := boxPerson.Query(models.Person_.ExtId.Equals(id, true))
	Persons, err := queryPerson.Find()
	queryPerson.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "person",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Persons) == 0 {
		sm = models.ServerMessage{
			DataType: "person",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Persons) != 1 {
		sm = models.ServerMessage{
			DataType: "person",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Persons[0], sm
}
