package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetConsignmentNoteInByExtId(obx *objectbox.ObjectBox, id string) (*models.ConsignmentNoteIn, models.ServerMessage) {

	var sm models.ServerMessage

	boxConsignmentNoteIn := models.BoxForConsignmentNoteIn(obx)

	queryConsignmentNoteIn := boxConsignmentNoteIn.Query(models.ConsignmentNoteIn_.ExtId.Equals(id, true))
	ConsignmentNoteIns, err := queryConsignmentNoteIn.Find()
	queryConsignmentNoteIn.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(ConsignmentNoteIns) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(ConsignmentNoteIns) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "ConsignmentNoteIn",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return ConsignmentNoteIns[0], sm
}
