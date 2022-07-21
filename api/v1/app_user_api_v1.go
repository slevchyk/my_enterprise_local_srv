package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/slevchyk/my_enterprise_local_srv/dao"
	"github.com/slevchyk/my_enterprise_local_srv/models"
)

func (apiV1 *ApiV1) AppUserPost(w http.ResponseWriter, r *http.Request) {

	var aus []models.AppUser
	var err error

	fvIsMain := r.FormValue("is_main")

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

		if fvIsMain == "false" {

			cu := map[string]interface{}{
				"id_settings": 2,
				"phone":       v.Phone,
				"pin":         v.Password}

			bs, err := json.Marshal(cu)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			hc := &http.Client{}
			// TODO: change to config
			// URL := cfg.MainSrv + "/api/clouddbuser"
			URL := "http://95.217.41.66:8811/api/clouddbuser"
			b := bytes.NewBuffer(bs)
			req, err := http.NewRequest("POST", URL, b)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// TODO: change to config file
			// req.SetBasicAuth(cfg.MainAuth.User, cfg.MainAuth.Password)

			req.SetBasicAuth("barkom", "^8Y!e4v?nqC3sAJ]")
			resp, err := hc.Do(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if resp.StatusCode != http.StatusOK {
				bs, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Error(w, string(bs), resp.StatusCode)
				return
			}
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

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
		query.Close()
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

func (api *ApiV1) AppUserAppGet(w http.ResponseWriter, r *http.Request) {

	var aus []*models.AppUser
	var err error

	fvToken := r.FormValue("token")
	au, err := models.GetAppUserByToken(api.obx, fvToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sa := models.ServerAnswer{Object: "AppUser",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForAppUser(api.obx)

	query := box.Query(models.AppUser_.ExtId.Equals(au.ExtId, true))
	aus, err = query.Find()
	query.Close()
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		query.Close()
		return
	}

	query = box.Query(models.AppUser_.ExtId.NotEquals(au.ExtId, true))
	aues, err := query.Find()
	query.Close()
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		query.Close()
		return
	}

	for _, v := range aues {
		aus = append(aus, v.CopyToExport())
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

	query := box.Query(models.AppUser_.Phone.Equals(fvPhone, true), models.AppUser_.Password.Equals(fvPassword, true))
	aus, err := query.Find()
	query.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		query.Close()
		return
	}

	if len(aus) == 0 {
		http.Error(w, "access denied", http.StatusUnauthorized)
		return
	}

	if len(aus) > 1 {
		http.Error(w, "found more than 1", http.StatusConflict)
		return
	}

	au := aus[0]
	au.Token = uuid.NewString()
	au.TokenExpirationDate = time.Now().UTC().AddDate(0, 0, 7)

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

func (api *ApiV1) AppUserCniRecipientPost(w http.ResponseWriter, r *http.Request) {
	var err error
	var existaucr *models.AppUserCniRecipient
	var aucrs []*models.AppUserCniRecipient
	var aucris []models.AppUserCniRecipientImport

	sa := models.ServerAnswer{
		Object:    "AppUserCniRecipient",
		WebMethod: "post",
		DateUTC:   time.Now().UTC()}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	err = json.Unmarshal(bs, &aucris)
	if err != nil {
		sa.Status = http.StatusInternalServerError
		sa.Error = err.Error()
		sa.Send(w)
		return
	}

	box := models.BoxForAppUserCniRecipient(api.obx)

	for _, v := range aucris {

		aucr, pd := parseAppUserCniRecipient(api.obx, v)

		if len(pd.Messages) > 0 {
			pd.Status = http.StatusBadRequest
			sa.ProcessedData = append(sa.ProcessedData, pd)
			continue
		}

		if aucr.Id > 0 {
			existaucr, err = box.Get(aucr.Id)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "AppUseraucrrecipient",
					DataId:   fmt.Sprint(aucr.Id),
					Action:   "query by id",
					Message:  err.Error(),
				})

				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}
			aucrs = append(aucrs, existaucr)
		} else if aucr.ExtId != "" {
			query := box.Query(models.AppUserCniRecipient_.ExtId.Equals(aucr.ExtId, true))
			aucrs, err = query.Find()
			query.Close()
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "ConsignmentNoteIn",
					DataId:   fmt.Sprint(aucr.Id),
					Action:   "query by ext id",
					Message:  err.Error(),
				})

				query.Close()

				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}
		}

		if len(aucrs) == 0 {

			if aucr.CreatedAt.IsZero() {
				aucr.CreatedAt = time.Now().UTC()
				aucr.UpdatedAt = time.Now().UTC()
			}

			srvId, err := box.Put(&aucr)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "AppUserCniRecipient",
					DataId:   fmt.Sprint(srvId),
					Action:   "insert",
					Message:  err.Error(),
				})

				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

			pd.Status = http.StatusOK
			pd.SrvId = srvId
			sa.ProcessedData = append(sa.ProcessedData, pd)

		} else if len(aucrs) == 1 {

			aucr.Id = aucrs[0].Id
			aucr.CreatedAt = aucrs[0].CreatedAt
			aucr.UpdatedAt = time.Now()

			err := box.Update(&aucr)
			if err != nil {
				pd.Status = http.StatusInternalServerError
				pd.Messages = append(pd.Messages, models.ServerMessage{
					DataType: "AppUserCniRecipient",
					DataId:   fmt.Sprint(aucr.Id),
					Action:   "update",
					Message:  err.Error(),
				})

				sa.ProcessedData = append(sa.ProcessedData, pd)
				continue
			}

			pd.Status = http.StatusOK
			pd.SrvId = aucrs[0].Id
			sa.ProcessedData = append(sa.ProcessedData, pd)

		} else {

			pd.Status = http.StatusInternalServerError
			pd.Messages = append(pd.Messages, models.ServerMessage{
				DataType: "AppUserCniRecipient",
				DataId:   fmt.Sprint(aucr.Id),
				Action:   "select",
				Message:  "more than 1",
			})

			sa.ProcessedData = append(sa.ProcessedData, pd)
		}
	}

	sa.Send(w)
}

