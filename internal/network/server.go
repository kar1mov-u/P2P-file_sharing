package network

import (
	"fmt"
	"io"
	"net"
	"os"

	pb "github.com/kar1mov-u/P2P-file_sharing/internal/network/proto"
	"google.golang.org/protobuf/proto"
)

type Server struct {
	Addr string
	Port string
}

func (srv *Server) startServer() {
	listener, err := net.Listen("tcp", srv.Addr+":"+srv.Port)
	fmt.Printf("Started listening on %s:%s \n", srv.Addr, srv.Port)

	if err != nil {
		fmt.Println("Failed to start server", err)
		os.Exit(1)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Failed on connecting to client", err)
			continue
		}

		go srv.handleReq(conn)

	}

}

func (srv *Server) handleReq(conn net.Conn) {
	defer func() {
		fmt.Println("Disconnecting from ", conn.RemoteAddr().String())
		conn.Close()
	}()

	fmt.Println("Connected to ", conn.RemoteAddr().String())

	buff := make([]byte, 4096)

	for {

		n, err := conn.Read(buff)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Peer disconnected ", conn.RemoteAddr().String())
				break
			} else {
				fmt.Println("Error on reading input", err)
				continue
			}
		}

		var req pb.ClientRequest
		err = proto.Unmarshal(buff[:n], &req)
		if err != nil {
			fmt.Println("Failed to serialize peer message", err)
			continue
		}

		fmt.Println(req.Type)

	}

}
