package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/kar1mov-u/P2P-file_sharing/internal/network"
)

func main() {

	peer := network.Peer{}
	var wg sync.WaitGroup
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Address: ")
	addr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed reading input", err)
	}
	addr = strings.TrimSpace(addr)

	fmt.Print("Enter Port: ")
	port, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed reading input", err)
	}
	port = strings.TrimSpace(port)

	wg.Add(1)
	go func() {
		defer wg.Done()
		peer.PeerStartServer(addr, port)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		peer.PeerLoop()
	}()

	wg.Wait()

}
