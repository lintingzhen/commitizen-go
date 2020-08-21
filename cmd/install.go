package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/lintingzhen/commitizen-go/git"
	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install this tool to git-core as git-cz",
	Run: func(cmd *cobra.Command, args []string) {
		appFilePath, _ := exec.LookPath(os.Args[0])
		if path, err := git.InstallSubCmd(appFilePath, "cz"); err != nil {
			fmt.Printf("Install commitizen failed, err=%v\n", err)
		} else {
			fmt.Printf("Install commitizen to %s\n", path)
		}
	},
}
