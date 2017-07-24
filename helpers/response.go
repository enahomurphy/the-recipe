package helpers

import (
	"fmt"
	"net/http"
)

// ResponseWriter handles request response
// sets the content type, the status code
// and the error message
func ResponseWriter(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, message)
}

// ServerError sends a 500  handles server error response
// when error triggered is from an unknown
// source or could possibly be an unknown
// error this method is called
// takes the response writer and the error
func ServerError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	errMsg := RespondMessage(err.Error(), status)
	ResponseWriter(w, status, errMsg)
}

// DecoderErrorResponse handles server error response
// When decoding user request fails, this is mostly
// due to incorrect header
// takes the response write as the only parameter
func DecoderErrorResponse(w http.ResponseWriter) {
	errMsg := fmt.Errorf("%s", "Error parsing body set content-type to application/json")
	ServerError(w, errMsg)
}

// BadRequest Sends a 400 request to the user
// this is mostly triggered when the user sends a invalid
// data to the server
func BadRequest(w http.ResponseWriter, errMsg error) {
	// Error interface
	w.Header().Set("Content-type", "Application/json")
	response := RespondMessages(errMsg.Error(), http.StatusBadRequest)
	fmt.Fprint(w, response)
}

// StatusNotFound Sends a 404 request to the user
// this is mostly triggered when the resource the user
// is trying to get dose not exist
func StatusNotFound(w http.ResponseWriter, err error) {
	errMsg := RespondMessage(err.Error(), http.StatusNotFound)
	ResponseWriter(w, http.StatusNotFound, errMsg)
}

// StatusOk Sends a 200 request to the user
// this is mostly triggered when the user
// access a resource
func StatusOk(w http.ResponseWriter, data interface{}) {
	msg := RespondWithData("", 200, data)
	ResponseWriter(w, http.StatusOK, msg)
}

// StatusOkMessage Sends a 200 request to the user
// this is mostly triggered when the user
// access a resource
func StatusOkMessage(w http.ResponseWriter, message string) {
	msg := RespondMessages(message, http.StatusOK)
	ResponseWriter(w, http.StatusOK, msg)
}

// StatusCreated Sends a 201 request to the user
// this is triggered when a resource is created
// is trying to get dose not exist
func StatusCreated(w http.ResponseWriter, msg string) {
	ResponseWriter(w, http.StatusCreated, msg)
}

// Unauthorized Sends a 401 request to the user
// this is triggered when the user making the request
// has not logged in
func Unauthorized(w http.ResponseWriter, message string) {
	msg := RespondMessages(message, http.StatusUnauthorized)
	ResponseWriter(w, http.StatusUnauthorized, msg)
}
