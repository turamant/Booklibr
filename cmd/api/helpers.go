package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)
type envelope map[string]interface{}

func (app * application) ReadIDParam(r *http.Request) (int64, error){
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid parametrs")
	}
	return id, nil
}

func (app *application) WriteJson(w http.ResponseWriter, status int, data envelope, headers http.Header) error{
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil{
		return err
	}
	js = append(js, '\n')
	for key, value := range headers{
		w.Header()[key] = value
	}
	w.Header().Set("Type-Content", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}