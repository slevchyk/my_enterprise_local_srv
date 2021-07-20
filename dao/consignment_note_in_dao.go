package dao

import (
	"fmt"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetConsignmentNoteInById(obx *objectbox.ObjectBox, id uint64) (*models.ConsignmentNoteIn, models.ServerMessage) {

	var sm models.ServerMessage

	if id == 0 {
		sm = models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			DataId: fmt.Sprint(id),
			Action:   "query",
			Message:  "id = 0",
		}
		return nil, sm
	} 

	boxConsignmentNoteIn := models.BoxForConsignmentNoteIn(obx)

	ConsignmentNoteIn, err := boxConsignmentNoteIn.Get(id)
	
	if err != nil {
		sm = models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			DataId: fmt.Sprint(id),
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	return ConsignmentNoteIn, sm
}

func GetConsignmentNoteInByExtId(obx *objectbox.ObjectBox, id string) (*models.ConsignmentNoteIn, models.ServerMessage) {

	var sm models.ServerMessage

	boxConsignmentNoteIn := models.BoxForConsignmentNoteIn(obx)

	queryConsignmentNoteIn := boxConsignmentNoteIn.Query(models.ConsignmentNoteIn_.ExtId.Equals(id, true))
	ConsignmentNoteIns, err := queryConsignmentNoteIn.Find()
	queryConsignmentNoteIn.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(ConsignmentNoteIns) == 0 {
		sm = models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(ConsignmentNoteIns) != 1 {
		sm = models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return ConsignmentNoteIns[0], sm
}
