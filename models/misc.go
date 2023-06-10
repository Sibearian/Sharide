// Fuctions for common tasks
package models

import "encoding/json"

func RemoveUser(idx int, arr []User) []User {
    arr[idx] = arr[len(arr) - 1]
    return arr[:len(arr) - 1]
}

func FindUser(ele User, arr []User) (idx int) {
    for idx, aele := range arr {
        if ele.Userid == aele.Userid {
            return idx
        }
    }
    return -1
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
    data, err := json.Marshal(obj)
    if err != nil {
        return
    }
    err = json.Unmarshal(data, &newMap)
    return
}
