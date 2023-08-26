package handler

import (
	"net/http"
	"strings"

	"github.com/kdsama/rate-limiter/services"
)

type Limiter struct {
	serv services.RateLimiter
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
