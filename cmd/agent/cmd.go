package agent

import (
	"github.com/spf13/cobra"
)

var (
	AgentCmd = &cobra.Command{
		Use:   "agent",
		Short: "agent server",
		Long:  `start a agent server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)
