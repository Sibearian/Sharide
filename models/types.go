package models

import (
	"math"

	"github.com/pierrre/geohash"
)

type User struct{
    Userid  string  `firestore:"userid" json:"userid"`
    Name    string  `firestore:"name" json:"name"`
    Gender  uint8   `firestore:"gender" json:"gender"`
    Mail    string  `firestore:"email" json:"mail"`
    Phone   string  `firestore:"phone" json:"phone"`
    Rating  float32 `firestore:"rating" json:"rating"`
    Number  int     `firestore:"number" json:"number"`
    Upiid   string  `firestore:"upi"    json:"upi"`
}

type Location struct {
    Lat     float64 `json:"lat" firestore:"lat"`
    Lng     float64 `json:"lng" firestore:"lng"`
    GeoHash    string  `firestore:"hash"`
}

type UserSlice struct {
    Userid  string  `json:"id" firestore:"id"`
    Gender  uint8   `json:"gender" firestore:"gender"`
}

func (a Location) DistanceTo(b Location) float64 {
    var rlat1 float64 = a.Lat * (math.Pi / 180)
    var rlat2 float64 = b.Lat * (math.Pi / 180)

    var difflat = rlat2 - rlat1
    var difflon = (a.Lng - b.Lng) * (math.Pi / 180)

    return 2 * 6371.07103 * math.Asin(math.Sqrt(math.Sin(difflat / 2) * math.Sin(difflat / 2) + math.Cos(rlat1) * math.Cos(rlat2) * math.Sin(rlat2) * math.Sin(difflon / 2) * math.Sin(difflon / 2)))
}


func (loc *Location) Hash() {
    loc.GeoHash = geohash.Encode(loc.Lat, loc.Lng, 7)
}

