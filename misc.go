// Fuctions for common tasks
package main

func removeUser(idx int, arr []User) []User {
    arr[idx] = arr[len(arr) - 1]
    return arr[:len(arr) - 1]
}

func findUser(ele User, arr []User) (idx int) {
    for idx, aele := range arr {
        if ele.Userid == aele.Userid {
            return idx
        }
    }
    return -1
}
