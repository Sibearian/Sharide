// Fuctions for common tasks
package pool

import (
	"ShaRide/models"
	"encoding/json"

	"github.com/pierrre/geohash"
)

func RemoveUser(idx int, arr []models.User) []models.User {
    arr[idx] = arr[len(arr) - 1]
    return arr[:len(arr) - 1]
}

func FindUser(ele models.User, arr []models.User) (idx int) {
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

func getNeighborSlice(hash string) []string {
    neighbors, _ := geohash.GetNeighbors(hash)
    return []string{
        neighbors.East,
        neighbors.West,
        neighbors.North,
        neighbors.South,
        neighbors.NorthEast,
        neighbors.NorthWest,
        neighbors.SouthEast,
        neighbors.SouthWest,
    }
}
