// ride fuction handlers
package app

import (
	"ShaRide/models"
	"ShaRide/share"
	"encoding/json"
	"fmt"
	"net/http"
)

func createRide(w http.ResponseWriter, r *http.Request) {
	var newShare share.Share
	err := json.NewDecoder(r.Body).Decode(&newShare)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	docRef, err := share.CreateShare(newShare, rideRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  "Server Error",
		})
		return
	}

	sendData(w, http.StatusOK, map[string]string{
		"pool_id": docRef.ID,
	})
	return
}

func getRiders(w http.ResponseWriter, r *http.Request) {
	var ride ReqRide

	err := json.NewDecoder(r.Body).Decode(&ride)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	rides, err := share.GetPassangers(ride.Rider, ride.Dist, rideRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, rides)
	return
}

func startRide(w http.ResponseWriter, r *http.Request) {
	var req shareId

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	loc, err := share.StartShare(req.RideId, rideRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, map[string]models.Location{
		"start": loc,
	})
	return
}

func pickupRide(w http.ResponseWriter, r *http.Request) {
	var req shareId

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	loc, err := share.PickUpShare(req.RideId, rideRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, map[string]models.Location{
		"end": loc,
	})
	return
}

func endRide(w http.ResponseWriter, r *http.Request) {
	var req shareId

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	err := share.EndShare(req.RideId, rideRef)
	if err != nil {
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

func reqJoinRide(w http.ResponseWriter, r *http.Request) {
	var user UserReq

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	err := share.ReqJoinShare(user.User, user.PoolId, rideRef)
	if err != nil {
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

func letJoinRide(w http.ResponseWriter, r *http.Request) {
	var user UserReq

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	err := share.LetJoinShare(user.User, user.PoolId, rideRef)
	if err != nil {
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
