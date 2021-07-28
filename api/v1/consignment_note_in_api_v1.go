package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) ConsignmentNoteInPost(w http.ResponseWriter, r *http.Request) {

	var cnis []*models.ConsignmentNoteIn
	var err error

	// var sas []models.ServerAnswer
	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

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

	box := models.BoxForConsignmentNoteIn(apiV1.obx)

	for _, v := range jsonData {

		var pd models.ServerProcessedData

		isJsonError := false

		srvIdFloat64, ok := v["srv_id"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking type",
				Message: "srv_id: incorrect type",
			})
			isJsonError = true
		}
		srvId := uint64(srvIdFloat64)
		pd.SrvId = srvId

		appId, ok := v["app_id"].(string)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking type",
				Message: "app_id: incorrect type",
			})
			isJsonError = true
		}
		pd.AppId = appId

		extId, ok := v["ext_id"].(string)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking type",
				Message: "ext_id: incorrect type",
			})
			isJsonError = true
		}
		pd.ExtId = extId

		dateStr, ok := v["date"].(string)
		if !ok || dateStr == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "date: incorrect type or empty",
			})
			isJsonError = true
		}

		date, err := time.Parse("2006-01-02T15:04:05", dateStr)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "date: can't convert to date format",
			})
			isJsonError = true
		}

		number, ok := v["number"].(string)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "number: incorrect type",
			})
			isJsonError = true
		}

		harvestTypeId, ok := v["harvest_type_id"].(string)
		if !ok || harvestTypeId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "HarvestType",
				Action:   "checking value",
				Message:  "harvest_type_id: incorrect type or empty",
			})
			isJsonError = true
		}

		vehicleId, ok := v["vehicle_id"].(string)
		if !ok || vehicleId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Vehicle",
				Action:   "checking value",
				Message:  "vehicle_id: incorrect type or empty",
			})
			isJsonError = true
		}

		departureDateStr, ok := v["departure_date"].(string)
		if !ok || departureDateStr == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "departure_date",
				Action:   "checking value",
				Message:  "departure_date: incorrect type or empty",
			})
			isJsonError = true
		}

		departureDate, err := time.Parse("2006-01-02T15:04:05", departureDateStr)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "departure_date",
				Action:   "checking value",
				Message:  "departure_date: can't convert to date format",
			})
			isJsonError = true
		}

		driverId, ok := v["driver_id"].(string)
		if !ok || driverId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Driver",
				Action:   "checking value",
				Message:  "driver_id: incorrect type or empty",
			})
			isJsonError = true
		}

		recipientId, ok := v["recipient_id"].(string)
		if !ok || recipientId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "recipient_id: incorrect type or empty",
			})
			isJsonError = true
		}

		senderId, ok := v["sender_id"].(string)
		if !ok || senderId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "sender_id: incorrect type or empty",
			})
			isJsonError = true
		}

		materiallyResponsiblePersonId, ok := v["materially_responsible_person_id"].(string)
		//TODO: change back
		// if !ok || materiallyResponsiblePersonId == "" {
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "materially_responsible_person_id: incorrect type or empty",
			})
			isJsonError = true
		}

		responsiblePersonId, ok := v["app_user_id"].(string)
		if !ok || responsiblePersonId == "" {		
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "responsible_person_id: incorrect type or empty",
			})
			isJsonError = true
		}

		gross, ok := v["gross"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "gross: incorrect type",
			})
			isJsonError = true
		}

		tare, ok := v["tare"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "tare: incorrect type",
			})
			isJsonError = true
		}

		net, ok := v["net"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "net: incorrect type",
			})
			isJsonError = true
		}

		humidity, ok := v["humidity"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "humidity: incorrect type",
			})
			isJsonError = true
		}

		weediness, ok := v["weediness"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "weediness: incorrect type",
			})
			isJsonError = true
		}

		isDeleted, ok := v["is_deleted"].(bool)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "is_deleted: incorrect type",
			})
			isJsonError = true
		}

		createdAtStr, ok := v["created_at"].(string)
		if !ok || createdAtStr == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "created_at: incorrect type or empty",
			})
			isJsonError = true
		}

		createdAt, err := time.Parse("2006-01-02T15:04:05", createdAtStr)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "created_at: cant convert to date format",
			})
			isJsonError = true
		}

		updatedAtStr, ok := v["updated_at"].(string)
		if !ok || updatedAtStr == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "updated_at: incorrect type or empty",
			})
			isJsonError = true
		}

		updatedAt, err := time.Parse("2006-01-02T15:04:05", updatedAtStr)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "updated_at: cant convert to date format",
			})
			isJsonError = true
		}

		if isJsonError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		isDataError := false

		// consignmentNoteIn, sm := dao.GetConsignmentNoteInByExtId(apiV1.obx, extId)
		// if consignmentNoteIn == nil {
		// 	sm.SourceId = extId
		// 	sm.DataId = extId
		// 	sa.Messages = append(sa.Messages, sm)
		// 	isDataError = true
		// }

		harvestType, sm := dao.GetHarvestTypeByExtId(apiV1.obx, harvestTypeId)
		if harvestType == nil {
			sm.DataType = "HarvestType"
			sm.DataId = harvestTypeId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		vehicle, sm := dao.GetVehicleByExtId(apiV1.obx, vehicleId)
		if vehicle == nil {
			sm.DataType = "Vehicle"
			sm.DataId = vehicleId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		driver, sm := dao.GetPersonByExtId(apiV1.obx, driverId)
		if driver == nil {
			sm.DataType = "Driver"
			sm.DataId = driverId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		recipient, sm := dao.GetStorageByExtId(apiV1.obx, recipientId)
		if recipient == nil {
			sm.DataType = "Storage"
			sm.DataId = recipientId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		sender, sm := dao.GetStorageByExtId(apiV1.obx, senderId)
		if sender == nil {
			sm.DataType = "Storage"
			sm.DataId = senderId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		materiallyResponsiblePerson, sm := dao.GetAppUserByExtId(apiV1.obx, materiallyResponsiblePersonId)
		// TODO:
		// if materiallyResponsiblePerson == nil {
		// sm.DataType = "AppUser"
		// sm.DataId = materiallyResponsiblePersonId
		// sm.Action = "db select by ext id"
		// sm.Message = "not found"
		// pd.Messages = append(pd.Messages, sm)

		// isDataError = true
		// }

		responsiblePerson, sm := dao.GetAppUserByExtId(apiV1.obx, responsiblePersonId)
		// TODO
		// if responsiblePerson == nil {
		// sm.DataType = "AppUser"
		// sm.DataId = responsiblePersonId
		// sm.Action = "db select by ext id"
		// sm.Message = "not found"
		// pd.Messages = append(pd.Messages, sm)

		// isDataError = true
		// }

		if isDataError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		if srvId > 0 {
			cni, err := box.Get(srvId)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "query by id",
					Message:  err.Error(),
				})
			}
			cnis = append(cnis, cni)
		} else if extId != "" {
			query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(extId, true))
			cnis, err = query.Find()
			query.Close()
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "query by ext id",
					Message:  err.Error(),
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				query.Close()
				continue
			}
		} else if appId != "" {
			query := box.Query(models.ConsignmentNoteIn_.AppId.Equals(appId, true))
			cnis, err = query.Find()
			query.Close()
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "query by app id",
					Message:  err.Error(),
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				query.Close()
				continue
			}
		}

		// query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(extId, true))
		// ConsignmentNoteIns, err := query.Find()
		// query.Close()
		// if err != nil {
		// 	sa.Messages = append(sa.Messages, models.ServerMessage{
		// 		SourceId: extId,
		// 		Action:   "query",
		// 		Message:  err.Error(),
		// 	})
		// 	query.Close()
		// 	continue
		// }

		ConsignmentNoteIn := models.ConsignmentNoteIn{
			ExtId:                       extId,
			AppId:                       appId,
			Date:                        date,
			Number:                      number,
			HarvestType:                 harvestType,
			Vehicle:                     vehicle,
			DepartureDate:               departureDate,
			Driver:                      driver,
			Recipient:                   recipient,
			Sender:                      sender,
			MateriallyResponsiblePerson: materiallyResponsiblePerson,
			ResponsiblePerson:           responsiblePerson,
			Gross:                       gross,
			Tare:                        tare,
			Net:                         net,
			Humidity:                    humidity,
			Weediness:                   weediness,
			IsDeleted:                   isDeleted,
			CreatedAt:                   createdAt,
			UpdatedAt:                   updatedAt,
		}

		if len(cnis) == 0 {

			if ConsignmentNoteIn.CreatedAt.IsZero() {
				ConsignmentNoteIn.CreatedAt = time.Now().UTC()
				ConsignmentNoteIn.UpdatedAt = time.Now().UTC()
			}

			srvId, err := box.Put(&ConsignmentNoteIn)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "insert",
					Message:  err.Error(),
				})
			}

			pd.Status = http.StatusOK
			pd.SrvId = srvId

		} else if len(cnis) == 1 {
			ConsignmentNoteIn.Id = cnis[0].Id
			ConsignmentNoteIn.CreatedAt = cnis[0].CreatedAt
			ConsignmentNoteIn.UpdatedAt = time.Now().UTC()

			err := box.Update(&ConsignmentNoteIn)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "update",
					Message:  err.Error(),
				})
			}

			pd.Status = http.StatusOK
			pd.SrvId = ConsignmentNoteIn.Id

		} else {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(srvId),
				Action:   "select",
				Message:  "more than 1",
			})
		}

		if len(pd.Messages) > 0 {
			pd.Status = http.StatusBadRequest
		}
		sa.ProcessedData = append(sa.ProcessedData, pd)

	}

	sa.Send(w)
}

