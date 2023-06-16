package pool

import (
	model "ShaRide/models"
)

type Pool struct {
    User        string              `firestore:"userid" json:"userid"`
    Seats       uint8               `firestore:"seats" json:"seats"`
    PrefGender  uint8               `firestore:"pref_gender" json:"pref_gender"`
    RideStatus  uint8               `firestore:"ride_status" json:"status"`

    Time        int64               `firestore:"time_posted" json:"time_posted"`
    TTL         int64               `firestore:"wait_till" json:"wait_till"`
    CompletedT  int64               `firestore:"pool_complition_time"`

    Start       model.Location      `firestore:"start" json:"start"`
    End         model.Location      `firestore:"drop" json:"drop"`
    Requests    []model.UserSlice   `firestore:"requests" json:"requests"`
    Members     []model.UserSlice   `firestore:"members" json:"members"`
}
