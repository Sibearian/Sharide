// Api endpoints are defined here
package main

import (
	"fmt"
	"log"
)

func main() {
    var guy  User
    guy.Name = "xzxc"
    guy.Gender = 1
    guy.Rating = 4.5
    guy.Userid = "101"

    var pool PoolPost
    pool.User = guy
    pool.Loc  = Location{12.939756, 77.565426}
    pool.Seats = 3

    app, err := createNewApp()
    if err != nil {
        log.Fatal(err)
    }

    firestore, err := createFirestore(app)
    if err != nil {
        log.Fatal(err)
    }

    pools := getCollectionRef(firestore, "pools")
    // if err := createPool(pool, pools); err != nil{
    //    log.Fatal(err)
    //}

    err = joinPool(guy, "XkJ1smdXIY6uTA1JxSDf", pools)
    if err != nil {
        fmt.Println(err)
    }

    err = leavePool(guy, "XkJ1smdXIY6uTA1JxSDf", pools)
    if err != nil {
        fmt.Println(err)
    }
}
