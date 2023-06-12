// Functions to porvide an interface between the api and the function
package app

import (
	"ShaRide/pool"
	"encoding/json"
	"net/http"
)


// Api Handel for creating pool endpoint
func createPool(w http.ResponseWriter, r *http.Request){
    var newPool pool.Pool
    err := json.NewDecoder(r.Body).Decode(&newPool)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "json is in wrong format",
        })
        return
    }

    docRef, err := pool.CreatePool(newPool, poolRef)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "Server Error",
        })
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "pool_id" : docRef.ID,
        "Status"  : "OK",
    })

}


// Api Handel for joining pool endpoint
func joinPool(w http.ResponseWriter, r *http.Request) {
    var userReq PoolReq
    err := json.NewDecoder(r.Body).Decode(&userReq)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "json is in wrong format",
        })
        return
    }

    if err = pool.JoinPool(userReq.ReqUser, userReq.PoolId, poolRef); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "Server Error",
        })
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "Status" : "OK",
    })
    
}

// Api Handel for leave pool endpoint
func leavePool(w http.ResponseWriter, r *http.Request) {
    var userReq PoolReq

    err := json.NewDecoder(r.Body).Decode(&userReq)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "json is in wrong format",
        })
        return
    }

    if err = pool.LeavePool(userReq.ReqUser, userReq.PoolId, poolRef); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status"  : "ERROR",
            "error"   : "Server Error",
        })
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "Status" : "OK",
    })
}
