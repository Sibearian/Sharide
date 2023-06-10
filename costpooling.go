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

func joinPool(user User, poolId string, collection *firestore.CollectionRef) error {
    docRef, doc, err := getDocRef(collection, poolId)
    if err != nil {
        return fmt.Errorf("Document could be found: %v", err)
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

func leavePool(user User, poolId string, collection *firestore.CollectionRef) error {
    docRef, doc, err := getDocRef(collection, poolId)
    if err != nil {
        return fmt.Errorf("Document could be found: %v", err)
    }

    var data PoolPost
    doc.DataTo(&data)

    if idx := findUser(user, data.Members); idx != -1 {
        data.Members = removeUser(idx, data.Members)
        updateDoc(docRef, data)
    }

    return nil
}
