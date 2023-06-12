package models

import (
	"math"

	"github.com/pierrre/geohash"
)

type User struct{
    Userid  string  `firestore:"userid" json:"userid"`
    Name    string  `firestore:"name" json:"name"`
    Gender  uint8   `firestore:"gender" json:"gender"`
    Rating  float32 `firestore:"rating" json:"rating"`
}

type Location struct {
    Lat     float64 `json:"lat" firestore:"lat"`
    Lng     float64 `json:"lng" firestore:"lng"`
    GeoHash    string  `json:"hash" firestore:"hash"`
}

type JoinPoolReq struct {
    ReqUser User    `json:"user"`
    PoolId  string  `json:"pool_id"`
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

