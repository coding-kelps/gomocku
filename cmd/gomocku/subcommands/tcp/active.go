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

func InitActiveCmd() *cobra.Command {
    activeCmd := cobra.Command{
        Use:   "active",
        Short: "initiate connection to the gomokurs environment manager",
        Run:    activeExecute,
    }

	activeCmd.Flags().String("address", "localhost:49912", "The listening address")

	return &activeCmd
}

func activeExecute(cmd *cobra.Command, args []string) {
	address, _ := cmd.Flags().GetString("address")

    conn, err := net.Dial("tcp", address)
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        return
    }
	defer conn.Close()

	tcp, err := adapters.NewTCPManagerInterface(conn)
	if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	ai := ai.NewRandomAI()
	coord := coordinator.NewCoordinator(tcp, ai)
	
	err = coord.Serve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}
