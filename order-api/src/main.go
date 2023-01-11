package main

import (
	"OrderAPI/src/cmd/order"
	"OrderAPI/src/pkg/utils"
)

func main() {
	env := utils.GetGoEnv()
	order.Execute(env)
}
