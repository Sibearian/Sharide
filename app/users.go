package app

import (
	"ShaRide/users"
	"encoding/json"
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