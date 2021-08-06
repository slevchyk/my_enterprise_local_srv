package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServerAnswer struct {
	Status        int                   `json:"status"`
	Object        string                `json:"object"`
	WebMethod     string                `json:"web_method"`
	DateUTC       time.Time             `json:"date_utc"`
	Error         string                `json:"error"`
	ProcessedData []ServerProcessedData `json:"processed_data"`
}

type ServerProcessedData struct {
	Status       int             `json:"status"`
	ChangedByAcc bool            `json:"changed_by_app"`
	AppId        string          `json:"app_id"`
	SrvId        uint64          `json:"srv_id"`
	ExtId        string          `json:"ext_id"`
	Messages     []ServerMessage `json:"messages"`
}

type ServerMessage struct {
	DataType string `json:"data_type"`
	DataId   string `json:"data_id"`
	RowId    string `json:"row_id"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

func (sa ServerAnswer) Send(w http.ResponseWriter) {

	if sa.Status == 0 {
		sa.Status = http.StatusOK
	}

	if sa.DateUTC.IsZero() {
		sa.DateUTC = time.Now().UTC()
	}

	bs, err := json.Marshal(sa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(sa.Status)
	w.Write(bs)
}
