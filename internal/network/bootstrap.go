package network

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

type BootstrapServer struct {
	Host  string
	Port  string
	Peers map[string]string
	mu    sync.RWMutex
}

func (bs *BootstrapServer) StartBootstrapServer() {
	bs.Peers = make(map[string]string)
	fmt.Printf("Starting Bootstrap server on: %s:%s\n", bs.Host, bs.Port)
	listener, err := net.Listen("tcp", bs.Host+":"+bs.Port)
	if err != nil {
		fmt.Println("Failed to start Bootstrap server", err)
		os.Exit(1)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed on establishing connection with client", err)
		}

		go bs.ServeClient(conn)

	}
}

func (bs *BootstrapServer) ServeClient(conn net.Conn) {

	defer func() {
		fmt.Println("Closing connection with", conn.RemoteAddr().String())
		bs.RemovePeer(GeneratePeerId(conn.RemoteAddr().String()))
		conn.Close()
	}()

	fmt.Println("Established connection with " + conn.RemoteAddr().String())
	//Add new connected peer to the list of known peers
	peerId := GeneratePeerId(conn.RemoteAddr().String())
	bs.AddPeer(peerId, conn.RemoteAddr().String())

	buff := make([]byte, 4096)
	for {
		n, err := conn.Read(buff)

		if err != nil {
			if err == io.EOF {
				fmt.Println("client disconnected :" + conn.RemoteAddr().String())
				// Remove CLient from the network
				bs.RemovePeer(peerId)
				break
			} else {
				fmt.Println("Error while reading from client", err)
				continue
			}
		}

		message := strings.TrimSpace(string(buff[:n]))
		fmt.Println("Message of client: " + string(buff[:n]))

		if message == "peer-discover" {
			resp := bs.GetPeers()
			respJson, err := json.Marshal(resp)
			if err != nil {
				fmt.Println("Failed to Serialize to Json", err)
				conn.Write([]byte("{}"))
				return
			}

			fmt.Println(respJson, "peers")
			_, err = conn.Write([]byte(respJson))
			if err != nil {
				fmt.Println("failed to write the Response", err)
			}
		}
	}
}

func (bs *BootstrapServer) AddPeer(peerId, peer string) {
	bs.mu.Lock()
	bs.Peers[peerId] = peer
	bs.mu.Unlock()
}

func (bs *BootstrapServer) GetPeers() map[string]string {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	peersCopy := make(map[string]string)
	for k, peer := range bs.Peers {
		peersCopy[k] = peer
	}
	return peersCopy
}

func (bs *BootstrapServer) RemovePeer(peerId string) {
	bs.mu.Lock()
	delete(bs.Peers, peerId)
	bs.mu.Unlock()
}

func GeneratePeerId(peerInfo string) string {
	hash := sha1.Sum([]byte(peerInfo))
	return hex.EncodeToString(hash[:])[:10]
}
