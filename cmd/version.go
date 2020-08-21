package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var revision string

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information and quit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Commitizen-go version 1.0.0, build revision %s\n", revision)
	},
}
