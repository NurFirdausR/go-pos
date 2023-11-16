package helper

import (
	"encoding/json"
	"net/http"
)

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("content-type", "Application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
