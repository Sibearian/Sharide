// ride fuction handlers
package app

import (
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
            "status"  : "ERROR",
            "error"   : "json is in wrong format",
        })
        return
    }

    docRef, err := share.CreateShare(newShare, rideRef)
    if err != nil {
        sendData(w, http.StatusInternalServerError, map[string]string{
            "status"  : "ERROR",
            "error"   : "Server Error",
        })
        return
    }

    sendData(w, http.StatusOK, map[string]string{
        "pool_id" : docRef.ID,
    })
    return 
}

func getRiders(w http.ResponseWriter, r *http.Request) {
    var ride ReqRide

    err := json.NewDecoder(r.Body).Decode(&ride)
    if err != nil {
        sendData(w, http.StatusBadRequest, map[string]string{
            "status"  : "ERROR",
            "error"   : "json is in wrong format",
        })
        return
    }

    rides, err := share.GetPassangers(ride.Rider, ride.Dist, rideRef)
    if err != nil {
        sendData(w, http.StatusInternalServerError, map[string]string{
            "status"  : "ERROR", 
            "error"   : fmt.Sprintf("%v", err),
        })
        return
    }

    sendData(w, http.StatusOK, rides)
    return 
}
