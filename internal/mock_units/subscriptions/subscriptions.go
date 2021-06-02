// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/twitchdev/twitch-cli/internal/database"
)

type Endpoint struct{}

var db database.CLIDatabase

func (e Endpoint) Path() string { return "/subscriptions" }

func (e Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db = r.Context().Value("db").(database.CLIDatabase)

	switch r.Method {
	case http.MethodGet:
		getStreams(w, r)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getStreams(w http.ResponseWriter, r *http.Request) {
	s, err := db.NewQuery(r, 100).GetSubscriptions(database.Subscription{})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	j, err := json.Marshal(s)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(j)
}
