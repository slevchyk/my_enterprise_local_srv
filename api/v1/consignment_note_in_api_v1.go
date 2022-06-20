package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (api *ApiV1) ConsignmentNoteInPost(w http.ResponseWriter, r *http.Request) {

	var err error
	var cniis []models.ConsignmentNoteInImport

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "post",
		DateUTC:   time.Now()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &cniis)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	for _, v := range cniis {

		pd := postConsignmentNoteIn(api.obx, v, true)

		if len(pd.Messages) > 0 {
			pd.Status = http.StatusBadRequest
		}
		sa.ProcessedData = append(sa.ProcessedData, pd)
	}

	sa.Send(w)
}

func (api *ApiV1) ConsignmentNoteInGet(w http.ResponseWriter, r *http.Request) {

	var cnies []*models.ConsignmentNoteInExport
	var cnis []*models.ConsignmentNoteIn
	var err error

	fvAll := r.FormValue("all")
	fvSrvId := r.FormValue("id")
	fvAppId := r.FormValue("app_id")
	fvExtId := r.FormValue("ext_id")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now()}

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
		cnis, err = box.GetAll()
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
		cnis = append(cnis, cni)

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
		cnis, err = query.Find()
		query.Close()
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
		cnis, err = query.Find()
		query.Close()
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
		cnis, err = query.Find()
		query.Close()
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

	boxGoodsConsignmentNoteIn := models.BoxForGoodsConsignmentNoteIn(api.obx)
	for _, cni := range cnis {

		query := boxGoodsConsignmentNoteIn.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id))
		gcnis, err := query.Find()
		query.Close()
		if err != nil {
			pd := models.ServerProcessedData{
				SrvId:  cni.Id,
				AppId:  cni.AppId,
				ExtId:  cni.ExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query goods rows",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)

			query.Close()
			return
		}

		cnies = append(cnies, &models.ConsignmentNoteInExport{
			Document:   cni,
			TableGoods: gcnis,
		})

	}

	bs, err := json.Marshal(cnies)
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
	var cniis []models.ConsignmentNoteInImport

	fvToken := r.FormValue("token")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "post",
		DateUTC:   time.Now()}

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

	err = json.Unmarshal(bs, &cniis)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	for _, v := range cniis {

		var pd models.ServerProcessedData

		pd.SrvId = v.Id
		pd.AppId = v.AppId
		pd.ExtId = v.ExtId

		//перевіримо чи дані які прийшли дійсно від авторизованого користувача
		if v.AppUserId != au.ExtId {
			pd.Status = http.StatusUnauthorized
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking token",
				Message: "app user != token",
			})
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		pd = postConsignmentNoteIn(api.obx, v, false)

		if len(pd.Messages) > 0 {
			pd.Status = http.StatusBadRequest
		}
		sa.ProcessedData = append(sa.ProcessedData, pd)

	}

	sa.Send(w)
}

func (api *ApiV1) ConsignmentNoteInAppGet(w http.ResponseWriter, r *http.Request) {

	var cnis []*models.ConsignmentNoteIn
	var err error
	var query *models.ConsignmentNoteInQuery
	var cnies []*models.ConsignmentNoteInExport

	fvToken := r.FormValue("token")
	fvAll := r.FormValue("all")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now()}

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

	cnis, err = query.Find()
	query.Close()
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)

		query.Close()
		return
	}

	boxGoodsConsignmentNoteIn := models.BoxForGoodsConsignmentNoteIn(api.obx)
	for _, cni := range cnis {

		query := boxGoodsConsignmentNoteIn.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id))
		gcnis, err := query.Find()
		query.Close()
		if err != nil {
			pd := models.ServerProcessedData{
				SrvId:  cni.Id,
				AppId:  cni.AppId,
				ExtId:  cni.ExtId,
				Status: http.StatusInternalServerError,
			}

			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "query goods rows",
				Message: err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			sa.Send(w)

			query.Close()
			return
		}

		cnies = append(cnies, &models.ConsignmentNoteInExport{
			Document:   cni,
			TableGoods: gcnis,
		})
	}

	bs, err := json.Marshal(cnies)
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
		DateUTC:   time.Now()}

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

