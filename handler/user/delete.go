/**
* @file   : delete
* @descrip: 
* @author : ch-yk
* @create : 2018-09-05 下午1:48
* @email  : commonheart.yk@gmail.com
**/

package user

import (

	"strconv"

	"api_gateway/internal/errno"
	"api_gateway/model"
	. "api_gateway/handler"

	"github.com/gin-gonic/gin"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [delete]
func Delete(c *gin.Context) {
	/*
	根据用户传入过来的 id 进行软删除:
	 DELETE /v1/user/:id
	*/
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
