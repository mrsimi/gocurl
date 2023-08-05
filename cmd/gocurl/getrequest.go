package gocurl

import (
	"fmt"
	
	"github.com/mrsimi/gocurl/pkg/gocurl"
	"github.com/spf13/cobra"
)

var url string
var getCmd = &cobra.Command{
	Use: "get",
	Aliases: []string{"g"},
	Short: "recieves a resource from a url",

	//make an httpRequest and display the result to the console
	Run: func(cmd *cobra.Command, args []string) {
		result := gocurl.GetRequest(url)
		fmt.Printf(result)
	},
}

func init(){
	getCmd.Flags().StringVarP(&url, "url", "u", "", "resource url")
	rootCmd.AddCommand(getCmd)
}