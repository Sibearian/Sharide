package main

import (
	"encoding/json"
	"math"
)

type User struct{
    Userid  string  `firestore:"userid" json:"userid"`
    Name    string  `firestore:"name" json:"name"`
    Gender  uint8   `firestore:"gender" json:"gender"`
    Rating  float32 `firestore:"rating" json:"rating"`
}


type PoolPost struct {
    User
    Seats       uint8       `firestore:"seats" json:"seats"`
    Loc         Location    `firestore:"loc" json:"loc"`
    Geohash     string      `firestore:"geohash" json:"geohash"`
    Requests    []User    `firestore:"participants" json:"participants"`
    Members     []User    `firestore:"members" json:"members"`
}

type Location struct {
    Lat     float64 `json:"lat" firestore:"lat"`
    Lng     float64 `json:"lng" firestore:"Lng"`
}

func (a Location) DistanceTo(b Location) float64 {
    var rlat1 float64 = a.Lat * (math.Pi / 180)
    var rlat2 float64 = b.Lat * (math.Pi / 180)

    var difflat = rlat2 - rlat1
    var difflon = (a.Lng - b.Lng) * (math.Pi / 180)

    return 2 * 6371.07103 * math.Asin(math.Sqrt(math.Sin(difflat / 2) * math.Sin(difflat / 2) + math.Cos(rlat1) * math.Cos(rlat2) * math.Sin(rlat2) * math.Sin(difflon / 2) * math.Sin(difflon / 2)))
}


func structToMap(obj interface{}) (newMap map[string]interface{}, err error) {
    data, err := json.Marshal(obj) // Convert to a json string
    if err != nil {
        return
    }
    err = json.Unmarshal(data, &newMap) // Convert to a map
    return
}
