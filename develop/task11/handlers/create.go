package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"task11/cache"
	"task11/domain"
	"task11/models"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request, c *cache.Cache) {

	if r.Method != http.MethodPost {
		domain.ErrorLogger(w, errors.New("method error"))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var decoded models.Event

	if decodingBodyErr := decoder.Decode(&decoded); decodingBodyErr != nil {
		domain.ErrorLogger(w, decodingBodyErr)
		return
	}

	dateQuery := decoded.Date
	timeQuery := decoded.Time
	userQuery := decoded.UserID

	if _, errParse := time.Parse("2006-01-02", dateQuery); errParse != nil {
		domain.ErrorLogger(w, errParse)
		return
	}

	if _, errParseTime := time.Parse("15:00", timeQuery); errParseTime != nil {
		domain.ErrorLogger(w, errParseTime)
		return
	}

	event := domain.NewEvent(dateQuery, timeQuery, userQuery)
	c.Create(event)

	domain.ResponseLogger(w, "event created")

}
