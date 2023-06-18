// All function that deal with the sharing module
package share

import (
	"ShaRide/db"
	"ShaRide/pool"
	"fmt"

	"cloud.google.com/go/firestore"
)

func CreateShare(user Share, shareRef *firestore.CollectionRef) (*firestore.DocumentRef, error) {
    user.Start.Hash()
    user.End.Hash()

    docRef, err := db.AddDoc(shareRef, user)
    if err != nil {
        return nil, fmt.Errorf("Failed to add a new document to the firestore, doc - \n%v\n: %v", user, nil)
    }
    return docRef, nil
}

func GetPassangers(user Share, dist float64, shareRef *firestore.CollectionRef) (map[string]Share, error) {
    var res = make(map[string]Share)

    user.Start.Hash()
    user.End.Hash()

    var endHashs = pool.Get300mBox(user.End.GeoHash)

    endFilter   := firestore.PropertyFilter{
        Path: "end.hash",
        Operator: "in",
        Value: endHashs,
    }

    genderFilter:= firestore.PropertyFilter{
        Path: "pref_gender",
        Operator: "==",
        Value: user.Gender,
    }

    driverFilter:= firestore.PropertyFilter{
        Path: "is_driver",
        Operator: "==",
        Value: false,
    }

    statusFilter:= firestore.PropertyFilter{
        Path: "ride_status",
        Operator: "==",
        Value: 0,
    }

    q := firestore.AndFilter{
        Filters: []firestore.EntityFilter{endFilter, genderFilter, driverFilter, statusFilter},
    }

    snaps, err := db.GetQueryDocs(shareRef.WhereEntity(q))
    if err != nil {
        return nil, err
    }

    var filter Share 
    // filter based on the passenger filter also
    for _, snap := range snaps {
        snap.DataTo(&filter)        
        fmt.Println(filter.Start.DistanceTo(user.End))
        if filter.Start.DistanceTo(user.End) >= dist {
            continue
        }
        res[snap.Ref.ID] = filter
    }

    return res, nil
}

func StartShare(shareId string, shareRef *firestore.CollectionRef) error {
    return db.UpdateDocField(shareRef, shareId, []firestore.Update{
        {
            Path: "ride_status",
            Value: 1,
        },
    })
}

func EndPool(shareid string, shareRef *firestore.CollectionRef) error {
    return db.UpdateDocField(shareRef, shareid, []firestore.Update{
        {
            Path: "ride_status",
            Value: 2,
        }, 
    })
}
