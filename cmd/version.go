package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string
var revision string

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information and quit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Commitizen-go version %s, build revision %s\n", version, revision)
	},
}
