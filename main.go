package main

import (
	"fmt"
	"github.com/n0npax/proxy_generator/cmd"
	"os"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
