package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"micro-example/define"
	userProto "micro-example/service/user/proto"
	"net/http"
)

var userCli userProto.UserService

type UserParam struct{
	Username string
	Password string
}

func init() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()

	cli := service.Client()

	userCli = userProto.NewUserService("micro.service.user", cli)
}

func SignUpHandler(c *gin.Context) {
	var userParam UserParam
	err := c.ShouldBind(&userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": define.CodeInvalidParameter,
			"msg": "invalid parameters",
		})
		return
	}
	log.Printf("user param is %v\n", userParam)
	respSignup, err := userCli.Signup(context.TODO(), &userProto.ReqSignup{
		Username: userParam.Username,
		Password: userParam.Password,
	})

	if err != nil {
		log.Printf("sign up failed: %v\n", err)
		c.JSON(http.StatusBadGateway, gin.H{
			"code": define.CodeSystemError,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": respSignup.Code,
		"msg": respSignup.Message,
	})
}

func LoginHandler(c *gin.Context)  {
	var userParam UserParam
	err := c.ShouldBind(&userParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": define.CodeInvalidParameter,
			"msg": "invalid parameters",
		})
		return
	}
	log.Printf("user param is %v\n", userParam)
	respLogin, err := userCli.Login(context.TODO(), &userProto.ReqLogin{
		Username: userParam.Username,
		Password: userParam.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code": define.CodeSystemError,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": respLogin.Code,
		"msg":  respLogin.Message,
	})

}
