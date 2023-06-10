// All fuctions that deal with the cost pooling module
package main

import (
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/pierrre/geohash"
)

func getHash(loc Location) string {
        return geohash.Encode(loc.Lat, loc.Lng, 7)
}


func createPool(pool PoolPost, collection *firestore.CollectionRef) error {
    pool.Geohash = getHash(pool.Loc)

    if err := addDoc(collection, pool); err != nil {
        return fmt.Errorf("Failed to add a new document to the firestore doc - \n%v\n: %v", pool, err)
    }
    return nil
}

func joinPool(user User, poolId string, collenction *firestore.CollectionRef) error {
    docRef, doc, err := getDocRef(collenction, poolId)
    if err != nil {
        return fmt.Errorf("Document could not reach the queary: %v", err)
    }

    var data PoolPost
    doc.DataTo(&data)

    if findUser(user, data.Members) == -1 {
        data.Members = append(data.Members, user)
        if idx := findUser(user, data.Requests); idx != -1 {
            data.Requests = removeUser(idx, data.Requests)
        }
        updateDoc(docRef, data)
    }

    return nil
}
