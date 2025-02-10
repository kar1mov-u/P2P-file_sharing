package main

import "github.com/kar1mov-u/P2P-file_sharing/internal/network"

func main() {

	bserver := network.BootstrapServer{Host: "0.0.0.0", Port: "8000"}
	bserver.StartBootstrapServer()
}
