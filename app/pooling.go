// Functions to porvide an interface between the api and the function
package app

import (
	"ShaRide/models"
	"ShaRide/pool"
	"encoding/json"
	"fmt"
	"net/http"
)

func createPool(w http.ResponseWriter, r *http.Request){
    var newPool pool.Pool
    err := json.NewDecoder(r.Body).Decode(&newPool)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "Status"  : fmt.Sprint(err),
        })
        return
    }

    docRef, err := pool.CreatePool(newPool, poolRef)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "Status"  : fmt.Sprint(err),
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

func joinPool(w http.ResponseWriter, r *http.Request) {
    var userReq models.JoinPoolReq
    err := json.NewDecoder(r.Body).Decode(&userReq)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "Status"  : fmt.Sprint(err),
        })
        return
    }

    if err = pool.JoinPool(userReq.ReqUser, userReq.PoolId, poolRef); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "Status"  : fmt.Sprint(err),
        })
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "Status" : "OK",
    })
    
}
