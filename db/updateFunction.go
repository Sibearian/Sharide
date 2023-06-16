// Update fuctions
package db

import "cloud.google.com/go/firestore"


func UpdateDocField(collectionRef *firestore.CollectionRef, docId string, value []firestore.Update) error {
    _, err := collectionRef.Doc(docId).Update(ctx, value)
    return err
}
