// Api endpoints are defined here
package main

import (
	"fmt"
	"log"
    "ShaRide/db"
    model "ShaRide/models"
    "ShaRide/pool"
)

func main() {
    var guy  model.User
    guy.Name = "xzxc"
    guy.Gender = 1
    guy.Rating = 4.5
    guy.Userid = "101"

    var sharePool pool.Pool
    sharePool.User = guy
    sharePool.Start = model.Location{Lat: 12.939756, Lng: 77.565426, GeoHash: "hh"}
    sharePool.End = model.Location{Lat: 12.950123, Lng: 77.573864, GeoHash: "df"}
    sharePool.Seats = 3
    sharePool.Start.Hash()
    sharePool.End.Hash()

    fmt.Print("Hello\n\n")

    app, err := db.CreateNewApp()
    if err != nil {
        log.Fatal(err)
    }

    firestore, err := db.CreateFirestore(app)
    if err != nil {
        log.Fatal(err)
    }
    defer firestore.Close()

    pools := db.GetCollectionRef(firestore, "pools")

    err = pool.CreatePool(sharePool, pools)

    // err = pool.JoinPool(guy, "XkJ1smdXIY6uTA1JxSDf", pools)
    // if err != nil {
        // fmt.Println(err)
    // }
// 
    // err = pool.LeavePool(guy, "XkJ1smdXIY6uTA1JxSDf", pools)
    // if err != nil {
        // fmt.Println(err)
    // }
// 
    // snaps, err := pool.GetExactPools(model.Location{Lat: 12.939756, Lng: 77.565426, GeoHash: "hh"}, model.Location{Lat: 12.950123, Lng: 77.573864, GeoHash: "df"}, pools)
// 
    // if err != nil {
        // fmt.Println(err)
    // }
// 
    // fmt.Printf("Got these Pools :\n%v\n", snaps)
}
