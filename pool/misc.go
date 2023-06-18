// Fuctions for common tasks
package pool

import (
	"ShaRide/models"
	"encoding/json"
	"math"

	"github.com/pierrre/geohash"
)

const earthRadius float64 = 6371000.0

func RemoveUser(idx int, arr []models.UserSlice) []models.UserSlice {
    arr[idx] = arr[len(arr) - 1]
    return arr[:len(arr) - 1]
}

func FindUser(user models.UserSlice, arr []models.UserSlice) (idx int) {
    for idx, member := range arr {
        if user == member {
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

func Get300mBox(hash string) []string {
    neighbors, _ := geohash.GetNeighbors(hash)
    return []string{
        hash,
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

func get600mBox(hash string) (res []string) {
    var hashs map[string]bool
    hashs = make(map[string]bool)

    for _, hash1 := range Get300mBox(hash) {
        for _, hash2 := range Get300mBox(hash1) {
            hashs[hash2] = true
        }
    }

    for k, _ := range hashs {
        res = append(res, k)
    }
    return 
}

type BoundingBox struct {
	MinLatitude  float64
	MaxLatitude  float64
	MinLongitude float64
	MaxLongitude float64
}

func GetBoundingBox(latitude, longitude, distance float64) BoundingBox {
	earthRadius := 6371000.0

	distanceRadians := distance / earthRadius

	latitudeRad := latitude * (math.Pi / 180)
	longitudeRad := longitude * (math.Pi / 180)

	minLatitude := latitudeRad - distanceRadians
	maxLatitude := latitudeRad + distanceRadians

	deltaLongitude := math.Asin(math.Sin(distanceRadians) / math.Cos(latitudeRad))
	minLongitude := longitudeRad - deltaLongitude
	maxLongitude := longitudeRad + deltaLongitude

	maxLatitude = maxLatitude * (180 / math.Pi)
	minLongitude = minLongitude * (180 / math.Pi)
	maxLongitude = maxLongitude * (180 / math.Pi)

	return BoundingBox{
		MinLatitude:  minLatitude,
		MaxLatitude:  maxLatitude,
		MinLongitude: minLongitude,
		MaxLongitude: maxLongitude,
	}
}
