package rpc

import (
	"fmt"
)

func GetRpcResult(request int) int {
	fmt.Println("GetRpcResult execute")
	return request * 100
}
