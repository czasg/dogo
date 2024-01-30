package admin

import (
	"github.com/spf13/cobra"
)

var (
	ApiCmd = &cobra.Command{
		Use:   "api",
		Short: "init admin api",
		Long:  "init admin api",
		RunE: func(cmd *cobra.Command, args []string) error {
			panic("not found init command")
		},
	}
)
