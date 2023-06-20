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
var rideRef *firestore.CollectionRef
var userRef *firestore.CollectionRef

func Start() {
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
	rideRef = db.GetCollectionRef(fs, "ride")
	userRef = db.GetCollectionRef(fs, "users")

	// Pool apis
	router.HandleFunc("/pool/create", createPool).Methods("POST")
	router.HandleFunc("/pool/let_join", joinPool).Methods("POST")
	router.HandleFunc("/pool/leave", leavePool).Methods("POST")
	router.HandleFunc("/pool/getpools", getPools).Methods("POST")
	router.HandleFunc("/pool/start", startPool).Methods("POST")
	router.HandleFunc("/pool/end", endPool).Methods("POST")
	router.HandleFunc("/pool/req_join", reqPool).Methods("POST")

	// Ride apis
	router.HandleFunc("/ride/create", createRide).Methods("POST")
	router.HandleFunc("/ride/req_join", reqJoinRide).Methods("POST")
	router.HandleFunc("/ride/let_join", letJoinRide).Methods("POST")
	router.HandleFunc("/ride/getrides", getRiders).Methods("POST")
	router.HandleFunc("/ride/start", startRide).Methods("POST")
	router.HandleFunc("/ride/pickup", pickupRide).Methods("POST")
	router.HandleFunc("/ride/end", endRide).Methods("POST")

	// User api
	router.HandleFunc("/user/feedback", giveRating).Methods("POST")
	router.HandleFunc("/user/set", setUser).Methods("POST")
	router.HandleFunc("/user/get", getUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