func (api *ApiV1) AppUserCniRecipientGet(w http.ResponseWriter, r *http.Request) {

	var aus []*models.AppUserCniRecipient
	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "AppUserCniRecipient",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForAppUserCniRecipient(api.obx)

	if fvId == "" {
		aus, err = box.GetAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {
		query := box.Query(models.AppUserCniRecipient_.ExtId.Equals(fvId, true))
		aus, err = query.Find()
		query.Close()
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

func (api *ApiV1) AppUserCniRecipientDelete(w http.ResponseWriter, r *http.Request) {

	var err error

	fvId := r.FormValue("id")

	sa := models.ServerAnswer{Object: "AppUserCniRecipient",
		WebMethod: "get",
		DateUTC:   time.Now().UTC()}

	box := models.BoxForAppUserCniRecipient(api.obx)

	if fvId == "" {
		err = box.RemoveAll()
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			sa.Send(w)
			return
		}

	} else {

		id, err := strconv.Atoi(fvId)
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			return
		}

		err = box.RemoveId(uint64(id))
		if err != nil {
			sa.Status = http.StatusInternalServerError
			sa.Error = err.Error()
			return
		}
	}

	w.WriteHeader(http.StatusOK)

}

func parseAppUserCniRecipient(obx *objectbox.ObjectBox, aucri models.AppUserCniRecipientImport) (models.AppUserCniRecipient, models.ServerProcessedData) {

	var pd models.ServerProcessedData
	var sm models.ServerMessage

	pd.SrvId = aucri.Id
	pd.ExtId = aucri.ExtId

	isDataError := false

	if aucri.ExtId == "" {

		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "AppUseraucrrecipient",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true

	}

	var appUser *models.AppUser
	if aucri.AppUserId == "" {

		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "AppUser",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true

	} else {
		appUser, sm = dao.GetAppUserByExtId(obx, aucri.AppUserId)
		if appUser == nil {
			sm.DataType = "AppUser"
			sm.DataId = aucri.AppUserId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	var recipient *models.Storage
	if aucri.RecipientId == "" {
		pd.Messages = append(pd.Messages, models.ServerMessage{
			DataType: "Storage",
			Action:   "checking data",
			Message:  "ext id isn't specified",
		})

		isDataError = true
	} else {
		recipient, sm = dao.GetStorageByExtId(obx, aucri.RecipientId)
		if recipient == nil {
			sm.DataType = "Storage"
			sm.DataId = aucri.RecipientId
			sm.Action = "db select by ext id"
			sm.Message = "not found"
			pd.Messages = append(pd.Messages, sm)

			isDataError = true
		}
	}

	aucr := models.AppUserCniRecipient{
		Id:        aucri.Id,
		ExtId:     aucri.ExtId,
		AppUser:   appUser,
		Recipient: recipient,
		IsActive:  aucri.IsActive,
	}

	pd.Status = http.StatusOK
	if isDataError {
		pd.Status = http.StatusBadRequest
	}

	return aucr, pd
}
