// Functions to porvide an interface between the api and the function
package app

import (
	"ShaRide/pool"
	"encoding/json"
	"fmt"
	"net/http"
)

// Api Handel for creating pool endpoint
func createPool(w http.ResponseWriter, r *http.Request) {
	var newPool pool.Pool
	err := json.NewDecoder(r.Body).Decode(&newPool)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	docRef, err := pool.CreatePool(newPool, poolRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  "Server Error",
		})
		return
	}

	sendData(w, http.StatusOK, map[string]string{
		"pool_id": docRef.ID,
		"Status":  "OK",
	})
	return
}

// Api Handel for joining pool endpoint
func joinPool(w http.ResponseWriter, r *http.Request) {
	var userReq UserReq
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = pool.JoinPool(userReq.User, userReq.PoolId, poolRef); err != nil {
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

// Api Handel for leave pool endpoint
func leavePool(w http.ResponseWriter, r *http.Request) {
	var userReq UserReq

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = pool.LeavePool(userReq.User, userReq.PoolId, poolRef); err != nil {
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

// handle for getting all the pools
func getPools(w http.ResponseWriter, r *http.Request) {
	var poolReq ReqPools

	err := json.NewDecoder(r.Body).Decode(&poolReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	pools, err := pool.GetPools(poolReq.Pref, poolReq.SLoc, poolReq.SLoc, poolReq.Dist, poolRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  fmt.Sprintf("%v", err),
		})
		return
	}

	sendData(w, http.StatusOK, pools)
	return
}

// Api Handel for leave pool endpoint
func startPool(w http.ResponseWriter, r *http.Request) {
	var userReq UserReq

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = pool.StartPool(userReq.PoolId, poolRef); err != nil {
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

func endPool(w http.ResponseWriter, r *http.Request) {
	var userReq UserReq

	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	if err = pool.EndPool(userReq.PoolId, poolRef); err != nil {
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

// Api Handel for joining pool endpoint
func reqPool(w http.ResponseWriter, r *http.Request) {
	var userReq UserReq
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		sendData(w, http.StatusBadRequest, map[string]string{
			"status": "ERROR",
			"error":  "json is in wrong format",
		})
		return
	}

	location, err := pool.ReqJoinPool(userReq.User, userReq.PoolId, poolRef)
	if err != nil {
		sendData(w, http.StatusInternalServerError, map[string]string{
			"status": "ERROR",
			"error":  "Server Error",
		})
		return
	}

	sendData(w, http.StatusOK, location)
	return
}
