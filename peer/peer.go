package peer

import (
	"encoding/base64"
	"net"
)

// Payload represent bytes for send/receive and UDP address for peer
type Payload struct {
	Bytes []byte
	Addr  *net.UDPAddr
}

// Endpoint each peer endpoint
type Endpoint struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

// Peer represent peer connection details
type Peer struct {
	ID         string       `json:"id,omitempty"`
	Username   string       `json:"username,omitempty"`
	Endpoint   Endpoint     `json:"endpoint,omitempty"`
	PublicKey  string       `json:"publicKey,omitempty"`
	PrivateKey [32]byte     `json:"-"`
	Addr       *net.UDPAddr `json:"-"`
}

// GetPublicKey get peer public key
func (p *Peer) GetPublicKey() ([32]byte, error) {
	var key [32]byte
	bs, err := base64.StdEncoding.DecodeString(p.PublicKey)
	if err != nil {
		return key, err
	}
	copy(key[:], bs)
	return key, nil
}

// SetPublicKey set peer public key
func (p *Peer) SetPublicKey(key [32]byte) {
	p.PublicKey = base64.StdEncoding.EncodeToString(key[:])
}

// Peers connection of peers
type Peers map[string]*Peer
