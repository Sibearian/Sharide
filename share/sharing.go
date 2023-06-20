// All function that deal with the sharing module
package share

import (
	"ShaRide/db"
	"ShaRide/models"
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

	endFilter := firestore.PropertyFilter{
		Path:     "end.hash",
		Operator: "in",
		Value:    endHashs,
	}

	genderFilter := firestore.PropertyFilter{
		Path:     "pref_gender",
		Operator: "==",
		Value:    user.User.Gender,
	}

	gen2Filter := firestore.PropertyFilter{
		Path:     "user.gender",
		Operator: "==",
		Value:    user.PrefGen,
	}

	statusFilter := firestore.PropertyFilter{
		Path:     "ride_status",
		Operator: "==",
		Value:    0,
	}

	q := firestore.AndFilter{
		Filters: []firestore.EntityFilter{endFilter, genderFilter, statusFilter, gen2Filter},
	}

	snaps, err := db.GetQueryDocs(shareRef.WhereEntity(q))
	if err != nil {
		return nil, err
	}

	fmt.Println(snaps)

	var filter Share
	for _, snap := range snaps {
		snap.DataTo(&filter)
		fmt.Println(filter.Start.DistanceTo(user.End))
		if filter.Start.DistanceTo(user.End) <= dist {
			continue
		}
		res[snap.Ref.ID] = filter
	}

	return res, nil
}

func StartShare(shareId string, shareRef *firestore.CollectionRef) (res models.Location, err error) {
	err = db.UpdateDocField(shareRef, shareId, []firestore.Update{
		{
			Path:  "ride_status",
			Value: 1,
		},
	})

	_, doc, err := db.GetDocRef(shareRef, shareId)
	var share Share
	err = doc.DataTo(&share)
	return share.Start, err
}

func PickUpShare(shareId string, shareRef *firestore.CollectionRef) (models.Location, error) {
	err := db.UpdateDocField(shareRef, shareId, []firestore.Update{
		{
			Path:  "ride_status",
			Value: 2,
		},
	})

	_, doc, _ := db.GetDocRef(shareRef, shareId)
	var share Share
	doc.DataTo(&share)
	return share.End, err
}

func EndShare(shareid string, shareRef *firestore.CollectionRef) error {
	return db.UpdateDocField(shareRef, shareid, []firestore.Update{
		{
			Path:  "ride_status",
			Value: 3,
		},
	})
}

func LetJoinShare(user models.UserSlice, shareid string, shareRef *firestore.CollectionRef) error {
	_, doc, err := db.GetDocRef(shareRef, shareid)
	if err != nil {
		return fmt.Errorf("Document could be found: %v", err)
	}

	var data Share
	doc.DataTo(&data)

	go userJoined(user, shareRef)

	db.UpdateDocField(shareRef, shareid, []firestore.Update{
		{
			Path:  "accepted",
			Value: user,
		},
	})

	return nil
}

func userJoined(user models.UserSlice, ref *firestore.CollectionRef) {
	q := ref.Where("requests", "array-contains", user)
	snaps, _ := db.GetQueryDocs(q)

	var data Share
	for _, snap := range snaps {
		snap.DataTo(&data)
		if idx := findRider(user, data.Req); idx != -1 {
			data.Req = removeRider(idx, data.Req)
			db.UpdateDoc(snap.Ref, data)
		}
	}
}

func ReqJoinShare(user models.UserSlice, shareId string, shareRef *firestore.CollectionRef) error {
	docRef, doc, err := db.GetDocRef(shareRef, shareId)
	if err != nil {
		return fmt.Errorf("Document could not be found: %v", err)
	}

	var data Share
	doc.DataTo(&data)

	if findRider(user, data.Req) == -1 {
		data.Req = append(data.Req, user)
		db.UpdateDoc(docRef, data)
	}

	return nil
}
