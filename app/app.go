// Applications
package app

import (
	"ShaRide/db"
	"log"
	"net/http"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
)

var app *firebase.App
var fs *firestore.Client
var poolRef *firestore.CollectionRef

func Start(){
    var err error
    router := mux.NewRouter()
    app, err = db.CreateNewApp()
    if err != nil {
        log.Fatal(err)
    }

    fs, err = db.CreateFirestore(app)
    if err != nil {
        log.Fatal(err)
    }
    defer fs.Close()

    poolRef = db.GetCollectionRef(fs, "pools")

    // Pool apis
    router.HandleFunc("/pool/create", createPool).Methods("POST")
    router.HandleFunc("/pool/let_join", joinPool).Methods("POST")
    router.HandleFunc("/pool/leave", leavePool).Methods("POST")
    router.HandleFunc("/pool/getpools", getPools).Methods("POST")
    router.HandleFunc("/pool/start", startPool).Methods("POST")
    router.HandleFunc("/pool/end", endPool).Methods("POST")
    router.HandleFunc("/pool/req_join", reqPool).Methods("POST")

    
    log.Fatal(http.ListenAndServe(":8080", router))
}

