package main

import (
	"fmt"
	"gopenapi/cmd"
	"os"
)

func main() {
	err := cmd.InitApp().Run(os.Args)
	if err != nil {
		fmt.Printf("[ERROR]: %s", err)
		os.Exit(1)
	}
}