func (api *ApiV1) ConsignmentNoteInAppChanged(w http.ResponseWriter, r *http.Request) {

	var err error

	fvToken := r.FormValue("token")

	sa := models.ServerAnswer{
		Object:    "ConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now()}

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

	query := box.Query(models.ConsignmentNoteIn_.AppUser.Equals(au.Id), models.ConsignmentNoteIn_.ChangedByAcc.Equals(true))
	ids, err := query.FindIds()
	if err != nil {
		sa.Status = http.StatusBadRequest
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	bs, err := json.Marshal(ids)
	if err != nil {
		sa.Status = http.StatusBadRequest
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func (api *ApiV1) ConsignmentNoteInDelete(w http.ResponseWriter, r *http.Request) {

	box := models.BoxForGoodsConsignmentNoteIn(api.obx)
	box.RemoveAll()

}

func parseConsignmentNoteIn(obx *objectbox.ObjectBox, cnii models.ConsignmentNoteInImport, isAcc bool) (models.ConsignmentNoteIn, []models.GoodsConsignmentNoteIn, models.ServerProcessedData) {

	var pd models.ServerProcessedData
	var sm models.ServerMessage
	var gcnis []models.GoodsConsignmentNoteIn

	pd.SrvId = cnii.Id
	pd.AppId = cnii.AppId
	pd.ExtId = cnii.ExtId

	isDataError := false

	date, err := time.Parse("2006-01-02T15:04:05", cnii.Date)
	if err != nil {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			Action:  "checking value",
			Message: "date: can't convert to date format",
		})

		isDataError = true
	}

	var harvestType *models.HarvestType
	if cnii.HarvestTypeId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "HarvestType",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		harvestType, sm = dao.GetHarvestTypeByExtId(obx, cnii.HarvestTypeId)
		if harvestType == nil {
			sm.DataType = "HarvestType"
			sm.DataId = cnii.HarvestTypeId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var vehicle *models.Vehicle
	if cnii.VehicleId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "Vehicle",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		vehicle, sm = dao.GetVehicleByExtId(obx, cnii.VehicleId)
		if vehicle == nil {
			sm.DataType = "Vehicle"
			sm.DataId = cnii.VehicleId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var trailer *models.Trailer
	if cnii.TrailerId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "Trailer",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		trailer, sm = dao.GetTrailerByExtId(obx, cnii.VehicleId)
		if trailer == nil {
			sm.DataType = "Trailer"
			sm.DataId = cnii.TrailerId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	departureDate, err := time.Parse("2006-01-02T15:04:05", cnii.DepartureDateId)
	if err != nil {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			Action:  "checking value",
			Message: "departureDate: can't convert to date format",
		})

		isDataError = true
	}

	var driver *models.ServiceWorker
	if cnii.DriverId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "Person",
			Action:   "checking data Driver",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		driver, sm = dao.GetServiceWorkerByExtId(obx, cnii.DriverId)
		if driver == nil {
			sm.DataType = "Driver"
			sm.DataId = cnii.DriverId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var recipient *models.Storage
	if cnii.RecipientId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "Storage",
			Action:   "checking data Recipient",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		recipient, sm = dao.GetStorageByExtId(obx, cnii.RecipientId)
		if recipient == nil {
			sm.DataType = "Storage"
			sm.DataId = cnii.RecipientId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var sender *models.Storage
	var manager *models.AppUser
	if cnii.OperationId == 0 {
		if cnii.ManagerId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "AppUser",
				Action:   "checking data Manager",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			manager, sm = dao.GetAppUserByExtId(obx, cnii.ManagerId)
			if manager == nil {
				sm.DataType = "AppUser"
				sm.DataId = cnii.ManagerId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}
	} else {
		if cnii.SenderId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Storage",
				Action:   "checking data Sender",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			sender, sm = dao.GetStorageByExtId(obx, cnii.SenderId)
			if sender == nil {
				sm.DataType = "Storage"
				sm.DataId = cnii.SenderId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}
	}

	var appUser *models.AppUser
	if cnii.AppUserId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "AppUser",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		appUser, sm = dao.GetAppUserByExtId(obx, cnii.AppUserId)
		if appUser == nil {
			sm.DataType = "AppUser"
			sm.DataId = cnii.AppUserId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var createdAt time.Time
	if isAcc {
		createdAt = time.Now()
	} else {
		createdAt, err = time.Parse("2006-01-02T15:04:05", cnii.CreatedAt)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "createdAt: can't convert to date format",
			})

			isDataError = true
		}
	}

	var updatedAt time.Time
	if isAcc {
		updatedAt = time.Now()
	} else {
		updatedAt, err = time.Parse("2006-01-02T15:04:05", cnii.UpdatedAt)
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "updatedAt: can't convert to date format",
			})

			isDataError = true
		}
	}

	for _, gcnii := range cnii.Goods {

		if isAcc {
			if gcnii.ExtId == "" {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					RowId:   gcnii.ExtId,
					Action:  "checking value",
					Message: "row ext id isn't specified",
				})

				isDataError = true
			}
		} else {
			if gcnii.AppId == "" {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					RowId:   gcnii.AppId,
					Action:  "checking value",
					Message: "row app id isn't specified",
				})

				isDataError = true
			}
		}

		var rowSubdivision *models.Subdivision
		if gcnii.SubdivisionId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Subdivision",
				Action:   "checking data Recipient",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			rowSubdivision, sm = dao.GetSubdivisionByExtId(obx, gcnii.SubdivisionId)
			if recipient == nil {
				sm.DataType = "Subdivision"
				sm.DataId = gcnii.SubdivisionId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}

		var rowGoodsGroup *models.GoodsGroup
		if gcnii.GoodsGroupId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "GoodsGroup",
				Action:   "checking data Recipient",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			rowGoodsGroup, sm = dao.GetGoodsGroupByExtId(obx, gcnii.GoodsGroupId)
			if recipient == nil {
				sm.DataType = "GoodsGroup"
				sm.DataId = gcnii.GoodsGroupId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}

		var rowGoods *models.Goods
		if gcnii.GoodsId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Goods",
				Action:   "checking data Recipient",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			rowGoods, sm = dao.GetGoodsByExtId(obx, gcnii.GoodsId)
			if recipient == nil {
				sm.DataType = "Goods"
				sm.DataId = gcnii.GoodsId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}

		var rowUnit *models.Unit
		if gcnii.UnitId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "Unit",
				Action:   "checking data Recipient",
				Message:  "ext id isn't specified",
			})

			isDataError = true
		} else {
			rowUnit, sm = dao.GetUnitByExtId(obx, gcnii.UnitId)
			if recipient == nil {
				sm.DataType = "Unit"
				sm.DataId = gcnii.UnitId
				sm.Action = "db select by ext id"
				sm.Message = "not found"
				pd.Messages = append(pd.Messages, sm)

				isDataError = true
			}
		}

		var rowCreatedAt time.Time
		if isAcc {
			rowCreatedAt = time.Now()
		} else {
			rowCreatedAt, err = time.Parse("2006-01-02T15:04:05", gcnii.CreatedAt)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "checking value",
					Message: "row createdAt: can't convert to date format",
				})

				isDataError = true
			}
		}

		var rowUpdatedAt time.Time
		if isAcc {
			rowUpdatedAt = time.Now()
		} else {
			rowUpdatedAt, err = time.Parse("2006-01-02T15:04:05", gcnii.UpdatedAt)
			if err != nil {
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "checking value",
					Message: "updatedAt: can't convert to date format",
				})

				isDataError = true
			}
		}

		gcnis = append(gcnis, models.GoodsConsignmentNoteIn{
			Id:                gcnii.Id,
			AppId:             gcnii.AppId,
			ExtId:             gcnii.ExtId,
			Subdivision:       rowSubdivision,
			GoodsGroup:        rowGoodsGroup,
			Goods:             rowGoods,
			Unit:              rowUnit,
			LoadingPercentage: gcnii.LoadingPercentage,
			Quantity:          gcnii.Quantity,
			CreatedAt:         rowCreatedAt,
			UpdatedAt:         rowUpdatedAt,
		})

	}

	cni := models.ConsignmentNoteIn{
		ExtId:         cnii.ExtId,
		AppId:         cnii.AppId,
		Date:          date,
		Number:        cnii.Number,
		OperationId:   cnii.OperationId,
		StatusId:      cnii.StatusId,
		ExtNumber:     cnii.ExtNumber,
		HarvestType:   harvestType,
		Vehicle:       vehicle,
		Trailer:       trailer,
		DepartureDate: departureDate,
		Driver:        driver,
		Recipient:     recipient,
		Manager:       manager,
		Sender:        sender,
		AppUser:       appUser,
		Comment:       cnii.Comment,
		Gross:         cnii.Gross,
		Tare:          cnii.Tare,
		Net:           cnii.Net,
		Humidity:      cnii.Humidity,
		Weediness:     cnii.Weediness,
		IsDeleted:     cnii.IsDeleted,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}

	pd.Status = http.StatusOK
	if isDataError {
		pd.Status = http.StatusBadRequest
	}

	return cni, gcnis, pd
}

