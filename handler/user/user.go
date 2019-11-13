/**
* @file   : user.go
* @descrip: 封装的是请求结构体和响应消息结构体(对应response的data类型), 这里不写明的，一般对应 handler.go 的通用消息格式
* @author : ch-yk
* @create : 2018-09-05 下午1:04
* @email  : commonheart.yk@gmail.com
**/

package user

import "api_gateway/model"
//模块: /v1/user

//对应  POST /v1/user
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//SendResponse(c, nil, rsp)，这里就是需要返回的 rsp
type CreateResponse struct {
	Username string `json:"username"`
}


//对应 GET /user
// 一般传递 {"offset": 0, "limit": 20} 即可
type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

//响应头返回的格式比较特殊: []*model.UserInfo, 结构体指针(地址)的数组
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}


