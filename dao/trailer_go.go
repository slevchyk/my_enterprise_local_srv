package dao

import (
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func GetTrailerByExtId(obx *objectbox.ObjectBox, id string) (*models.Trailer, models.ServerMessage) {

	var sm models.ServerMessage

	boxTrailer := models.BoxForTrailer(obx)

	queryTrailer := boxTrailer.Query(models.Trailer_.ExtId.Equals(id, true))
	Trailers, err := queryTrailer.Find()
	queryTrailer.Close()

	if err != nil {
		sm = models.ServerMessage{
			DataType: "Trailer",
			Action:   "query",
			Message:  err.Error(),
		}
		return nil, sm
	}

	if len(Trailers) == 0 {
		sm = models.ServerMessage{
			DataType: "Trailer",
			Action:   "query",
			Message:  "not found",
		}
		return nil, sm

	} else if len(Trailers) != 1 {
		sm = models.ServerMessage{
			DataType: "Trailer",
			Action:   "query",
			Message:  "more than 1",
		}
		return nil, sm
	}

	return Trailers[0], sm
}
