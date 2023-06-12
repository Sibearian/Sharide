// Types of the api requests
package app

import "ShaRide/models"

type PoolReq struct {
    ReqUser models.User `json:"user"`
    PoolId  string      `json:"pool_id"`
}

type Resp struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data"`
}
