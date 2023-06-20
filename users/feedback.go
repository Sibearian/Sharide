package users

import (
	"ShaRide/db"
	"ShaRide/models"

	"cloud.google.com/go/firestore"
)

func UpdateFeedBack(rating uint8, userid string, userRef *firestore.CollectionRef) error {
	_, snap, err := db.GetDocRef(userRef, userid)
	if err != nil {
		return err
	}

	var user models.User
	snap.DataTo(&user)

	return db.UpdateDocField(userRef, userid, []firestore.Update{
		{
			Path: "rating",
			Value: (float32(user.Number) * user.Rating + float32(rating)) / float32(user.Number + 1),
		},
		{
			Path: "number",
			Value: user.Number + 1,
		},
	})
}

func GetUser(userid string, userRef *firestore.CollectionRef) (*models.User, error) {
	_, snap, err := db.GetDocRef(userRef, userid)
	if err != nil {
		return nil, err
	}

	var user models.User
	snap.DataTo(&user)

	return &user, nil
}

func CreateUser(user models.User, userRef *firestore.CollectionRef) (error) {
	return db.SetDoc(userRef, user.Userid, user)
}