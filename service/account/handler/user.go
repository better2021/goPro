package handler

import "context"

type User struct {

}

// Signup:处理用户注册请求
func Signup(context.Context,req*proto.ReqSignuo,res*proto.RespSignup) error{
	username := req.Username
	passwd := req.Password

	// 参数简单效验
	if len(username) < 3 || len(passwd) < 5 {
		res.Code = common.StatusParaminvalid
		res.Message = "注册参数无效"
		return nil
	}
}