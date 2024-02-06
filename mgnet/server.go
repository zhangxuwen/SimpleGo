package mgnet

import (
	"errors"
	"fmt"
	"marcoGo/mgiface"
	"net"
)

// implementation interface of IServer
// define Server
type Server struct {

	// server name
	Name string

	// the IP version bound to the server
	IPVersion string

	// the IP listened by the server
	IP string

	// the port listened by the server
	Port string
}

// the function for start server
func (s *Server) Start() {

	// the flag of the server is starting
	fmt.Printf("[marcoGo start] at IP: %s, Port: %s\n", s.IP, s.Port)

	// get the address of TCP
	addr, err := net.ResolveTCPAddr(s.IPVersion, s.IP+":"+s.Port)
	if err != nil {
		fmt.Printf("[marcoGo error] for %s", err)
		return
	}

	// listen address of server
	linstener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Printf("[marcoGo error] for %s", err)
		return
	}

	var cid uint32
	cid = 0

	// wait for connection
	for {

		// continue if connect
		connect, err := linstener.AcceptTCP()
		if err != nil {
			fmt.Printf("[marcoGo error] for %s", err)
			continue
		}

		// example
		dealConn := NewConnection(connect, cid, func(conn *net.TCPConn, data []byte, cnt int) error {

			fmt.Println(data[:cnt])
			if _, err := conn.Write(data[:cnt]); err != nil {
				return errors.New("call back wall")
			}
			return nil
		})
		cid++
		go dealConn.Start()

	}
}

// stop server
func (s *Server) Stop() {

	// do something

}

// run server
func (s *Server) Serve() {

	// start server
	s.Start()

	// do something

	// block
	select {}
}

/*
init the server moudule
*/
func NewServer(name string) mgiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "8999",
	}
	return s
}
