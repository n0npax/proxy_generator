package main

import (
	"github.com/n0npax/proxy_generator/cmd"
//	"github.com/spf13/cobra"
	"os"
	"fmt"
)



func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

