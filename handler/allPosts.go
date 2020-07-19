package handler

import (
"encoding/json"
"fmt"
	"instagram-clone-server/service"
	"net/http"
)

type Posting struct {
	title string
}

func AllPosts(datastore service.PostDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := datastore.GetAllPosts()
		if err != nil {
			fmt.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(ps); err == nil {
			w.Write(payload)
		}
	}
}
