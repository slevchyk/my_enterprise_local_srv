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

	sa := models.ServerAnswer{SourceType: "GoodsConsignmentNoteIn",
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

		extId, ok := v["ext_id"].(string)
		if !ok || extId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value ext_id",
				Message:  "ext_id: incorrect type or empty",
			})
			isJsonError = true
		}

		appId, ok := v["app_id"].(int64)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "app_id: incorrect type",
			})
			isJsonError = true
		}

		consignmentNoteInId, ok := v["consignment_note_in_id"].(string)
		if !ok || consignmentNoteInId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "consignment_note_in_id: incorrect type or empty",
			})
			isJsonError = true
		}

		subdivisionId, ok := v["subdivision_id"].(string)
		if !ok || subdivisionId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "subdivision_id: incorrect type or empty",
			})
			isJsonError = true
		}

		goodsGroupId, ok := v["goods_group_id"].(string)
		if !ok || goodsGroupId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "goods_group_id: incorrect type or empty",
			})
			isJsonError = true
		}

		goodsId, ok := v["goods_id"].(string)
		if !ok || goodsId == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "goods_id: incorrect type or empty",
			})
			isJsonError = true
		}

		unit_id, ok := v["unit_id"].(string)
		if !ok || unit_id == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "unit_id: incorrect type or empty",
			})
			isJsonError = true
		}

		loadingPercentage, ok := v["loading_percentage"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "loading_percentage: incorrect type or empty",
			})
			isJsonError = true
		}

		quantity, ok := v["quantity"].(float32)
		if !ok {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "quantity: incorrect type",
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

		updatedAtStr, ok := v["updated_at"].(string)
		if !ok || updatedAtStr == "" {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: extId,
				Action:   "checking value",
				Message:  "updated_at: incorrect type or empty",
			})
			isJsonError = true
		}

		if isJsonError {
			continue
		}

		isDataError := false

		unit, sm := dao.GetUnitByExtId(apiV1.obx, unit_id)
		if unit == nil {
			sm.SourceId = extId
			sm.DataId = unit_id
			sa.Messages = append(sa.Messages, sm)
			isDataError = true
		}

		if isDataError {
			continue
		}

		query := box.Query(models.GoodsConsignmentNoteIn_.ExtId.Equals(extId, true))
		GoodsConsignmentNoteIns, err := query.Find()
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

		GoodsConsignmentNoteIn := models.GoodsConsignmentNoteIn{
			ExtId:             extId,
			AppId:             appId,
			Unit:              unit,
			LoadingPercentage: loadingPercentage,
			Quantity:          quantity,
		}

		if len(GoodsConsignmentNoteIns) == 0 {
			GoodsConsignmentNoteIn.CreatedAt = time.Now().UTC()
			GoodsConsignmentNoteIn.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&GoodsConsignmentNoteIn)
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					Status:   http.StatusInternalServerError,
					SourceId: extId,
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

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{SourceType: "GoodsConsignmentNoteIn",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

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
