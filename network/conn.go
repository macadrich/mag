package network

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"net"
	"strconv"
)

// Conn is a connection to a remote peer
type Conn interface {
	Send([]byte)
	Protocol() string
	GetAddr() net.Addr
	GetEncryptKey() ([32]byte, error)
	SetEncryptKey([32]byte)
}

// Payload represent bytes for send/receive and UDP address for peer
type Payload struct {
	Bytes []byte
	Addr  *net.UDPAddr
}

// PeerConns map peer connections
type PeerConns map[string]PeerConn

// PeerConn remote peer connection
type PeerConn struct {
	send   chan *Payload
	addr   *net.UDPAddr
	secret string
}

// Send p2p send data
func (c *PeerConn) Send(b []byte) {
	c.send <- &Payload{Bytes: b, Addr: c.addr}
}

// Protocol return protocol name
func (c *PeerConn) Protocol() string {
	return "UDP"
}

// GetAddr peer address
func (c *PeerConn) GetAddr() net.Addr {
	return c.addr
}

// GetEncryptKey get peer encrypted key
func (c *PeerConn) GetEncryptKey() ([32]byte, error) {
	return convertEncryptKey(c.secret)
}

// SetEncryptKey set peer encrypted key
func (c *PeerConn) SetEncryptKey(secret [32]byte) {
	c.secret = base64.StdEncoding.EncodeToString(secret[:])
}

// NewUDPConn initialize new UDPConn
func NewUDPConn(send chan *Payload, addr *net.UDPAddr) *PeerConn {
	return &PeerConn{
		send: send, // payload
		addr: addr, // address
	}
}

func convertEncryptKey(keyText string) ([32]byte, error) {
	// ensure secret has been set
	var keyInByte [32]byte
	if keyText == "" {
		return keyInByte, errors.New("key has not been set")
	}

	// decode to byte slice
	bs, err := base64.StdEncoding.DecodeString(keyText)
	if err != nil {
		return keyInByte, errors.New("could not decode key")
	}

	// copy byte slice into byte array
	copy(keyInByte[:], bs)
	return keyInByte, nil
}

// GenPort generate random port
func GenPort() string {
	return ":" + strconv.Itoa(rand.Intn(65535-10000)+10000)
}
