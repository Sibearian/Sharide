// function to manage the firebase fucntions
package db

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


func CreateNewApp() (*firebase.App, error) {
    conf := &firebase.Config{ProjectID: "bms-shareride"}

    app, err := firebase.NewApp(ctx, conf)
    if err != nil {
        return nil, fmt.Errorf("Failed to create a new app: %v", err)
    }

    return app, nil 
}


func CreateFirestore(app *firebase.App) (*firestore.Client, error) {
    client, err := app.Firestore(ctx)
    if err != nil {
        return nil, fmt.Errorf("Failed to create firestore: %v", err)
    }

    return client, nil
}

func AddDoc(collection *firestore.CollectionRef, data interface{}) (*firestore.DocumentRef, error) {
    docRef, _, err := collection.Add(ctx, data)
    if err != nil {
        return nil, fmt.Errorf("An error has occured while adding the document: %v", err)
    }
    return docRef, nil
}

func SetDoc(collection *firestore.CollectionRef, docId string, data interface{}) (error) {
    _, err := collection.Doc(docId).Set(ctx, data)
    if err != nil {
        return fmt.Errorf("An error has occured while adding the document: %v", err)
    }
    return nil
}

func UpdateDoc(docRef *firestore.DocumentRef, data interface{}) error {
    _, err := docRef.Set(ctx, data)
    if err != nil {
        return fmt.Errorf("Could not update the document: %v", err)
    }
    return nil
}


