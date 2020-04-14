package order

import (
	"context"
	"errors"
	"fmt"
	gorder "gomicroservice/grpcrpc/order"
	"io"
	"strconv"
)

// OrderServiceImpl .
type OrderServiceImpl struct {
	// 	// OrderServiceServer is the server API for OrderService service.
	// type OrderServiceServer interface {
	// 	GetOrderInfo(context.Context, *OrderRequest) (*OrderInfo, error)
	// 	GetOrderInfoStream(*OrderRequest, OrderService_GetOrderInfoStreamServer) error
	// 	GetOrderInfoClientStream(OrderService_GetOrderInfoClientStreamServer) error
	// 	GetOrderInfo2Stream(OrderService_GetOrderInfo2StreamServer) error
	// }
}

// GetOrderInfo .
func (os *OrderServiceImpl) GetOrderInfo(ctx context.Context, request *gorder.OrderRequest) (*gorder.OrderInfo, error) {
	if request.OrderID == "123" {
		response := &gorder.OrderInfo{
			OrderID:     "123",
			OrderName:   "abc124",
			OrderStatus: "ok",
		}
		return response, nil
	}

	return nil, errors.New("not found")
}

// GetOrderInfoStream .
func (os *OrderServiceImpl) GetOrderInfoStream(request *gorder.OrderRequest, stream gorder.OrderService_GetOrderInfoStreamServer) error {
	orderMap := map[int]gorder.OrderInfo{
		1: gorder.OrderInfo{OrderID: "1", OrderName: "11", OrderStatus: "one"},
		2: gorder.OrderInfo{OrderID: "2", OrderName: "22", OrderStatus: "two"},
		3: gorder.OrderInfo{OrderID: "3", OrderName: "33", OrderStatus: "three"},
	}
	for _, order := range orderMap {
		stream.Send(&order)
	}
	return nil
}

// GetOrderInfoClientStream .
func (os *OrderServiceImpl) GetOrderInfoClientStream(stream gorder.OrderService_GetOrderInfoClientStreamServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			result := &gorder.OrderInfo{OrderStatus: "done"}
			return stream.SendAndClose(result)
		}
		if err != nil {
			return err
		}
		fmt.Println(request)
	}
}

// GetOrderInfo2Stream .
func (os *OrderServiceImpl) GetOrderInfo2Stream(stream gorder.OrderService_GetOrderInfo2StreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			panic(err)
		}
		orderMap := map[int]gorder.OrderInfo{
			1: gorder.OrderInfo{OrderID: "1", OrderName: "11", OrderStatus: "one"},
			2: gorder.OrderInfo{OrderID: "2", OrderName: "22", OrderStatus: "two"},
			3: gorder.OrderInfo{OrderID: "3", OrderName: "33", OrderStatus: "three"},
		}
		orderID, _ := strconv.Atoi(req.OrderID)
		result := orderMap[orderID]
		err = stream.Send(&result)
		if err == io.EOF {
			return err
		}
		if err != nil {
			panic(err)
		}
	}
	return nil
}
