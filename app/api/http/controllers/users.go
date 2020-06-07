package controllers

import (
	"net/http"

	"github.com/spadesk1991/skip-micro/app/common"

	"github.com/gin-gonic/gin"
	"github.com/spadesk1991/skip-micro/app/services/impl"
	"github.com/spadesk1991/skip-micro/app/services/pb"
)

func Registry(c *gin.Context) {
	userService := new(impl.UsersService)
	req := new(pb.RegRequest)
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err_code": -1, "msg": err.Error()})
		return
	}
	err = common.Validate(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err_code": -1, "msg": err.Error()})
		return
	}
	res := new(pb.RegResponse)
	err = userService.Registry(c, req, res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err_code": -1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"err_code": 0, "data": res.Data})
}
