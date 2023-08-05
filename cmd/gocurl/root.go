package gocurl

import (
	"fmt"
	"os"

	"github.com/mrsimi/gocurl/pkg/gocurl"
	"github.com/spf13/cobra"
)

var url string
var rootCmd = &cobra.Command{
	Use: "gocurl [request type]",
	Short: "gocurl - a go version of the curl tool",
	Long: `gocurl is a go version of the curl command-line tool.
it is command-line tool for getting or sending data including files using URL syntax`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		switch requestType := args[0]; requestType{
		case "GET", "get":
			result := gocurl.GetRequest(url)
			fmt.Println(result)
		default: 
			fmt.Println("Request Type not supported yet")
		}
	},
}

func Execute(){
	//setting all the flags
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "resource url")

	if err := rootCmd.Execute(); err != nil{
		fmt.Fprintln(os.Stderr, "Whopps. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}