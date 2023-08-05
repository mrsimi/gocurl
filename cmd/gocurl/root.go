package gocurl

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gocurl ",
	Short: "gocurl - a go version of the curl tool",
	Long: `gocurl is a go version of the curl command-line tool.
it is command-line tool for getting or sending data including files using URL syntax`,

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		fmt.Fprintln(os.Stderr, "Whopps. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}