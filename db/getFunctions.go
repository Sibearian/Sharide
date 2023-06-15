// All the functions to get data
package db

import (
	"fmt"

	"cloud.google.com/go/firestore"
)

func GetCollectionRef(client *firestore.Client, collection string) *firestore.CollectionRef {
    return client.Collection(collection)
}

func GetDocRef(collection *firestore.CollectionRef, docId string) (*firestore.DocumentRef, *firestore.DocumentSnapshot, error) {
    docRef := collection.Doc(docId)

    doc, err := docRef.Get(ctx)
    if err != nil {
        return docRef, nil, fmt.Errorf("In getDocRef: %v", err)
    }

    return docRef, doc, nil
}

func GetQueryDocs(q firestore.Query) ([]*firestore.DocumentSnapshot, error){
    return q.Documents(ctx).GetAll()
}
