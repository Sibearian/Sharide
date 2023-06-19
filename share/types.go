// Types to help the user to interact with the sharing service
package share

import "ShaRide/models"

type Share struct {
    User    models.UserSlice`firestore:"user" json:"user"`
    PrefGen uint8           `firestore:"pref_gender" json:"pref_gender"`
    Start   models.Location `firestore:"start" json:"start"`
    End     models.Location `firestore:"end" json:"end"`
    Req     []models.UserSlice`firestore:"req"`
    Accept  string          `firestore:"accepted"`
    RStatus uint8           `firestore:"ride_status"`
}

func findRider(rider models.UserSlice, requests []models.UserSlice) (int) {
    for idx, ride := range requests {
        if ride == rider {
            return idx
        }
    }
    return -1
}

func removeRider(idx int, requests []models.UserSlice) []models.UserSlice {
    requests[idx] = requests[len(requests) - 1] 
    return requests[:len(requests) - 1]
}
