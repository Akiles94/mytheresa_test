package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configFile = ""

//Root command
var RootCmd = &cobra.Command{
	Use:   "info",
	Short: "MyTheresa Promotions Test",
	Long:  `REST API for MyTheresa challenge`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("use: ./ serve")
	},
}

// Execute commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
