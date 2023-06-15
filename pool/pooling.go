// All fuctions that deal with the cost pooling module
package pool

import (
	"ShaRide/db"
	model "ShaRide/models"
	"fmt"

	"cloud.google.com/go/firestore"
)


func CreatePool(pool Pool, collection *firestore.CollectionRef) (*firestore.DocumentRef, error) {
    pool.Start.Hash()
    pool.End.Hash()
    docRef, err := db.AddDoc(collection, pool)
    if err != nil {
        return nil, fmt.Errorf("Failed to add a new document to the firestore doc - \n%v\n: %v", pool, err)
    }
    return docRef, nil
}


func JoinPool(user model.UserSlice, poolId string, collection *firestore.CollectionRef) error {
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


func LeavePool(user model.UserSlice, poolId string, collection *firestore.CollectionRef) error {
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


func GetPools(gender uint8, start, end model.Location, distance float64, poolRef *firestore.CollectionRef) (res []Pool, err error) {
    start.Hash()
    end.Hash()

    var startHashs []string
    if distance <= 150 {
        startHashs = []string{start.GeoHash}
    } else if distance <= 300 {
        startHashs = get300mBox(start.GeoHash)
    } else if distance <= 600 {    
        startHashs = get600mBox(start.GeoHash)
    } else {
        startHashs = []string{start.GeoHash[0:len(start.GeoHash) -1]}
    }

    startQuery := firestore.PropertyFilter{
        Path: "start.hash",
        Operator: "in",
        Value: startHashs,
    }

    genderFilter    := firestore.PropertyFilter{
        Path: "pref_gender",
        Operator: "==",
        Value: gender,
    }

    statusFilter    := firestore.PropertyFilter{
        Path: "ride_status",
        Operator: "==",
        Value: 0,
    }

    query := firestore.AndFilter{
        Filters: []firestore.EntityFilter{startQuery, genderFilter, statusFilter},
    }

    docs, err := db.GetQueryDocs(poolRef.WhereEntity(query))
    if err != nil {
        return nil, err
    }

    var pool Pool
    var box BoundingBox = GetBoundingBox(end.Lat, end.Lng, distance)
    fmt.Println(box)
    for _, doc := range docs {
        doc.DataTo(&pool)
        fmt.Println(pool.End)
        if pool.End.DistanceTo(end) <= distance {
            res = append(res, pool)
        }
        fmt.Printf("res : %v\n\n", res)
    }


    return res, nil
    
}
