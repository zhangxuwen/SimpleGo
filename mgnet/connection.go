package mgnet

import (
	"fmt"
	"marcoGo/mgiface"
	"net"
)

/*
connection moudule
*/
type Connection struct {

	// the socket TCP of connection
	Connection *net.TCPConn

	// connection ID
	ConnectionID uint32

	// connection status
	isClosed bool

	// API of handle transaction
	handleAPI mgiface.HandleFunc

	// channel as flag show if stop
	ExitChan chan bool
}

// init moudule connection
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI mgiface.HandleFunc) *Connection {

	c := &Connection{
		Connection:   conn,
		ConnectionID: connID,
		handleAPI:    callbackAPI,
		isClosed:     false,
		ExitChan:     make(chan bool, 1),
	}

	return c
}

// read transaction of connection
func (c *Connection) StartReader() {

	fmt.Println("[marcoGo read is ready]")

	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Connection.Read(buf)
		if err != nil {
			fmt.Println("[marcoGo recv buffer error]", err)
			continue
		}

		// run the handle API
		if err := c.handleAPI(c.Connection, buf, cnt); err != nil {

			fmt.Println("[marcoGo handle error] connectionID = ", c.ConnectionID)
			break

		}

	}

}

// start connection
func (c *Connection) Start() {

	fmt.Println("[marcoGo connection start] connectionID = ", c.ConnectionID)

	// start transaction of the connection

	// start read transaction

	// start write transaction
}

// close connection
func (c *Connection) Stop() {

	fmt.Println("[marcoGo connection stop] ConnectionID = ", c.ConnectionID)

	// if the connection has closed
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	// close the connection socket
	c.Connection.Close()

	// close the channel
	close(c.ExitChan)
}

// get the socket of connection
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Connection
}

// get the ID of connection
func (c *Connection) GetConnectionID() uint32 {
	return c.ConnectionID
}

// get the client's TCP IP and port of connection
func (c *Connection) GetRemoteAddress() net.Addr {
	return c.Connection.RemoteAddr()
}

// send data to client
func (c *Connection) Send(data []byte) error {
	return nil
}
