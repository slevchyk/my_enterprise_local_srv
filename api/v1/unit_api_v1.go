package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) UnitPost(w http.ResponseWriter, r *http.Request) {

	var aus []models.Unit
	var err error

	sa := models.ServerAnswer{
		Object:    "Unit",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &aus)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForUnit(apiV1.obx)

	for _, v := range aus {

		pd := models.ServerProcessedData{
			ExtId: v.ExtId,
		}

		isDataError := false

		if v.ExtId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "ext id is empty",
			})
			isDataError = true
		}

		if v.Name == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "name is empty",
			})
			isDataError = true
		}

		if isDataError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}
	

		query := box.Query(models.Unit_.ExtId.Equals(v.ExtId, true))
		Units, err := query.Find()
		query.Close()

		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query",
				Message: err.Error(),
			})

			pd.Status = http.StatusInternalServerError
			sa.ProcessedData = append(sa.ProcessedData, pd)

			query.Close()
			continue
		}

		if len(Units) == 0 {
			v.CreatedAt = time.Now().UTC()
			v.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&v)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "insert",
					Message: err.Error(),
				})
			}

		} else if len(Units) == 1 {
			v.Id = Units[0].Id
			v.CreatedAt = Units[0].CreatedAt
			v.UpdatedAt = time.Now().UTC()

			pd.SrvId = v.Id

			err := box.Update(&v)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "update",
					Message: err.Error(),
				})
			}
		} else {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "more than 1",
				Message: err.Error(),
			})
		}

		sa.ProcessedData = append(sa.ProcessedData, pd)
	}

	sa.Send(w)
}

func (api *ApiV1) UnitGet(w http.ResponseWriter, r *http.Request) {

	var aus []*models.Unit
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "Unit",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForUnit(api.obx)

	if fvId == "" {
		aus, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.Unit_.ExtId.Equals(fvId, true))
		aus, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			query.Close()
			return
		}
	}

	bs, err := json.Marshal(aus)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
