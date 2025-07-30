package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}

type Res struct {
	Num int
}

func main() {
	req := Req{
		Num1: 1,
		Num2: 2,
	}

	client, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	var res Res
	err = client.Call("Server.Add", req, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
