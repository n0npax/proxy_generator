package cmd

import (
	"fmt"
	"os"
	"io"
	"github.com/n0npax/proxy_generator/parser"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.Flags().String("endpoint", "", "endpoint to db")
	RootCmd.Flags().String("output", "", "delault stdut, if path will redirect output to given file")
	RootCmd.Flags().String("template-file","./templates/nginx.conf.tmpl", "")
}
// RootCmd is main cobra command
var RootCmd = &cobra.Command{
	Use: "generates nginx config with redirections based on given input",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := cmd.Flag("endpoint").Value.String()
		outputFlag := cmd.Flag("output").Value.String()
		templateFile := cmd.Flag("template-file").Value.String()
		
		var redirections []parser.NginxRedirection

		// output stream decision
		var output io.Writer
		if outputFlag == "" {
			output = os.Stdout
		}		else{
			var err error
			output, err = os.Create(outputFlag)
			if err != nil {
				panic(err)
			}
		}
		// input data
		if endpoint != "" {
			fmt.Printf("Not implemented Yet")
			os.Exit(1)		
		}
		
		if len(args) %2 != 0 {
			fmt.Printf("Need args in pairs")
			os.Exit(1)
		} else {
			for i:=0; i < len(args); i+=2 {
				redirections = append(redirections,
					parser.NginxRedirection{
						InternalURL: args[i], ExternalURL: args[i+1]})
			}
		}

		parser.GenerateConfig(redirections, output, templateFile)
	},
}