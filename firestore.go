// function to manage the firebase fucntions
package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var ctx context.Context

func init(){
    ctx = context.Background()
}


func createNewApp() (*firebase.App, error) {
    conf := &firebase.Config{ProjectID: "mdp-bmsce"}

    app, err := firebase.NewApp(ctx, conf)
    if err != nil {
        return nil, fmt.Errorf("Failed to create a new app: %v", err)
    }

    return app, nil 
}


func createFirestore(app *firebase.App) (*firestore.Client, error) {
    client, err := app.Firestore(ctx)
    if err != nil {
        return nil, fmt.Errorf("Failed to create firestore: %v", err)
    }

    return client, nil
}


func getCollectionRef(client *firestore.Client, collection string) *firestore.CollectionRef {
    return client.Collection(collection)
}


func addDoc(collection *firestore.CollectionRef, data interface{}) error {
    _, _, err := collection.Add(ctx, data)
    if err != nil {
        return fmt.Errorf("An error has occured while adding the document: %v", err)
    }
    return nil
}


func updateDoc(docRef *firestore.DocumentRef, data interface{}) error {
    _, err := docRef.Set(ctx, data)
    if err != nil {
        return fmt.Errorf("Could not update the document: %v", err)
    }
    return nil
}


func getDocRef(collection *firestore.CollectionRef, docId string) (*firestore.DocumentRef, *firestore.DocumentSnapshot, error) {
    docRef := collection.Doc(docId)

    doc, err := docRef.Get(ctx)
    if err != nil {
        return docRef, nil, fmt.Errorf("In getDocRef: %v", err)
    }

    return docRef, doc, nil
}
