package service

import (
	"413zk/bff/basic/config"
	"413zk/bff/handler/request"
	__ "413zk/proto"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var form request.Login
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"code": 400,
		})
		return
	}
	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name:     form.Name,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  r.Msg,
		"code": r.Code,
	})
	return
}
func Register(c *gin.Context) {
	var form request.Register
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "参数错误",
			"code": 400,
		})
		return
	}
	r, err := config.UserClient.Login(c, &__.LoginReq{
		Name:     form.Name,
		Password: form.Password,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  r.Msg,
		"code": r.Code,
	})
	return
}
