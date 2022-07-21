package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/objectbox/objectbox-go/objectbox"
)

type ApiV1 struct {
	obx *objectbox.ObjectBox
	abn int //Application Build Number
}

func NewApiV1(obx *objectbox.ObjectBox) *ApiV1 {
	return &ApiV1{obx: obx}
}

func (api *ApiV1) SetAbn(r *http.Request) {

	fvBuildNumber := r.FormValue("build_number")

	if fvBuildNumber == "" {
		api.abn = 0
		return
	}

	abn, err := strconv.Atoi(fvBuildNumber)
	if err != nil {
		api.abn = 0
		return
	}

	api.abn = abn
}

func parseDate(date string, abn int) (time.Time, error) {

	if abn < 5 {
		return time.Parse("2006-01-02T15:04:05", date)
	} else {
		return time.Parse("2006-01-02T15:04:05Z", date)
	}

}
