package tcp

import (
	"github.com/spf13/cobra"
)

func InitPassiveCmd() *cobra.Command {
	passiveCmd := cobra.Command{
		Use:   "passive",
		Short: "wait for the gomokurs environment manager to initiate connection",
		Run: 	passiveExecute,
	}

	passiveCmd.Flags().String("address", "localhost:49912", "The listening address")

	return &passiveCmd
}

func passiveExecute(cmd *cobra.Command, args []string) {
	_, _ = cmd.Flags().GetString("address")
}
