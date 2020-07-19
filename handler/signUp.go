package handler

import (
	"errors"
	"fmt"
	"instagram-clone-server/service"
	"instagram-clone-server/util"
	"log"
	"net/http"
)

func SignUp(datastore service.UserDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i service.SignUpInput

		// DecodeJSONBody handles the other errors.
		err := util.DecodeJSONBody(w, r, &i)
		if err != nil {
			var mr *util.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				log.Print(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		u, err := datastore.CreateUser(i)
		fmt.Print(err)

		if err != nil {
			msg := "User with this email already exists"
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		handleTokenResponse(w, u.Id)
		w.WriteHeader(http.StatusCreated)
	}
}