package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
