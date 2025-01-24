package tcp

import (
	"fmt"
	"io"
	"net"

	"github.com/coding-kelps/gomocku/pkg/domain/listener/models"
)

func (tcp *Tcp) Listen(ch chan<-models.ManagerCommand) error {
	for {
        conn, err := tcp.listener.Accept()
        if err != nil {
			fmt.Println("Error:", err)
			continue
        }

		go handleConnection(conn, ch)
	}
}

func handleConnection(conn net.Conn, ch chan<-models.ManagerCommand) {
	defer conn.Close()

	handlers := map[byte]func(conn net.Conn)(models.ManagerCommand, error){
		0x01: StartHandler,
		0x02: TurnHandler,
		0x03: BeginHandler,
		0x04: BoardHandler,
		0x05: BoardTurnHandler,
		0x06: BoardDoneHandler,
		0x07: InfoHandler,	
		0x08: EndHandler,
		0x09: AboutHandler,
	}

	for {
		var cmdID [1]byte
		if _, err := io.ReadFull(conn, cmdID[:]); err != nil {
			if err == io.EOF {
				fmt.Printf("Client %s disconnected\n", conn.RemoteAddr())
			} else {
				fmt.Printf("Error reading command ID from %s: %v", conn.RemoteAddr(), err)
			}

			return
		}

		handler, ok := handlers[cmdID[0]]
		if !ok {
            fmt.Printf("Unknown command ID %d received, closing connection", cmdID[0])

            return
		}

		cmd, err := handler(conn)
        if err != nil {
            fmt.Printf("Error handling command %d: %v", cmdID[0], err)

			// TO DO: Implement error notification to client
            continue
        }

		ch <- cmd
	}
}
