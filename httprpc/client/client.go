package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	req := float64(3.45)
	var resp float64
	err = client.Call("MathUtil.CalcArea", req, &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	syncCall := client.Go("MathUtil.CalcArea", req, &resp, nil)
	replayDone := <-syncCall.Done
	fmt.Println(replayDone)
	fmt.Println(syncCall)
}
