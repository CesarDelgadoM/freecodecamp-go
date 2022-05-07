package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errorf struct {
	Error string `json:"error"`
}

func Json(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.WriteHeader(statusCode)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		fmt.Fprintf(rw, "%s\n", err.Error())
	}
}

func Error(rw http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Json(rw, statusCode, errorf{
			Error: err.Error(),
		})
	} else {
		Json(rw, statusCode, nil)
	}
}
