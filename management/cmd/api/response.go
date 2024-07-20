package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (s *server) writeJSON(w http.ResponseWriter, data jsonResponse, status int) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) errorJSON(w http.ResponseWriter, err error, status int) error {
	resPayload := jsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return s.writeJSON(w, resPayload, status)
}
