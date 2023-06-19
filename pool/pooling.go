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

func GetPools(gender uint8, start, end model.Location, distance float64, poolRef *firestore.CollectionRef) (map[string]Pool, error) {
    start.Hash()
    var res = make(map[string]Pool)

    var startHashs []string
    if distance <= 150 {
        startHashs = []string{start.GeoHash}
    } else if distance <= 300 {
        startHashs = Get300mBox(start.GeoHash)
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
    for _, doc := range docs {
        err := doc.DataTo(&pool)
        if err != nil {
            continue
        }
        if pool.End.DistanceTo(end) <= distance && len(pool.Members) < int(pool.Seats) {
            res[doc.Ref.ID] = pool
        }
    }
    return res, nil
}

func StartPool(poolid string, poolRef *firestore.CollectionRef) error {
    return db.UpdateDocField(poolRef, poolid, []firestore.Update{
        {
            Path: "ride_status",
            Value: 1,
        },
    })
}

func EndPool(poolid string, poolRef *firestore.CollectionRef) error {
    return db.UpdateDocField(poolRef, poolid, []firestore.Update{
        {
            Path: "ride_status",
            Value: 2,
        },
        {
            Path: "pool_complition_time",
            Value: firestore.ServerTimestamp,
        },
    })
}


func JoinPool(user model.UserSlice, poolId string, collection *firestore.CollectionRef) error {
    _, doc, err := db.GetDocRef(collection, poolId)
    if err != nil {
        return fmt.Errorf("Document could be found: %v", err)
    }

    var data Pool
    doc.DataTo(&data)

    go userJoined(user, collection)

    return nil
}

func ReqJoinPool(user model.UserSlice, poolId string, collection *firestore.CollectionRef) (model.Location, error) {
    docRef, doc, err := db.GetDocRef(collection, poolId)
    if err != nil {
        var temp model.Location
        return temp, fmt.Errorf("Document could be found: %v", err)
    }

    var data Pool
    doc.DataTo(&data)

    if FindUser(user, data.Requests) == -1 {
        data.Requests = append(data.Requests, user)
        db.UpdateDoc(docRef, data)
    }

    return data.Start, nil 
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

func userJoined(user model.UserSlice, colRef *firestore.CollectionRef) {
    q := colRef.Where("requests", "array-contains", user)
    snaps, _ := db.GetQueryDocs(q)

    var data Pool
    for _, snap := range snaps {
        snap.DataTo(&data)
        if idx := FindUser(user, data.Requests); idx != -1 {
            data.Requests = RemoveUser(idx, data.Members)
            db.UpdateDoc(snap.Ref, data)
        }
    }
}
