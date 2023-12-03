package domain

import (
	"encoding/json"
	"log"
	"net/http"
	time2 "time"

	"task11/models"
)

func NewEvent(date, time string, userID int) *models.Event {
	Uid := time2.Now().Unix()

	return &models.Event{
		UserID: userID,
		Date:   date,
		Time:   time,
		Uid:    Uid,
	}
}

func ErrorLogger(w http.ResponseWriter, errorInput error) {
	log.Println(errorInput)
	w.WriteHeader(http.StatusInternalServerError)
	details := models.Details{ErrCode: http.StatusInternalServerError, ErrMessage: errorInput.Error()}
	response, _ := json.MarshalIndent(models.Error{Err: details}, "", "\t")
	w.Write(response)
}

func ResponseLogger(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	result := models.Result{StatusCode: http.StatusOK, Message: message}
	response, _ := json.MarshalIndent(models.Response{Body: result}, "", "\t")
	w.Write(response)
}
