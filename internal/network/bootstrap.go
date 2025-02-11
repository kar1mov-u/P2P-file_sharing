package network

import (
	"crypto/sha1"
	"encoding/hex"
)

type BootstrapServer struct {
	Server *Server
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
