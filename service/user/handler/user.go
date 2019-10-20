package handler

import (
	"context"
	"micro-example/define"
	proto "micro-example/service/user/proto"
)

type User struct{}

func (u *User) Signup(ctx context.Context, req *proto.ReqSignup, resp *proto.RespSignup) error {
	resp.Code = define.CodeSuccess
	resp.Message = "success"
	return nil
}


func (u *User) Login(ctx context.Context, req *proto.ReqLogin, resp *proto.RespLogin) error {
	if req.Username == "admin" && req.Password == "admin" {
		resp.Code = define.CodeSuccess
		resp.Message = "success"
		return nil
	}
	resp.Code = define.CodeInvalidUsernameOrPassword
	resp.Message = "invalid username or password"
	return nil
}