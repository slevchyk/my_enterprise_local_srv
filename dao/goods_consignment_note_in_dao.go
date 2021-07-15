package dao

import (
	"net/http"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetGoodsConsignmentNoteInByExtId(obx *objectbox.ObjectBox, id string) (*models.GoodsConsignmentNoteIn, models.ServerMessage) {

	var sm models.ServerMessage

	boxGoodsConsignmentNoteIn := models.BoxForGoodsConsignmentNoteIn(obx)

	queryGoodsConsignmentNoteIn := boxGoodsConsignmentNoteIn.Query(models.GoodsConsignmentNoteIn_.ExtId.Equals(id, true))
	GoodsConsignmentNoteIns, err := queryGoodsConsignmentNoteIn.Find()
	queryGoodsConsignmentNoteIn.Close()

	if err != nil {
		sm = models.ServerMessage{
			Status:   http.StatusInternalServerError,
			DataType: "GoodsConsignmentNoteIn",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(GoodsConsignmentNoteIns) == 0 {
		sm = models.ServerMessage{
			Status:   http.StatusNotFound,
			DataType: "GoodsConsignmentNoteIn",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(GoodsConsignmentNoteIns) != 1 {
		sm = models.ServerMessage{
			Status:   http.StatusBadRequest,
			DataType: "GoodsConsignmentNoteIn",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return GoodsConsignmentNoteIns[0], sm
}
