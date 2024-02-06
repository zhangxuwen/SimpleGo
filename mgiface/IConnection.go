package mgiface

import "net"

// define connection interface
type IConnection interface {

	// start connection
	Start()

	// close connection
	Stop()

	// get the socket of connection
	GetTCPConnection() *net.TCPConn

	// get the ID of connection
	GetConnectionID() uint32

	// get the client's TCP IP and port of connection
	GetRemoteAddress() net.Addr

	// send data to client
	Send(data []byte) error
}

// define function for handling the transaltion
type HandleFunc func(*net.TCPConn, []byte, int) error
