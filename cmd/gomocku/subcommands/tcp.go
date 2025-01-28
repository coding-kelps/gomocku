package subcommands

import (
	"github.com/coding-kelps/gomocku/cmd/gomocku/subcommands/tcp"

	"github.com/spf13/cobra"
)

func InitTcpCmd() *cobra.Command {
	tcpCmd := cobra.Command{
		Use:	"tcp",
		Short:	"run the gomocku mock AI using TCP",
	}

	tcpCmd.AddCommand(tcp.InitActiveCmd())
	tcpCmd.AddCommand(tcp.InitPassiveCmd())

	return &tcpCmd
}