func (api *ApiV1) ConsignmentNoteInGet(w http.ResponseWriter, r *http.Request) {

	var gs []*models.ConsignmentNoteIn
	var err error

	fvSrvId := r.FormValue("id")
	fvAppId := r.FormValue("app_id")
	fvExtId := r.FormValue("ext_id")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	// if fvSrvId == "" && fvAppId == "" && fvExtId == "" {
	// 	sa.Status = http.StatusBadRequest
	// 	sa.Error = "id isn't specified"
	// 	sa.Send(w)
	// 	return
	// }

	box := models.BoxForConsignmentNoteIn(api.obx)

	var srvId uint64
	if fvSrvId != "" {
		srvIdInt, err := strconv.Atoi(fvSrvId)
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}
		srvId = uint64(srvIdInt)
	} else {
		srvId = 0
	}

	// if srvId == 0 && fvAppId == "" && fvExtId == "" {
	// 	gs, err = box.GetAll()
	// 	if err != nil {
	// 		sa.Status = http.StatusInternalServerError
	// 		sa.Error = err.Error()
	// 		sa.Send(w)
	// 		return
	// 	}

	// } else if srvId != 0 {
	if srvId != 0 {
		srvIdInt, err := strconv.Atoi(fvSrvId)
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()

			pd := models.ServerProcessedData{
				SrvId:  srvId,
				AppId:  fvAppId,
				ExtId:  fvExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query by id",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)
			return
		}

		srvId = uint64(srvIdInt)
		cni, err := box.Get(srvId)
		gs = append(gs, cni)

		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()

			pd := models.ServerProcessedData{
				SrvId:  srvId,
				AppId:  fvAppId,
				ExtId:  fvExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query by id",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)
			return
		}
	} else if fvAppId != "" {
		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(fvAppId, true))
		gs, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()

			pd := models.ServerProcessedData{
				SrvId:  srvId,
				AppId:  fvAppId,
				ExtId:  fvExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query by app id",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)

			query.Close()
			return
		}
	} else if fvExtId != "" {
		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(fvExtId, true))
		gs, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()

			pd := models.ServerProcessedData{
				SrvId:  srvId,
				AppId:  fvAppId,
				ExtId:  fvExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query by ext id",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)

			query.Close()
			return
		}
	} else if fvExtId == "" {
		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals("", true))
		gs, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()

			pd := models.ServerProcessedData{
				SrvId:  srvId,
				AppId:  fvAppId,
				ExtId:  fvExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query all new for ext db",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)

			query.Close()
			return
		}
	}

	bs, err := json.Marshal(gs)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
