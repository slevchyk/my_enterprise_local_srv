package models

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServerAnswer struct {
	Status     int
	Error      string
	SourceType string
	WebMethod  string
	DateUTC    time.Time	
	Messages   []ServerMessage
}

type ServerMessage struct {
	Status   int    `json:"status"`
	SourceId string `json:"source_id"`
	DataType string `json:"data_type"`
	DataId   string `json:"data_id"`
	Action   string `json:"action"`
	Message  string `json:"message"`
}

func (sa ServerAnswer) Send(w http.ResponseWriter) {

	if sa.Status == 0 {
		if len(sa.Messages) == 0 {
			sa.Status = http.StatusOK
		} else {
			sa.Status = http.StatusBadRequest
		}
	}

	if sa.DateUTC.IsZero() {
		sa.DateUTC = time.Now().UTC()
	}

	bs, err := json.Marshal(sa)
	if err != nil {
		http.Error(w, sa.Error, sa.Status)
		return
	}

	w.WriteHeader(sa.Status)
	w.Write(bs)
}
