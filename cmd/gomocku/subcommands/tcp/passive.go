package tcp

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/coordinator"
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
	address, _ := cmd.Flags().GetString("address")

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }

		tcp, err := adapters.NewTCPManagerInterface(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)

			return
		}

		ai := ai.NewRandomAI()
		coord := coordinator.NewCoordinator(tcp, ai)
	
		err = coord.Serve()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		conn.Close()
		return
	}
}
