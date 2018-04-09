package cmd

import (
	"fmt"
	"github.com/n0npax/proxy_generator/dbreader"
	"github.com/n0npax/proxy_generator/parser"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func init() {
	RootCmd.Flags().String("db-endpoint", "", "in example sqlite://${PATH}")
	RootCmd.Flags().String("output", "", "default stdut, if path will redirect output to given file")
	RootCmd.Flags().String("template-file", "./templates/nginx.conf.tmpl", "")
}

// RootCmd is main cobra command
var RootCmd = &cobra.Command{
	Long: `Command line tool to generating nginx.conf based on given template and set of redirections
Can use command line arguments or database`,
	Short: "Command line tool to generating nginx.conf based on given template and set of redirections",
	Use: "proxy_generator example http://example.com cnn https://cnn.com",
	Run: func(cmd *cobra.Command, args []string) {
		dbEndpoint := cmd.Flag("db-endpoint").Value.String()
		outputFlag := cmd.Flag("output").Value.String()
		templateFile := cmd.Flag("template-file").Value.String()

		var redirections []parser.NginxRedirection

		// output stream decision
		var output io.Writer
		if outputFlag == "" {
			output = os.Stdout
		} else {
			var err error
			output, err = os.Create(outputFlag)
			if err != nil {
				panic(err)
			}
		}
		// input data
		if dbEndpoint != "" {
			redirections = dbreader.ReadNginxRedirection(dbEndpoint)
		}
		if len(args)%2 != 0 {
			fmt.Println("Need args in pairs")
			os.Exit(1)
		} else {
			for i := 0; i < len(args); i += 2 {
				redirections = append(redirections,
					parser.NginxRedirection{
						InternalURL: args[i], ExternalURL: args[i+1]})
			}
		}
		if len(redirections) == 0 {
			fmt.Println("No redirrections found. Will not generate anything file")
			os.Exit(1)
		}
		parser.GenerateConfig(redirections, output, templateFile)
	},
}
