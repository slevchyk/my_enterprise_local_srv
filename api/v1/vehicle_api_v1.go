package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) VehiclePost(w http.ResponseWriter, r *http.Request) {

	var vis []models.VehicleImport
	var err error

	sa := models.ServerAnswer{
		Object:    "Vehicle",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := io.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &vis)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForVehicle(apiV1.obx)

	for _, vi := range vis {

		pd := models.ServerProcessedData{
			ExtId: vi.ExtId,
		}

		isDataError := false

		if vi.ExtId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "ext id is empty",
			})
			isDataError = true
		}

		if vi.Name == "" {
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

		query := box.Query(models.Vehicle_.ExtId.Equals(vi.ExtId, true))
		Vehicles, err := query.Find()
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

		trailer, _ := dao.GetTrailerByExtId(apiV1.obx, vi.DefTrailerExtId)
		driver, _ := dao.GetServiceWorkerByExtId(apiV1.obx, vi.DefDriverExtId)

		v := models.Vehicle{

			ExtId:      vi.ExtId,
			Name:       vi.Name,
			IsDeleted:  vi.IsDeleted,
			Length:     float64(vi.Length),
			Width:      float64(vi.Width),
			Height:     float64(vi.Height),
			MinWeight:  float64(vi.MinWeight),
			MaxWeight:  float64(vi.MaxWeight),
			Comment:    vi.Comment,
			PhotoPath:  vi.PhotoPath,
			NfcId:      vi.NfcId,
			DefTrailer: trailer,
			DefDriver:  driver,
		}

		if len(Vehicles) == 0 {
			v.CreatedAt = time.Now().UTC()
			v.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&v)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "insert",
					Message: err.Error(),
				})
				pd.Status = http.StatusInternalServerError
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

		} else if len(Vehicles) == 1 {
			v.Id = Vehicles[0].Id
			v.CreatedAt = Vehicles[0].CreatedAt
			v.UpdatedAt = time.Now().UTC()

			pd.SrvId = v.Id

			err := box.Update(&v)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "update",
					Message: err.Error(),
				})
				pd.Status = http.StatusInternalServerError
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}
		} else {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "select",
				Message: "more than 1",
			})
			pd.Status = http.StatusConflict
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		pd.Status = http.StatusOK
		sa.ProcessedData = append(sa.ProcessedData, pd)
	}

	sa.Send(w)
}

func (api *ApiV1) VehicleGet(w http.ResponseWriter, r *http.Request) {

	var vs []*models.Vehicle
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "Vehicle",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForVehicle(api.obx)

	if fvId == "" {
		vs, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.Vehicle_.ExtId.Equals(fvId, true))
		vs, err = query.Find()
		query.Close()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			query.Close()
			return
		}
	}

	bs, err := json.Marshal(vs)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
