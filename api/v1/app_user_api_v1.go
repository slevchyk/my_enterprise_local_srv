package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) AppUserPost(w http.ResponseWriter, r *http.Request) {

	var aus []models.AppUser
	var err error

	sa := models.ServerAnswer{
		Object:    "AppUser",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &aus)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForAppUser(apiV1.obx)

	for _, v := range aus {

		pd := models.ServerProcessedData{
			ExtId: v.ExtId,
		}

		isDataError := false

		if v.ExtId == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "ext id is empty",
			})
			isDataError = true
		}

		if v.FirstName == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "first name is empty",
			})
			isDataError = true
		}

		if v.LastName == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "last name is empty",
			})
			isDataError = true
		}

		if v.Phone == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "phone is empty",
			})
			isDataError = true
		}

		if v.Password == "" {
			pd.Messages = append(pd.Messages, models.ServerMessage{
				Action:  "checking value",
				Message: "password is empty",
			})
			isDataError = true
		}

		if isDataError {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		query := box.Query(models.AppUser_.ExtId.Equals(v.ExtId, true))
		appUsers, err := query.Find()
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

		if len(appUsers) == 0 {
			v.CreatedAt = time.Now().UTC()
			v.UpdatedAt = time.Now().UTC()

			_, err := box.Put(&v)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					Action:  "insert",
					Message: err.Error(),
				})

				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

		} else if len(appUsers) == 1 {
			v.Id = appUsers[0].Id
			v.CreatedAt = appUsers[0].CreatedAt
			v.UpdatedAt = time.Now().UTC()

			// pd.SrvId = string(v.Id)
			pd.SrvId = v.Id

			err := box.Update(&v)
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

	sa.Send(w)
}

func (api *ApiV1) AppUserGet(w http.ResponseWriter, r *http.Request) {

	var aus []*models.AppUser
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "AppUser",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForAppUser(api.obx)

	if fvId == "" {
		aus, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.AppUser_.ExtId.Equals(fvId, true))
		aus, err = query.Find()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			query.Close()
			return
		}
	}

	bs, err := json.Marshal(aus)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}

func (api *ApiV1) AppUserAuth(w http.ResponseWriter, r *http.Request) {

	fvPhone := r.FormValue("phone")
	fvPassword := r.FormValue("password")

	errMessage := ""

	if fvPhone == "" {
		errMessage += fmt.Sprintln("phone isn't provided")
	}

	if fvPassword == "" {
		errMessage += fmt.Sprintln("password isn't provided")
	}

	if errMessage != "" {
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	box := models.BoxForAppUser(api.obx)

	query := box.Query(models.AppUser_.Phone.Equals(fvPhone, true), models.AppUser_.Password.Equals(fvPhone, true))
	aus, err := query.Find()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		query.Close()
		return
	}

	if len(aus) > 1 {
		http.Error(w, "found more than 1", http.StatusConflict)
		return
	}

	au := aus[0]
	au.Token = uuid.NewString()
	_, err = box.Put(au)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(aus[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
