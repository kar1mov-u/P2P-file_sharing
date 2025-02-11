package network

import (
	"fmt"
	"net"

	pb "github.com/kar1mov-u/P2P-file_sharing/internal/network/proto"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	Addr string
	Port string
}

func (cl *Client) PeerDiscoverReq(addr, port string) *pb.PeerDiscoverResponse {
	resp := &pb.PeerDiscoverResponse{}
	req := pb.ClientRequest{Type: "peer-discover"}
	conn, err := net.Dial("tcp", addr+":"+port)
	if err != nil {
		fmt.Println("Failed to make request to ", addr, port, err)
		return resp

	}

	data, err := proto.Marshal(&req)
	if err != nil {
		fmt.Println("Failed Serializng Request", err)
		return resp
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Failed on writing to the wire", err)
	}

	buff := make([]byte, 4096)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("Failed on Reading response", err)
		return resp
	}

	err = proto.Unmarshal(buff[:n], resp)
	if err != nil {
		fmt.Println("Failed serialize response", err)
	}

	return resp

}
