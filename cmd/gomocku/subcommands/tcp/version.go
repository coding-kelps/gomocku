package tcp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/coding-kelps/gomocku/pkg/adapters/tcp"
)

func InitVersionCmd() *cobra.Command {
	versionCmd := cobra.Command{
		Use:   "version",
		Short: "return the protocol version of the TCP manager interface",
		Run: 	versionExecute,
	}

	return &versionCmd
}

func versionExecute(cmd *cobra.Command, args []string) {
	fmt.Fprintf(os.Stderr, "%s\n", tcp.ProtocolVersion)
}
