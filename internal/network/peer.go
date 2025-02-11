package network

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Peer struct {
	PeerServer *Server
	PeerClient *Client
	PeerList   map[string]string
	Mu         sync.RWMutex
}

func (pr *Peer) PeerStartServer(addr, port string) {
	pr.PeerServer = &Server{Addr: addr, Port: port}
	pr.PeerServer.startServer()
}

func (pr *Peer) PeerLoop() {
	fmt.Println("Entering the Program")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed reading input", err)
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

	}

}
