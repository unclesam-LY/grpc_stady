package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func (s Server) Add(req Req, res *Res) error {
	res.Num = req.Num1 + req.Num2
	fmt.Println("请求来了", req)
	return nil
}

func main() {
	// 注册rpc服务
	err := rpc.Register(new(Server))
	if err != nil {
		fmt.Println(err)
		return
	}
	rpc.HandleHTTP()

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = http.Serve(listen, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
