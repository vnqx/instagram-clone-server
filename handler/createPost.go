package handler

import (
	"errors"
	"fmt"
	"instagram-clone-server/service"
	"instagram-clone-server/util"
	"log"
	"net/http"
)

func CreatePost(datastore service.PostDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, err := verifyToken(w, r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var i service.CreatePostInput
		err = util.DecodeJSONBody(w, r, &i)
		if err != nil {
			var mr *util.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				// Default to 500 Internal Server Error
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}
		//Create the posting in the db with the given input
		err = datastore.CreatePost(i, claims.UserId)
		if err != nil {
			fmt.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}