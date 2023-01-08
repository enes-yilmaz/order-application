package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func PrintBanner(projectName, port string) {
	_, err := exec.LookPath("figlet")

	if err != nil {
		fmt.Println("project: " + projectName)
	} else {
		cmd := exec.Command("figlet", projectName)
		cmd.Stdin = strings.NewReader("and old falcon")
		var out bytes.Buffer
		cmd.Stdout = &out
		_ = cmd.Run()
		fmt.Println(out.String())
	}

	fmt.Printf("running on port %s\n", port)
}

func GetGoEnv() string {
	e, ok := os.LookupEnv("ENV")
	if !ok {
		e = defaultGoEnv
	}
	return e
}

func GetSwagHostEnv() string {
	e, ok := os.LookupEnv("SWAG_URL")
	if !ok {
		e = defaultSwagUrl
	}
	return e
}

const defaultSwagUrl = "localhost:4000/customer-api"
const defaultGoEnv = "prod"
