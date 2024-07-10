package api

import (
	"encoding/json"
	"io"	
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) GoodsPost(w http.ResponseWriter, r *http.Request) {

	var gs []models.Goods
	var err error

	sa := models.ServerAnswer{
		Object:    "Goods",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := io.ReadAll(r.Body)
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

	box := models.BoxForGoods(apiV1.obx)
	boxUnit := models.BoxForUnit(apiV1.obx)

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

		unit_ext_id, ok := v["unit_ext_id"].(string)
		if !ok || unit_ext_id == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "unit_ext_id: incorrect type or empty",
			})
			jsonError = true
		}

		if jsonError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		queryUnit := boxUnit.Query(models.Unit_.ExtId.Equals(unit_ext_id, true))
		units, err := queryUnit.Find()
		queryUnit.Close()

		if err != nil {
			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "unit",
				DataId:   unit_ext_id,
				Action:   "query",
				Message:  err.Error(),
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		if len(units) == 0 {
			pd.Status = http.StatusNotFound
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "unit",
				DataId:   unit_ext_id,
				Action:   "query",
				Message:  "not found",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		} else if len(units) != 1 {
			pd.Status = http.StatusConflict
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "unit",
				DataId:   unit_ext_id,
				Action:   "query",
				Message:  "more than 1",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		unit := units[0]

		query := box.Query(models.Goods_.ExtId.Equals(extId, true))
		goodss, err := query.Find()
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

		goods := models.Goods{
			ExtId: extId,
			Name:  name,
			Unit:  unit,
			IsOiliness: v["is_oiliness"].(bool),
			IsOilinessDry: v["is_oiliness_dry"].(bool),
			IsErucicAcid: v["is_erucic_acid"].(bool),			
			IsGlucosinolates: v["is_glucosinolates"].(bool),
			IsMycotoxins: v["is_mycotoxins"].(bool),
			IsProtein: v["is_protein"].(bool),
			IsProteinDry: v["is_protein_dry"].(bool),
			IsAcid: v["is_acid"].(bool),			
			IsDeleted: v["is_deleted"].(bool),
		}

		if len(goodss) == 0 {
			goods.CreatedAt = time.Now().UTC()
			goods.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&goods)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "insert",
					Message: err.Error(),
				})
				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

		} else if len(goodss) == 1 {
			goods.Id = goodss[0].Id
			goods.CreatedAt = goodss[0].CreatedAt
			goods.UpdatedAt = time.Now().UTC()

			err := box.Update(&goods)
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
				Action:  "select by ext id",
				Message: "more than 1",
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

func (api *ApiV1) GoodsGet(w http.ResponseWriter, r *http.Request) {

	var gs []*models.Goods
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "Goods",
		WebMethod: "get",
		DateUTC:   time.Now().UTC().UTC()}

	// obx, err := objectbox.NewBuilder().Model(models.ObjectBoxModel()).Build()
	// if err != nil {
	// 	sa.Status = http.StatusInternalServerError
	// 	sa.Error = err.Error()
	// 	sa.Send(w)
	// 	return
	// }
	// defer obx.Close()

	box := models.BoxForGoods(api.obx)

	if fvId == "" {
		gs, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.Goods_.ExtId.Equals(fvId, true))
		gs, err = query.Find()
		query.Close()
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

func (api *ApiV1) GoodsDeleteAll(w http.ResponseWriter, r *http.Request) {
	box := models.BoxForGoods(api.obx)
	box.RemoveAll()
	w.WriteHeader(http.StatusOK)
}
