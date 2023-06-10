// All fuctions that deal with the cost pooling module
package pool

import (
	"fmt"
    "ShaRide/db"
    model "ShaRide/models"

	"cloud.google.com/go/firestore"
)



func CreatePool(pool Pool, collection *firestore.CollectionRef) error {
    pool.Start.Hash()
    pool.End.Hash()

    if err := db.AddDoc(collection, pool); err != nil {
        return fmt.Errorf("Failed to add a new document to the firestore doc - \n%v\n: %v", pool, err)
    }
    return nil
}


func JoinPool(user model.User, poolId string, collection *firestore.CollectionRef) error {
    docRef, doc, err := db.GetDocRef(collection, poolId)
    if err != nil {
        return fmt.Errorf("Document could be found: %v", err)
    }

    var data Pool
    doc.DataTo(&data)

    if FindUser(user, data.Members) == -1 {
        data.Members = append(data.Members, user)
        if idx := FindUser(user, data.Requests); idx != -1 {
            data.Requests = RemoveUser(idx, data.Requests)
        }
        db.UpdateDoc(docRef, data)
    }

    return nil
}


func LeavePool(user model.User, poolId string, collection *firestore.CollectionRef) error {
    docRef, doc, err := db.GetDocRef(collection, poolId)
    if err != nil {
        return fmt.Errorf("Document could be found: %v", err)
    }

    var data Pool
    doc.DataTo(&data)

    if idx := FindUser(user, data.Members); idx != -1 {
        data.Members = RemoveUser(idx, data.Members)
        db.UpdateDoc(docRef, data)
    }

    return nil
}
