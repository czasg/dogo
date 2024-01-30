package initialize

import (
	"github.com/spf13/cobra"
	"proj/cmd/initialize/admin"
)

var (
	InitCmd = &cobra.Command{
		Use:   "init",
		Short: "init app",
		Long:  "init app by command",
		RunE: func(cmd *cobra.Command, args []string) error {
			panic("not found init command")
		},
	}
)

func init() {
	InitCmd.AddCommand(admin.ApiCmd)
}
