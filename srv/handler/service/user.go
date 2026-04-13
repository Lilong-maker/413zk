package service

import (
	__ "413zk/proto"
	"413zk/srv/basic/config"
	"413zk/srv/handler/model"
	"context"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Login(_ context.Context, in *__.LoginReq) (*__.LoginResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.LoginResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	if user.Password != in.Password {
		return &__.LoginResp{
			Msg:  "密码错误",
			Code: 400,
		}, nil
	}
	return &__.LoginResp{
		Msg:  "登录成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Register(_ context.Context, in *__.RegisterReq) (*__.RegisterResp, error) {

	var user model.User
	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.RegisterResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	m := model.User{
		Name:     in.Name,
		Password: in.Password,
	}
	err = m.RedisterAdd(config.DB)
	if err != nil {
		return &__.RegisterResp{
			Msg:  "注册失败",
			Code: 400,
		}, nil
	}

	return &__.RegisterResp{
		Msg:  "注册成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderAdd(_ context.Context, in *__.OrderAddReq) (*__.OrderAddResp, error) {

	var order model.Order
	err := order.FindOrder(config.DB, in.Name)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "商品不存在",
			Code: 400,
		}, nil
	}
	m := model.Order{
		Name:    order.Name,
		Address: order.Address,
		Num:     order.Num,
	}
	err = m.OrderAdd(config.DB)
	if err != nil {
		return &__.OrderAddResp{
			Msg:  "商品添加失败",
			Code: 400,
		}, nil
	}

	return &__.OrderAddResp{
		Msg:  "商品添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) GwcAdd(_ context.Context, in *__.GwcAddReq) (*__.GwcAddResp, error) {

	var gwc model.Gwc
	err := gwc.Findgwc(config.DB, in.Name)
	if err != nil {
		return &__.GwcAddResp{
			Msg:  "购物车不存在",
			Code: 400,
		}, nil
	}
	m := model.Gwc{
		Name: gwc.Name,
		Num:  gwc.Num,
	}
	err = m.GwcAdd(config.DB)
	if err != nil {
		return &__.GwcAddResp{
			Msg:  "添加失败",
			Code: 400,
		}, nil
	}

	return &__.GwcAddResp{
		Msg:  "添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderList(_ context.Context, in *__.OrderListReq) (*__.OrderListResp, error) {

	var list model.Order
	err := list.OrderList(config.DB, in.Id)
	if err != nil {
		return &__.OrderListResp{
			Msg:  "列表查询失败",
			Code: 400,
		}, nil
	}

	return &__.OrderListResp{
		Msg:  "列表查询成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) OrderLists(_ context.Context, in *__.OrderListReq) (*__.OrderListResp, error) {

	var list model.Order
	err := list.OrderList(config.DB, in.Id)
	if err != nil {
		return &__.OrderListResp{
			Msg:  "库存扣减失败",
			Code: 400,
		}, nil
	}

	return &__.OrderListResp{
		Msg:  "库存扣减成功",
		Code: 200,
	}, nil
}
