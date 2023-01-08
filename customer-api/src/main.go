package main

import (
	"CustomerAPI/src/cmd/customer"
	"CustomerAPI/src/pkg/utils"
)

func main() {
	
	env := utils.GetGoEnv()
	customer.Execute(env)

}
