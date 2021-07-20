package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) GoodsConsignmentNoteInPost(w http.ResponseWriter, r *http.Request) {

	var gcnis []models.GoodsConsignmentNoteIn
	var err error

	sa := models.ServerAnswer{Object: "GoodsConsignmentNoteIn",
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

	box := models.BoxForGoodsConsignmentNoteIn(apiV1.obx)

	for _, v := range jsonData {

		isJsonError := false

		var pd models.ServerProcessedData

		var srvId uint64
		srvIdFloat64, ok := v["srv_id"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking type",
				Message: "srv_id: incorrect type",
			})
			isJsonError = true
		} else {
			srvId = uint64(srvIdFloat64)
			pd.SrvId = srvId
		}

		appId, ok := v["app_id"].(string)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "app_id: incorrect type",
			})
			isJsonError = true
		}
		pd.AppId = appId

		extId, ok := v["ext_id"].(string)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value ext_id",
				Message: "ext_id: incorrect type",
			})
			isJsonError = true
		}
		pd.ExtId = extId

		var consignmentNoteInId uint64
		consignmentNoteInIdFloat64, ok := v["consignment_note_in_id"].(float64)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "consignment_note_in_id: incorrect type or empty",
			})
			isJsonError = true
		} else {
			consignmentNoteInId = uint64(consignmentNoteInIdFloat64)
		}

		subdivisionId, ok := v["subdivision_id"].(string)
		if !ok || subdivisionId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "subdivision_id: incorrect type or empty",
			})
			isJsonError = true
		}

		goodsGroupId, ok := v["goods_group_id"].(string)
		if !ok || goodsGroupId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "goods_group_id: incorrect type or empty",
			})
			isJsonError = true
		}

		goodsId, ok := v["goods_id"].(string)
		if !ok || goodsId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "goods_id: incorrect type or empty",
			})
			isJsonError = true
		}

		unitId, ok := v["unit_id"].(string)
		if !ok || unitId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "unit_id: incorrect type or empty",
			})
			isJsonError = true
		}

		loadingPercentage, ok := v["loading_percentage"].(float32)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "loading_percentage: incorrect type or empty",
			})
			isJsonError = true
		}

		quantity, ok := v["quantity"].(float32)
		if !ok {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "quantity: incorrect type",
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

		consignmentNoteIn, sm := dao.GetConsignmentNoteInById(apiV1.obx, consignmentNoteInId)
		if consignmentNoteIn == nil {
			pd.Messages = append(pd.Messages, sm)
			isDataError = true
		}

		subdivision, sm := dao.GetSubdivisionByExtId(apiV1.obx, subdivisionId)
		if subdivision == nil {
			pd.Messages = append(pd.Messages, sm)
			isDataError = true
		}

		goodsGroup, sm := dao.GetGoodsGroupByExtId(apiV1.obx, goodsGroupId)
		if goodsGroup != nil {
			pd.Messages = append(pd.Messages, sm)
			isDataError = true
		}

		goods, sm := dao.GetGoodsByExtId(apiV1.obx, goodsId)
		if goods != nil {
			pd.Messages = append(pd.Messages, sm)
			isDataError = true
		}

		unit, sm := dao.GetUnitByExtId(apiV1.obx, unitId)
		if unit == nil {
			pd.Messages = append(pd.Messages, sm)
			isDataError = true
		}

		if isDataError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		query := box.Query(models.GoodsConsignmentNoteIn_.ExtId.Equals(extId, true))
		GoodsConsignmentNoteIns, err := query.Find()
		query.Close()
		if err != nil {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "GoodsConsignmentNoteIn",
				Action:   "query",
				Message:  err.Error(),
			})

			pd.Status = http.StatusInternalServerError
			sa.ProcessedData = append(sa.ProcessedData, pd)

			query.Close()
			continue
		}

		GoodsConsignmentNoteIn := models.GoodsConsignmentNoteIn{
			ExtId:             extId,
			AppId:             appId,
			ConsignmentNoteIn: consignmentNoteIn,
			Subdivision:       subdivision,
			GoodsGroup:        goodsGroup,
			Goods:             goods,
			Unit:              unit,
			LoadingPercentage: loadingPercentage,
			Quantity:          quantity,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		}

		if len(GoodsConsignmentNoteIns) == 0 {
			if GoodsConsignmentNoteIn.CreatedAt.IsZero() {
				GoodsConsignmentNoteIn.CreatedAt = time.Now().UTC()
				GoodsConsignmentNoteIn.UpdatedAt = time.Now().UTC()
			}

			_, err := box.Put(&GoodsConsignmentNoteIn)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					Action:   "insert",
					Message:  err.Error(),
				})
			}

		} else if len(GoodsConsignmentNoteIns) == 1 {
			GoodsConsignmentNoteIn.Id = GoodsConsignmentNoteIns[0].Id
			GoodsConsignmentNoteIn.CreatedAt = GoodsConsignmentNoteIns[0].CreatedAt
			GoodsConsignmentNoteIn.UpdatedAt = time.Now().UTC()

			err := box.Update(&GoodsConsignmentNoteIn)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "GoodsConsignmentNoteIn",
					Action:   "update",
					Message:  err.Error(),
				})
			}
		} else {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "GoodsConsignmentNoteIn",
				Action:   "more than 1",
				Message:  err.Error(),
			})
		}

		sa.ProcessedData = append(sa.ProcessedData, pd)
	}

	err = json.Unmarshal(bs, &gcnis)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	sa.Send(w)
}

func (api *ApiV1) GoodsConsignmentNoteInGet(w http.ResponseWriter, r *http.Request) {

	var gs []*models.GoodsConsignmentNoteIn
	var err error

	fvConsignemntNoteInId := r.FormValue("consignment_note_in_id")
	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "GoodsConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	if fvConsignemntNoteInId == "" {
		sa.Status = http.StatusBadRequest
		sa.Object = "goods consignament note in"
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

	box := models.BoxForGoodsConsignmentNoteIn(api.obx)

	if fvId == "" {
		gs, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.GoodsConsignmentNoteIn_.ExtId.Equals(fvId, true))
		gs, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
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
