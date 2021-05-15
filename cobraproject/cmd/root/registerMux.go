package root

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/chinmayb/gotraining/cobraproject/pkg/notepad"
	"github.com/gorilla/mux"
)

func HTTPError(body interface{}, err error, w http.ResponseWriter) {
	body := &errorBody{
		Error: err.Error(),
	}
	buf, merr := json.Marshal(body)
	if merr != nil {
		log.Infof("Failed to marshal error message %q: %v", body, merr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Register(addr string, port int, server notepad.NotePadServer) error {
	r := mux.NewRouter()
	s := r.PathPrefix("/notepad").Subrouter()
	s.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		resp, err := server.List(ctx)

	}).Methods("GET")

	s.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var createReq notepad.NotePad
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		if err := json.Unmarshal(b, &createReq); err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
			return
		}
		resp, err := server.Create(ctx, &createReq)

	}).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", addr, port),
	}

	return nil
}
