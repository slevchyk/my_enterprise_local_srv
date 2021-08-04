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

func (api *ApiV1) ConsignmentNoteInPost(w http.ResponseWriter, r *http.Request) {

	var err error

	// var sas []models.ServerAnswer
	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "post",
		Status:    http.StatusOK,
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

	box := models.BoxForConsignmentNoteIn(api.obx)

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

		appUserId, ok := v["app_user_id"].(string)
		if !ok || appUserId == "" {
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

		if isJsonError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		isDataError := false

		harvestType, sm := dao.GetHarvestTypeByExtId(api.obx, harvestTypeId)
		if harvestType == nil {
			sm.DataType = "HarvestType"
			sm.DataId = harvestTypeId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		vehicle, sm := dao.GetVehicleByExtId(api.obx, vehicleId)
		if vehicle == nil {
			sm.DataType = "Vehicle"
			sm.DataId = vehicleId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		driver, sm := dao.GetPersonByExtId(api.obx, driverId)
		if driver == nil {
			sm.DataType = "Driver"
			sm.DataId = driverId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		recipient, sm := dao.GetStorageByExtId(api.obx, recipientId)
		if recipient == nil {
			sm.DataType = "Storage"
			sm.DataId = recipientId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		sender, sm := dao.GetStorageByExtId(api.obx, senderId)
		if sender == nil {
			sm.DataType = "Storage"
			sm.DataId = senderId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		appUser, sm := dao.GetAppUserByExtId(api.obx, appUserId)
		if appUser == nil {
			sm.DataType = "AppUser"
			sm.DataId = appUserId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		if isDataError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

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

		ConsignmentNoteIn := models.ConsignmentNoteIn{
			ExtId:         extId,
			Date:          date,
			Number:        number,
			HarvestType:   harvestType,
			Vehicle:       vehicle,
			DepartureDate: departureDate,
			Driver:        driver,
			Recipient:     recipient,
			Sender:        sender,
			AppUser:       appUser,
			Gross:         gross,
			Tare:          tare,
			Net:           net,
			Humidity:      humidity,
			Weediness:     weediness,
			IsDeleted:     isDeleted,
			ChangedByAcc:  true,
		}

		ConsignmentNoteIn.Id = cni.Id
		ConsignmentNoteIn.AppId = cni.AppId
		ConsignmentNoteIn.CreatedAt = cni.CreatedAt
		ConsignmentNoteIn.UpdatedAt = time.Now().UTC()

		err = box.Update(&ConsignmentNoteIn)
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(srvId),
				Action:   "update",
				Message:  err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)

			continue
		}

		pd.Status = http.StatusOK
		pd.SrvId = ConsignmentNoteIn.Id

		sa.ProcessedData = append(sa.ProcessedData, pd)

	}

	sa.Send(w)
}

func (api *ApiV1) ConsignmentNoteInGet(w http.ResponseWriter, r *http.Request) {

	var gs []*models.ConsignmentNoteIn
	var err error

	fvAll := r.FormValue("all")
	fvSrvId := r.FormValue("id")
	fvAppId := r.FormValue("app_id")
	fvExtId := r.FormValue("ext_id")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

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

	if fvAll == "true" {
		gs, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}
	} else if srvId != 0 {
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
	} else {
		query := box.Query(models.ConsignmentNoteIn_.ChangedByApp.Equals(true))
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

func (api *ApiV1) ConsignmentNoteInAppPost(w http.ResponseWriter, r *http.Request) {

	var err error

	fvToken := r.FormValue("token")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	if fvToken == "" {
		sa.Status = http.StatusUnauthorized
		sa.Error = "no token"
		sa.Send(w)
		return
	}

	au, err := models.GtAppUserByToken(api.obx, fvToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

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

	box := models.BoxForConsignmentNoteIn(api.obx)

	for _, v := range jsonData {

		var pd models.ServerProcessedData
		var cnis []*models.ConsignmentNoteIn

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

		// materiallyResponsiblePersonId, ok := v["materially_responsible_person_id"].(string)
		// if !ok || materiallyResponsiblePersonId == "" {
		// 	pd.Messages = append(pd.Messages, models.ServerMessage{
		// 		Action:  "checking value",
		// 		Message: "materially_responsible_person_id: incorrect type or empty",
		// 	})
		// 	isJsonError = true
		// }

		appUserId, ok := v["app_user_id"].(string)
		if !ok || appUserId == "" {
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

		if appUserId != au.ExtId {
			pd.Status = http.StatusUnauthorized
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking token",
				Message: "app user != token",
			})
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		isDataError := false

		harvestType, sm := dao.GetHarvestTypeByExtId(api.obx, harvestTypeId)
		if harvestType == nil {
			sm.DataType = "HarvestType"
			sm.DataId = harvestTypeId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		vehicle, sm := dao.GetVehicleByExtId(api.obx, vehicleId)
		if vehicle == nil {
			sm.DataType = "Vehicle"
			sm.DataId = vehicleId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		driver, sm := dao.GetPersonByExtId(api.obx, driverId)
		if driver == nil {
			sm.DataType = "Driver"
			sm.DataId = driverId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		recipient, sm := dao.GetStorageByExtId(api.obx, recipientId)
		if recipient == nil {
			sm.DataType = "Storage"
			sm.DataId = recipientId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		sender, sm := dao.GetStorageByExtId(api.obx, senderId)
		if sender == nil {
			sm.DataType = "Storage"
			sm.DataId = senderId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

		// materiallyResponsiblePerson, sm := dao.GetAppUserByExtId(apiV1.obx, materiallyResponsiblePersonId)

		appUser, sm := dao.GetAppUserByExtId(api.obx, appUserId)
		if appUser == nil {
			sm.DataType = "AppUser"
			sm.DataId = appUserId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}

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

		ConsignmentNoteIn := models.ConsignmentNoteIn{
			ExtId:         extId,
			AppId:         appId,
			Date:          date,
			Number:        number,
			HarvestType:   harvestType,
			Vehicle:       vehicle,
			DepartureDate: departureDate,
			Driver:        driver,
			Recipient:     recipient,
			Sender:        sender,
			// MateriallyResponsiblePerson: materiallyResponsiblePerson,
			AppUser:      appUser,
			Gross:        gross,
			Tare:         tare,
			Net:          net,
			Humidity:     humidity,
			Weediness:    weediness,
			IsDeleted:    isDeleted,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			ChangedByApp: true,
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

			if cnis[0].ChangedByAcc {
				pd.Status = http.StatusLocked
				pd.ChangedByAcc = true
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(srvId),
					Action:   "update",
					Message:  "changed by accounting db",
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

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

func (api *ApiV1) ConsignmentNoteInAppGet(w http.ResponseWriter, r *http.Request) {

	var cs []*models.ConsignmentNoteIn
	var err error
	var query *models.ConsignmentNoteInQuery

	fvToken := r.FormValue("token")
	fvAll := r.FormValue("all")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	if fvToken == "" {
		sa.Status = http.StatusBadRequest
		sa.Error = "token isn't specified"
		sa.Send(w)
		return
	}

	au, err := models.GtAppUserByToken(api.obx, fvToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	box := models.BoxForConsignmentNoteIn(api.obx)

	if fvAll == "true" {
		query = box.Query(models.ConsignmentNoteIn_.AppUser.Equals(au.Id))
	} else {
		query = box.Query(models.ConsignmentNoteIn_.AppUser.Equals(au.Id), models.ConsignmentNoteIn_.ChangedByAcc.Equals(true))
	}

	cs, err = query.Find()
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)

		query.Close()
		return
	}

	bs, err := json.Marshal(cs)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func (api *ApiV1) ConsignmentNoteInAppProcessed(w http.ResponseWriter, r *http.Request) {

	var err error

	fvToken := r.FormValue("token")

	sa := models.ServerAnswer{
		Status:    http.StatusOK,
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

	au, err := models.GtAppUserByToken(api.obx, fvToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var jsonData []int
	err = json.Unmarshal(bs, &jsonData)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForConsignmentNoteIn(api.obx)

	for _, v := range jsonData {

		var pd models.ServerProcessedData

		srvId := uint64(v)
		pd.SrvId = srvId
		pd.Status = http.StatusOK

		c, err := box.Get(srvId)
		if err != nil {
			pd.Status = http.StatusNotFound
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataId:  fmt.Sprintln(srvId),
				Action:  "select by srv id",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		if c.AppUser.Id != au.Id {
			pd.Status = http.StatusNotFound
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataId:  fmt.Sprintln(srvId),
				Action:  "select by srv id",
				Message: "no such data",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		sa.ProcessedData = append(sa.ProcessedData, pd)
		c.ChangedByAcc = false
		box.Put(c)
	}

	sa.Send(w)
}
