package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) StoragePost(w http.ResponseWriter, r *http.Request) {

	var ss []models.Storage
	var err error

	sa := models.ServerAnswer{SourceType: "Storage",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &ss)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	// obx, err := objectbox.NewBuilder().Model(models.ObjectBoxModel()).Build()
	// if err != nil {
	// 	sa.Status = http.StatusInternalServerError
	// 	sa.Error = err.Error()
	// 	sa.Send(w)
	// 	return
	// }
	// defer obx.Close()

	box := models.BoxForStorage(apiV1.obx)

	for _, v := range ss {

		query := box.Query(models.Storage_.ExtId.Equals(v.ExtId, true))
		storages, err := query.Find()
		if err != nil {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				SourceId: v.ExtId,
				Action:   "query",
				Message:  err.Error(),
			})
			query.Close()
			continue
		}

		if len(storages) == 0 {
			v.CreatedAt = time.Now().UTC()
			v.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&v)
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					Status:   http.StatusInternalServerError,
					SourceId: v.ExtId,
					Action:   "insert",
					Message:  err.Error(),
				})
			}

		} else if len(storages) == 1 {
			v.Id = storages[0].Id
			v.CreatedAt = storages[0].CreatedAt
			v.UpdatedAt = time.Now().UTC()

			err := box.Update(&v)
			if err != nil {
				sa.Messages = append(sa.Messages, models.ServerMessage{
					Status:   http.StatusInternalServerError,
					SourceId: v.ExtId,
					Action:   "update",
					Message:  err.Error(),
				})
			}
		} else {
			sa.Messages = append(sa.Messages, models.ServerMessage{
				Status:   http.StatusBadRequest,
				SourceId: v.ExtId,
				Action:   "more than 1",
				Message:  err.Error(),
			})
		}

		query.Close()
	}

	sa.Send(w)
}

func (api *ApiV1) StorageGet(w http.ResponseWriter, r *http.Request) {

	var ss []*models.Storage
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{SourceType: "Storage",
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

	box := models.BoxForStorage(api.obx)

	if fvId == "" {
		ss, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.Storage_.ExtId.Equals(fvId, true))
		ss, err = query.Find()
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
