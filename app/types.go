// Types of the api requests
package app

import (
	"ShaRide/models"
	"ShaRide/share"
	"encoding/json"
	"net/http"
)

type UserReq struct {
	User   models.UserSlice `json:"user"`
	PoolId string           `json:"poolid"`
}

type ReqPools struct {
	User models.User     `firestore:"user" json:"user"`
	Pref uint8           `json:"pref_gender"`
	Dist float64         `json:"dist"`
	SLoc models.Location `firestore:"start_location" json:"start_location"`
	ELoc models.Location `firestore:"end_location" json:"end_location"`
}

func sendData(w http.ResponseWriter, status int, send interface{}) {
	(w).WriteHeader(http.StatusOK)
	(w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(send)
}

type ReqRide struct {
	Rider share.Share `json:"rider"`
	Dist  float64     `json:"start_distance"`
}

type shareId struct {
	RideId string `json:"rideid"`
}

type feedbackReq struct {
	Userid string `json:"userid"`
	Rating uint8  `json:"rating"`
}

type userId struct {
	UserId string `json:"userid"`
}
