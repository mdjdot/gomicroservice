package main

import (
	"fmt"
	"gomicroservice/protorpc/message"
	"net/rpc"
	"time"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	timeStamp := time.Now().Unix()
	request := message.OrderRequest{
		OrderID:   "201907310002",
		Timestamp: timeStamp,
	}
	var resp message.OrderInfo
	err = client.Call("OrderService.GetOrderInfo", request, &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", resp)
}
