package app

import (
	"ShaRide/models"
	"ShaRide/users"
	"encoding/json"
	"fmt"
	"net/http"
)

func giveRating(w http.ResponseWriter, r *http.Request){
	var rating feedbackReq
	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = users.UpdateFeedBack(rating.Rating, rating.Userid, userRef); err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  "Server Error",
		})
		return
	}

	sendData(w, http.StatusOK, map[string]string{
		"Status": "OK",
	})
	return
}

func setUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = users.SetUser(user, userRef); err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, map[string]string{
		"Status": "OK",
	})
	return
}

func getUser(w http.ResponseWriter, r *http.Request){
	var user userId
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	userDoc, err := users.GetUser(user.UserId, userRef)
	if ; err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, map[string]models.User{
		userDoc.Userid : *userDoc,
	})
	return
}