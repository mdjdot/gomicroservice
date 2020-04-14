package main

import (
	"errors"
	"gomicroservice/protorpc/message"
	"net/http"
	"net/rpc"
	"time"
)

// OrderService 订单服务
type OrderService struct{}

// GetOrderInfo 获取订单信息
func (o *OrderService) GetOrderInfo(request message.OrderRequest, response *message.OrderInfo) error {
	orderMap := map[string]message.OrderInfo{
		"201907300001": message.OrderInfo{OrderID: "201907300001", OrderName: "衣服", OrderStatus: "已付款"},
		"201907310001": message.OrderInfo{OrderID: "201907310001", OrderName: "零食", OrderStatus: "已付款"},
		"201907310002": message.OrderInfo{OrderID: "201907310002", OrderName: "食品", OrderStatus: "未付款"},
	}
	current := time.Now().Unix()
	if request.Timestamp > current {
		return errors.New("请求信息错误")
	}
	*response = orderMap[request.OrderID]
	return nil
}

func main() {
	orderService := &OrderService{}
	rpc.Register(orderService)
	rpc.HandleHTTP()
	panic(http.ListenAndServe(":8080", nil))
}
