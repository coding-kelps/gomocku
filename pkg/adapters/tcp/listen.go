package tcp

import (
	"fmt"
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

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		fmt.Printf("Received: %s\n", buffer[:n])
		ch <- models.UnknownCommand{}
	}
}
