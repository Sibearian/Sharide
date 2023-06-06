package main

import (
	"math"
	"time"
)

type Gender uint8

const (
    male Gender = iota
    female
    any
)

type PoolRequestMember struct {
    ReqTime             time.Time   `json:"reqTime"`
    Name                string      `json:"name"`
    PickUp              Location    `json:"pickup"`
    DropOff             Location    `json:"drop"`
    Gender                          `json:"gender"`
    WaitTillSeconds     int32       `json:"wait"`
}

type Location struct {
    Lat     float64 `json:"lat"`
    Lng     float64 `json:"lng"`
}

func (a Location) DistanceTo(b Location) float64 {
    var rlat1 float64 = a.Lat * (math.Pi / 180)
    var rlat2 float64 = b.Lat * (math.Pi / 180)

    var difflat = rlat2 - rlat1
    var difflon = (a.Lng - b.Lng) * (math.Pi / 180)

    return 2 * 6371.07103 * math.Asin(math.Sqrt(math.Sin(difflat / 2) * math.Sin(difflat / 2) + math.Cos(rlat1) * math.Cos(rlat2) * math.Sin(rlat2) * math.Sin(difflon / 2) * math.Sin(difflon / 2)))
}