func postGoodsConsignmentNoteIn(obx *objectbox.ObjectBox, cni models.ConsignmentNoteIn, gcnis []models.GoodsConsignmentNoteIn) models.ServerProcessedData {

	var rowId string
	var existGcnis []*models.GoodsConsignmentNoteIn
	var toDeleteIds []uint64
	var err error
	var spd models.ServerProcessedData

	spd.Status = http.StatusOK
	spd.SrvId = cni.Id
	spd.AppId = cni.AppId
	spd.ExtId = cni.ExtId

	box := models.BoxForGoodsConsignmentNoteIn(obx)

	// виберемо всі рядки документа із БД
	query := box.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id))
	dbIds, err := query.FindIds()
	if err != nil {
		spd.Status = http.StatusInternalServerError
		spd.Messages = append(spd.Messages, models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			DataId:   fmt.Sprint(cni.Id),
			Action:   "query by cni id",
			Message:  err.Error(),
		})

		return spd
	}

	// знайдемо слайс всіх ід які є в базі але вже немає в докумені
	for _, v := range dbIds {
		found := false
		for _, gcni := range gcnis {
			if v == gcni.Id {
				found = true
				break
			}
		}

		if !found {
			toDeleteIds = append(toDeleteIds, v)
		}
	}

	// видалимо з бд рядки яких вже немає в документі
	if len(toDeleteIds) > 0 {
		_, err = box.RemoveIds(toDeleteIds...)
		if err != nil {
			spd.Status = http.StatusInternalServerError
			spd.Messages = append(spd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cni.Id),
				Action:   "delete rows",
				Message:  err.Error(),
			})

			return spd
		}
	}

	// створимо/оновимо рядки
	isError := false
	for _, gcni := range gcnis {

		if gcni.Id > 0 {
			rowId = fmt.Sprint(gcni.Id)
		} else if gcni.ExtId != "" {
			rowId = fmt.Sprint(gcni.ExtId)
		} else {
			rowId = fmt.Sprint(gcni.AppId)
		}

		var spr models.ServerProcessedRow

		spr.Status = http.StatusOK
		spr.SrvId = gcni.Id
		spr.AppId = gcni.AppId
		spr.ExtId = gcni.ExtId

		if gcni.Id > 0 {
			query := box.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id), models.GoodsConsignmentNoteIn_.Id.Equals(gcni.Id))
			existGcnis, err = query.Find()
			query.Close()
			if err != nil {
				spr.Status = http.StatusInternalServerError
				spr.Messages = append(spd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					DataId:   rowId,
					Action:   "query by id",
					Message:  err.Error(),
				})

				spd.Rows = append(spd.Rows, spr)
				isError = true
				continue
			}

		} else if gcni.ExtId != "" {
			query := box.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id), models.GoodsConsignmentNoteIn_.ExtId.Equals(gcni.ExtId, true))
			existGcnis, err = query.Find()
			query.Close()
			if err != nil {
				spr.Status = http.StatusInternalServerError
				spr.Messages = append(spd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					RowId:    rowId,
					Action:   "query by ext id",
					Message:  err.Error(),
				})

				spd.Rows = append(spd.Rows, spr)
				isError = true
				continue
			}
		} else if gcni.AppId != "" {
			query := box.Query(models.GoodsConsignmentNoteIn_.ConsignmentNoteIn.Equals(cni.Id), models.GoodsConsignmentNoteIn_.AppId.Equals(gcni.AppId, true))
			existGcnis, err = query.Find()
			query.Close()
			if err != nil {
				spr.Status = http.StatusInternalServerError
				spr.Messages = append(spd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					RowId:    rowId,
					Action:   "query by app id",
					Message:  err.Error(),
				})

				spd.Rows = append(spd.Rows, spr)
				isError = true
				continue
			}
		} else {
			spr.Status = http.StatusBadRequest
			spr.Messages = append(spd.Messages, models.ServerMessage{
				DataType: "GoodsConsignmentNoteIn",
				RowId:    rowId,
				Action:   "id isn't specified",
				Message:  err.Error(),
			})

			spd.Rows = append(spd.Rows, spr)
			isError = true
			continue
		}

		if len(existGcnis) == 0 {
			gcni.Id = 0
			gcni.ConsignmentNoteIn = &cni
			id, err := box.Put(&gcni)
			if err != nil {
				spr.Status = http.StatusInternalServerError
				spr.Messages = append(spd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					RowId:    rowId,
					Action:   "insert",
					Message:  err.Error(),
				})

				spd.Rows = append(spd.Rows, spr)
				isError = true
				continue
			}

			spr.SrvId = id
			spd.Rows = append(spd.Rows, spr)

		} else if len(existGcnis) == 1 {
			gcni.Id = existGcnis[0].Id
			gcni.ConsignmentNoteIn = &cni
			gcni.CreatedAt = existGcnis[0].CreatedAt
			gcni.UpdatedAt = time.Now()

			err = box.Update(&gcni)
			if err != nil {
				spr.Status = http.StatusInternalServerError
				spr.Messages = append(spd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					RowId:    rowId,
					Action:   "update",
					Message:  err.Error(),
				})

				spd.Rows = append(spd.Rows, spr)
				isError = true
				continue
			}

			spr.SrvId = gcni.Id
			spd.Rows = append(spd.Rows, spr)
			continue
		} else {
			spr.Status = http.StatusConflict
			spr.Messages = append(spd.Messages, models.ServerMessage{
				DataType: "GoodsConsignmentNoteIn",
				DataId:   fmt.Sprint(gcni.Id),
				RowId:    rowId,
				Action:   "select",
				Message:  "more than 1 found",
			})

			spd.Rows = append(spd.Rows, spr)
			isError = true
			continue
		}
	}

	if isError {
		spd.Status = http.StatusBadRequest
	}

	return spd
}

