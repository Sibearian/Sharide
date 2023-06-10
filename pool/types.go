package pool

import (
	model "ShaRide/models"
)

type Pool struct {
    model.User
    Seats       uint8           `firestore:"seats" json:"seats"`
    Time        int64           `firestore:"time_posted" json:"time_posted"`
    TTL         int16           `firestore:"ttl" json:"wait_till"`
    Start       model.Location  `firestore:"start" json:"start"`
    End         model.Location  `firestore:"drop" json:"drop"`
    Requests    []model.User    `firestore:"requests" json:"requests"`
    Members     []model.User    `firestore:"members" json:"members"`
}
