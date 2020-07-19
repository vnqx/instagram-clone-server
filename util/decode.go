package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"io"
	"log"
	"net/http"
	"strings"
)

type MalformedRequest struct {
	Status int
	Msg    string
}

func (mr *MalformedRequest) Error() string {
	return mr.Msg
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// We are using the package gddo/httputil/header to parse
	// and extract the value here, so the check works
	// even if the client includes additional charset or boundary
	// information in the header.
	if r.Header.Get("Content-Type") != "" {
		v, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if v != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return &MalformedRequest{Status: http.StatusUnsupportedMediaType, Msg: msg}
		}
	}

	// Enforce a maximum read of 1MB from the response body.
	// A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	// If the number of bytes will get modified remember to
	// update 'request body too large' message.
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	// Set up the decoder
	dec := json.NewDecoder(r.Body)

	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec.DisallowUnknownFields()

	// Store the request's decoded body in i
	err := dec.Decode(&dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at positon %d)", syntaxError.Offset)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg:    msg,
			}
		// Catch any type errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)",
				unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg:    msg,
			}
		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and interpolate
		// it in our custom error message.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg:    msg,
			}
		// An io.EOF error is returned by Decode() if the request body
		// is empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &MalformedRequest{
				Status: http.StatusBadRequest,
				Msg:    msg,
			}
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &MalformedRequest{
				Status: http.StatusRequestEntityTooLarge,
				Msg:    msg,
			}
		// Otherwise default to logging the error and returning it
		default:
			log.Println(err.Error())
			return err
		}
	}

	// Call decode again, using a pointer to an empty anonymous struct
	// as the destination. If the request body only contained a single JSON
	// object this will return an io.EOF error. So if we get anything else,
	// we know that there is additional data in the request body.
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &MalformedRequest{
			Status: http.StatusBadRequest,
			Msg:    msg,
		}
	}

	return nil
}