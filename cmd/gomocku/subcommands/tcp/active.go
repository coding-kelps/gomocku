package tcp

import (
    "fmt"
    "net"

	"github.com/spf13/cobra"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/listener"
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
        fmt.Printf("error: %v\n", err)
        return
    }
	defer conn.Close()

	tcp := adapters.NewTCP(conn)

	ai := ai.NewRandomAI()
	listener := listener.NewListener(tcp, ai)
	
	err = listener.Listen()
	if err != nil {
		fmt.Printf("%e\n", err)
	}
}
