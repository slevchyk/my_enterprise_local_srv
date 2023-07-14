package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) SubdivisionPost(w http.ResponseWriter, r *http.Request) {

	var gs []models.Subdivision
	var err error

	sa := models.ServerAnswer{
		Object:    "Subdivision",
		WebMethod: "post",
		DateUTC:   time.Now()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	var jsonData []map[string]interface{}
	err = json.Unmarshal(bs, &jsonData)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForSubdivision(apiV1.obx)
	boxLocality := models.BoxForLocality(apiV1.obx)

	for _, v := range jsonData {

		jsonError := false

		var pd models.ServerProcessedData

		extId, ok := v["ext_id"].(string)
		if !ok || extId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value ext_id",
				Message: "ext_id: incorrect type or empty",
			})
			jsonError = true
		}
		pd.ExtId = extId

		name, ok := v["name"].(string)
		if !ok || name == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "name: incorrect type or empty",
			})
			jsonError = true
		}

		locality_ext_id, ok := v["locality_ext_id"].(string)
		if !ok || locality_ext_id == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "locality_ext_id: incorrect type or empty",
			})
			jsonError = true
		}

		if jsonError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		queryLocality := boxLocality.Query(models.Locality_.ExtId.Equals(locality_ext_id, true))
		localities, err := queryLocality.Find()
		queryLocality.Close()

		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "locality",
				DataId:   locality_ext_id,
				Action:   "query",
				Message:  err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		if len(localities) == 0 {
			pd.Status = http.StatusNotFound
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "locality",
				DataId:   locality_ext_id,
				Action:   "query",
				Message:  "not found",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		} else if len(localities) != 1 {
			pd.Status = http.StatusConflict
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "locality",
				DataId:   locality_ext_id,
				Action:   "query",
				Message:  "more than 1",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		locality := localities[0]

		query := box.Query(models.Subdivision_.ExtId.Equals(extId, true))
		Subdivisions, err := query.Find()
		query.Close()
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			query.Close()
			continue
		}

		Subdivision := models.Subdivision{
			ExtId:    extId,
			Name:     name,
			Locality: locality,
			IsDeleted: v["is_deleted"].(bool),
		}

		if len(Subdivisions) == 0 {
			Subdivision.CreatedAt = time.Now().UTC()
			Subdivision.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&Subdivision)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "insert",
					Message: err.Error(),
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

		} else if len(Subdivisions) == 1 {
			Subdivision.Id = Subdivisions[0].Id
			Subdivision.CreatedAt = Subdivisions[0].CreatedAt
			Subdivision.UpdatedAt = time.Now().UTC()

			err := box.Update(&Subdivision)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "update",
					Message: err.Error(),
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}
		} else {
			pd.Status = http.StatusConflict
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "more than 1",
				Message: err.Error(),
			})
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		pd.Status = http.StatusOK
		sa.ProcessedData = append(sa.ProcessedData, pd)
	}

	err = json.Unmarshal(bs, &gs)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	sa.Send(w)
}

func (api *ApiV1) SubdivisionGet(w http.ResponseWriter, r *http.Request) {

	var ss []*models.Subdivision
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "Subdivision",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForSubdivision(api.obx)

	if fvId == "" {
		ss, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.Subdivision_.ExtId.Equals(fvId, true))
		ss, err = query.Find()
		query.Close()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			query.Close()
			return
		}
	}

	bs, err := json.Marshal(ss)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
