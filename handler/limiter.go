package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kdsama/rate-limiter/services"
)

type Limiter struct {
	serv services.RateLimiter
}

type InputLimiter struct {
	Url          string `json:"url" `
	ShortUrl     string `json:"shorturl" `
	Expiry       int64  `json:"expiryTimestamp" `
	Limit        int64  `json:"limit" `
	BrowserCache bool   `json:"bcache" `
}

func NewLimiter() *Limiter {
	return &Limiter{}
}

func (l *Limiter) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		l.HandleGet(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid request"))

	}

}

func (l *Limiter) HandleSave(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		l.HandleSaveLimiter(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid request"))

	}

}

func (l *Limiter) HandleSaveLimiter(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t InputLimiter
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return

	}

	ok := validateLimiter(t)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))
		return
	}
	l.serv.Save("someuser", t.Url, t.BrowserCache, t.Expiry)
}

func (l *Limiter) HandleGet(w http.ResponseWriter, r *http.Request) {
	// read Endpoint
	pathSegments := strings.Split(r.URL.Path, "/")

	// The first element will be an empty string due to the leading slash
	// The second element will be the desired endpoint
	if len(pathSegments) == 2 {
		endpoint := pathSegments[1]
		//
		resp, err, bc := l.serv.Get(endpoint)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		if bc {
			http.Redirect(w, r, resp, http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, resp, http.StatusFound)
		}

		// http.Redirect()

	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid request"))
	}
}

func validateLimiter(t InputLimiter) bool {
	if t.Url == "" {
		return false
	}

	return true
}
