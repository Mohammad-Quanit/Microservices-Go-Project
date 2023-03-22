package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Hello struct {
	Msg string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output the field if value is empty
	Date string `json:",omitempty"`
	// convert output to string and rename id
	Id int `json:"id,string"`
}

func HelloHttp(rw http.ResponseWriter, r *http.Request) {
	h := Hello{Msg: "Hello World!", Id: 24}
	data, err := json.Marshal(h)
	if err != nil {
		panic(err.Error())
	}
	// fmt.Fprintf(rw, "Hello %s from response writer", data)
	fmt.Fprint(rw, string(data))
}
