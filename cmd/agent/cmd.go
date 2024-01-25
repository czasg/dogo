package agent

import (
	"github.com/spf13/cobra"
	"proj/cmd"
)

var (
	agentCmd = &cobra.Command{
		Use:   "agent",
		Short: "agent server",
		Long:  `start a agent server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	cmd.App.AddCommand(agentCmd)
}
