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
    GeoHash    string  `firestore:"hash" json:"-"`
}

type UserSlice struct {
    Userid  string  `json:"id" firestore:"id"`
    Gender  uint8   `json:"gender" firestore:"gender"`
}

func (a Location) DistanceTo(b Location) float64 {
    lat1, lng1 := a.Lat, a.Lng
    lat2, lng2 := b.Lat, b.Lng
    
    radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)
	
	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)
	
	dist := math.Sin(radlat1) * math.Sin(radlat2) + math.Cos(radlat1) * math.Cos(radlat2) * math.Cos(radtheta);
	if dist > 1 {
		dist = 1
	}
	
	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515
    dist = dist * 1.609344
	
	return dist * 1000
}

func (loc *Location) Hash() {
    loc.GeoHash = geohash.Encode(loc.Lat, loc.Lng, 7)
}

