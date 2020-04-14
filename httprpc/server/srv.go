package main

import (
	"math"
	"net"
	"net/http"
	"net/rpc"
)

// MathUtil 定义结构体
type MathUtil struct {
}

// CalcArea 计算面积方法
func (m *MathUtil) CalcArea(redium float64, result *float64) error {
	*result = math.Pi * redium * redium
	return nil
	// return errors.New("err")
}

func main() {
	mathUtil := &MathUtil{}
	err := rpc.Register(mathUtil)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	http.Serve(listen, nil)
}