func postConsignmentNoteIn(obx *objectbox.ObjectBox, cnii models.ConsignmentNoteInImport, isAcc bool) models.ServerProcessedData {

	var existCni *models.ConsignmentNoteIn
	var err error
	var cnis []*models.ConsignmentNoteIn
	var pd models.ServerProcessedData

	pd.SrvId = cnii.Id
	pd.AppId = cnii.AppId
	pd.ExtId = cnii.ExtId

	box := models.BoxForConsignmentNoteIn(obx)

	cni, gcnis, pd := parseConsignmentNoteIn(obx, cnii, isAcc)
	cni.ChangedByApp = true

	if pd.Status != http.StatusOK {
		// sa.ProcessedData = append(sa.ProcessedData, pd)
		//continue
		return pd
	}

	if cnii.Id > 0 {
		existCni, err = box.Get(cnii.Id)
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cnii.Id),
				Action:   "query by id",
				Message:  err.Error(),
			})
		}
		cnis = append(cnis, existCni)
	} else if cnii.ExtId != "" {
		query := box.Query(models.ConsignmentNoteIn_.ExtId.Equals(cnii.ExtId, true))
		cnis, err = query.Find()
		query.Close()
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cnii.Id),
				Action:   "query by ext id",
				Message:  err.Error(),
			})
			// sa.ProcessedData = append(sa.ProcessedData, pd)
			query.Close()
			// continue
			return pd
		}
	} else if cnii.AppId != "" {
		query := box.Query(models.ConsignmentNoteIn_.AppId.Equals(cnii.AppId, true))
		cnis, err = query.Find()
		query.Close()
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cnii.Id),
				Action:   "query by app id",
				Message:  err.Error(),
			})
			// sa.ProcessedData = append(sa.ProcessedData, pd)
			query.Close()
			// continue
			return pd
		}
	}

	if len(cnis) == 0 {

		if cni.CreatedAt.IsZero() {
			cni.CreatedAt = time.Now()
			cni.UpdatedAt = time.Now()
		}

		srvId, err := box.Put(&cni)
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(srvId),
				Action:   "insert",
				Message:  err.Error(),
			})
		}

		gspd := postGoodsConsignmentNoteIn(obx, cni, gcnis)
		if gspd.Status != http.StatusOK {
			return pd
		}

		pd.Status = http.StatusOK
		pd.Rows = append(pd.Rows, gspd.Rows...)
		pd.Messages = append(pd.Messages, gspd.Messages...)

	} else if len(cnis) == 1 {

		if !isAcc && cnis[0].ChangedByAcc {
			pd.Status = http.StatusLocked
			pd.ChangedByAcc = true
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cnii.Id),
				Action:   "update",
				Message:  "changed by accounting db",
			})

			return pd
		}

		cni.Id = cnis[0].Id
		cni.CreatedAt = cnis[0].CreatedAt
		cni.UpdatedAt = time.Now()
		cni.ChangedByAcc = isAcc
		cni.ChangedByApp = !isAcc

		err := box.Update(&cni)
		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "ConsignmentNoteIn",
				DataId:   fmt.Sprint(cnii.Id),
				Action:   "update",
				Message:  err.Error(),
			})

			return pd
		}

		gspd := postGoodsConsignmentNoteIn(obx, cni, gcnis)
		pd.Rows = append(pd.Rows, gspd.Rows...)
		pd.Messages = append(pd.Messages, gspd.Messages...)

		if gspd.Status != http.StatusOK {
			pd.Status = gspd.Status
			return pd
		}
	} else {
		pd.Status = http.StatusInternalServerError
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "ConsignmentNoteIn",
			DataId:   fmt.Sprint(cnii.Id),
			Action:   "select",
			Message:  "more than 1",
		})
	}

	return pd
}
