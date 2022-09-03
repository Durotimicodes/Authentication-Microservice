package main

import (
	"errors"
	"fmt"
	"net/http"
)

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	//request payload
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//read the JSON data and serialize it into the requestpayload struct
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	//validate user against the database
	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	} 

	//once the user is found, validate password 
	validPW , err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !validPW {
		//supply the same error message for security
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	//define payload
	payload := jsonResponse{
		Error: false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data: user,
	}

	//write the JSON to standard out
	app.writeJSON(w, http.StatusAccepted, payload)

}
