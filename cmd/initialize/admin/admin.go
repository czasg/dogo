package admin

import (
	"github.com/spf13/cobra"
	"proj/cmd/initialize/admin/account"
	"proj/cmd/initialize/admin/model"
)

var (
	AdminCmd = &cobra.Command{
		Use:   "admin",
		Short: "init admin",
		Long:  "init admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			panic("not found init command")
		},
	}
)

func init() {
	AdminCmd.AddCommand(
		model.ModelCmd,
		account.AccountCmd,
	)
}
