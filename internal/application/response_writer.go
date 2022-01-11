package application

import (
	"encoding/json"
	"net/http"
)

func writeResponse(object interface{}, w http.ResponseWriter) error {
	if json, err := json.Marshal(object); err != nil {
		writeErrorResponse(http.StatusInternalServerError, err, w)
		return err
	} else {
		w.Write(json)
		return nil
	}
}

func writeErrorResponse(errorStatus int, err error, w http.ResponseWriter) error {
	if errorJson, err := json.Marshal(err); err == nil {
		w.WriteHeader(errorStatus)
		w.Write(errorJson)
		return nil
	} else {
		return err
	}
}
