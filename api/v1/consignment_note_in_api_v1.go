package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) ConsignmentNoteInPost(w http.ResponseWriter, r *http.Request) {

	var cnis []models.ConsignmentNoteIn
	var err error

	sa := models.ServerAnswer{SourceType: "ConsignmentNoteIn",
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

		isJsonError := false

		extId, ok := v["ext_id"].(string)
		if !ok || extId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value ext_id",
				Message:  "ext_id: incorrect type or empty",
			})
			isJsonError = true
		}

		appId, ok := v["app_id"].(string)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "app_id: incorrect type",
			})
			isJsonError = true
		}

		dateStr, ok := v["date"].(string)
		if !ok || dateStr == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "date: incorrect type or empty",
			})
			isJsonError = true
		}

		date, err := time.Parse("2006-01-02T15:04:05", dateStr)
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "date: can't convert to date format",
			})
			isJsonError = true
		}

		number, ok := v["number"].(string)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "number: incorrect type",
			})
			isJsonError = true
		}

		harvestTypeId, ok := v["harvest_type_id"].(string)
		if !ok || harvestTypeId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "harvest_type_id: incorrect type or empty",
			})
			isJsonError = true
		}

		vehicleId, ok := v["vehicle_id"].(string)
		if !ok || vehicleId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "vehicle_id: incorrect type or empty",
			})
			isJsonError = true
		}

		departureDateStr, ok := v["departure_date"].(string)
		if !ok || departureDateStr == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "departure_date: incorrect type or empty",
			})
			isJsonError = true
		}

		departureDate, err := time.Parse("2006-01-02T15:04:05", departureDateStr)
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "departure_date: can't convert to date format",
			})
			isJsonError = true
		}

		driverId, ok := v["driver_id"].(string)
		if !ok || driverId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "driver_id: incorrect type or empty",
			})
			isJsonError = true
		}

		recipientId, ok := v["recipient_id"].(string)
		if !ok || recipientId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "recipient_id: incorrect type or empty",
			})
			isJsonError = true
		}

		senderId, ok := v["sender_id"].(string)
		if !ok || senderId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "sender_id: incorrect type or empty",
			})
			isJsonError = true
		}

		materiallyResponsiblePersonId, ok := v["materially_responsible_person_id"].(string)
		if !ok || materiallyResponsiblePersonId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "materially_responsible_person_id: incorrect type or empty",
			})
			isJsonError = true
		}

		responsiblePersonId, ok := v["responsible_person_id"].(string)
		if !ok || responsiblePersonId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "responsible_person_id: incorrect type or empty",
			})
			isJsonError = true
		}

		gross, ok := v["gross"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "gross: incorrect type",
			})
			isJsonError = true
		}

		tare, ok := v["tare"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "tare: incorrect type",
			})
			isJsonError = true
		}

		net, ok := v["net"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "net: incorrect type",
			})
			isJsonError = true
		}

		humidity, ok := v["humidity"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "humidity: incorrect type",
			})
			isJsonError = true
		}

		weediness, ok := v["weediness"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "weediness: incorrect type",
			})
			isJsonError = true
		}

		isDeleted, ok := v["is_deleted"].(bool)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "is_deleted: incorrect type",
			})
			isJsonError = true
		}

		createdAtStr, ok := v["created_at"].(string)
		if !ok || createdAtStr == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "created_at: incorrect type or empty",
			})
			isJsonError = true
		}

		createdAt, err := time.Parse("2006-01-02T15:04:05", createdAtStr)
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "created_at: cant convert to date format",
			})
			isJsonError = true
		}

		updatedAtStr, ok := v["updated_at"].(string)
		if !ok || updatedAtStr == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "updated_at: incorrect type or empty",
			})
			isJsonError = true
		}

		updatedAt, err := time.Parse("2006-01-02T15:04:05", updatedAtStr)
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "updated_at: cant convert to date format",
			})
			isJsonError = true
		}

		if isJsonError {
			continue
		}

		isDataError := false

		consignmentNoteIn, sm := dao.GetConsignmentNoteInByExtId(apiV1.obx, extId)
		if consignmentNoteIn == nil {
			sm.SourceId = extId
			sm.DataId = extId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		harvestType, sm := dao.GetHarvestTypeByExtId(apiV1.obx, harvestTypeId)
		if harvestType == nil {
			sm.SourceId = extId
			sm.DataId = harvestTypeId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		vehicle, sm := dao.GetVehicleByExtId(apiV1.obx, vehicleId)
		if vehicle != nil {
			sm.SourceId = extId
			sm.DataId = vehicleId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		driver, sm := dao.GetPersonByExtId(apiV1.obx, driverId)
		if driver != nil {
			sm.SourceId = extId
			sm.DataId = driverId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		recipient, sm := dao.GetStorageByExtId(apiV1.obx, recipientId)
		if recipient == nil {
			sm.SourceId = extId
			sm.DataId = recipientId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		sender, sm := dao.GetStorageByExtId(apiV1.obx, senderId)
		if sender == nil {
			sm.SourceId = extId
			sm.DataId = senderId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		materiallyResponsiblePerson, sm := dao.GetAppUserByExtId(apiV1.obx, materiallyResponsiblePersonId)
		if materiallyResponsiblePerson == nil {
			sm.SourceId = extId
			sm.DataId = materiallyResponsiblePersonId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		responsiblePerson, sm := dao.GetAppUserByExtId(apiV1.obx, responsiblePersonId)
		if responsiblePerson == nil {
			sm.SourceId = extId
			sm.DataId = responsiblePersonId
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		if isDataError {
			continue
		}

		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(extId, true))
		ConsignmentNoteIns, err := query.Find()
		query.Close()
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "query",
				Message:  err.Error(),
			})
			query.Close()
			continue
		}

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

		if len(ConsignmentNoteIns) == 0 {

			query := box.Query(models.ConsignmentNoteIn_.AppId.Equals(appId, true))
			ConsignmentNoteIns, err = query.Find()
			query.Close()
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					SourceId: extId,
					Action:   "query",
					Message:  err.Error(),
				})
				query.Close()
				continue
			}

			if ConsignmentNoteIn.CreatedAt.IsZero() {
				ConsignmentNoteIn.CreatedAt = time.Now().UTC()
				ConsignmentNoteIn.UpdatedAt = time.Now().UTC()
			}

			_, err := box.Put(&ConsignmentNoteIn)
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					Status:   http.StatusInternalServerError,
					SourceId: extId,
					Action:   "insert",
					Message:  err.Error(),
				})
			}

		} else if len(ConsignmentNoteIns) == 1 {
			ConsignmentNoteIn.Id = ConsignmentNoteIns[0].Id
			ConsignmentNoteIn.CreatedAt = ConsignmentNoteIns[0].CreatedAt
			ConsignmentNoteIn.UpdatedAt = time.Now().UTC()

			err := box.Update(&ConsignmentNoteIn)
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					Status:   http.StatusInternalServerError,
					SourceId: extId,
					Action:   "update",
					Message:  err.Error(),
				})
			}
		} else {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				Status:   http.StatusBadRequest,
				SourceId: extId,
				Action:   "more than 1",
				Message:  err.Error(),
			})
		}
	}

	err = json.Unmarshal(bs, &cnis)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	sa.Send(w)
}

func (api *ApiV1) ConsignmentNoteInGet(w http.ResponseWriter, r *http.Request) {

	var gs []*models.ConsignmentNoteIn
	var err error

	fvConsignemntNoteInId := r.FormValue("consignment_note_in_id")
	fvId := r.FormValue("id")

	sa := models.ServerAnswer{SourceType: "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	if fvConsignemntNoteInId == "" {
		sa.Status = http.StatusBadRequest
		sa.SourceType = "goods consignament note in"
		sa.Error = "Consignament note in id is not specified"

	}

	// obx, err := objectbox.NewBuilder().Model(models.ObjectBoxModel()).Build()
	// if err != nil {
	// 	sa.Status = http.StatusInternalServerError
	// 	sa.Error = err.Error()
	// 	sa.Send(w)
	// 	return
	// }
	// defer obx.Close()

	box := models.BoxForConsignmentNoteIn(api.obx)

	if fvId == "" {
		gs, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(fvId, true))
		gs, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: fvId,
				Action:   "query",
				Message:  err.Error(),
			})
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
