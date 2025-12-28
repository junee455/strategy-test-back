package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func ReadClientMetaData(r *http.Request) (*ClientMetaData, error) {
	var requestBodyJson ClientMetaData

	stAuthCookie, err := r.Cookie("StAuth")

	if err != nil {
		return nil, err
	}

	// var gameId = mux.Vars(r)["id"]
	var stAuthText string
	stAuthText, err = url.QueryUnescape(stAuthCookie.Value)

	if err != nil {
		fmt.Println("unsescape fail")
		return nil, err
	}

	fmt.Printf("StAuth cookie: %v\n", stAuthText)

	if err := json.Unmarshal([]byte(stAuthText), &requestBodyJson); err != nil {
		return nil, err
	}

	return &requestBodyJson, nil
}
