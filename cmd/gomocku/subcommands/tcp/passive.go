package tcp

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/coding-kelps/gomocku/pkg/adapters"
	"github.com/coding-kelps/gomocku/pkg/domain/ai"
	"github.com/coding-kelps/gomocku/pkg/domain/listener"
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
		fmt.Printf("error: %v\n", err)
	}

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
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
}
