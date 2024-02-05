package mgnet

import (
	"fmt"
	"marcoGo/mgiface"
	"net"
)

// IServer 的接口实现，定义一个Server的服务器模块
type Server struct {

	// 服务器的名称
	Name string

	// 服务器绑定的ip版本
	IPVersion string

	// 服务器监听的IP
	IP string

	// 服务器监听的端口
	Port string
}

// 启动服务器
func (s *Server) Start() {

	// 执行启动服务器的标志
	fmt.Printf("[marcoGo start] at IP: %s, Port: %s\n", s.IP, s.Port)

	// 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, s.IP+":"+s.Port)
	if err != nil {
		fmt.Printf("[marcoGo error] for %s", err)
		return
	}

	// 监听服务器的地址
	linstener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Printf("[marcoGo error] for %s", err)
		return
	}

	// 阻塞的等待客户端连接，处理客户端链接业务
	for {
		// 如果有客户端链接过来，阻塞会返回
		connect, err := linstener.Accept()
		if err != nil {
			fmt.Printf("[marcoGo error] for %s", err)
			continue
		}

		// 业务
		go func() {
			for {
				buf := make([]byte, 512)
				client, err := connect.Read(buf)
				if err != nil {
					fmt.Printf("[marcoGo error] for %s", err)
					continue
				}

				// 回显
				_, err = connect.Write(buf[:client])
				if err != nil {
					fmt.Printf("[marcoGo error] for %s", err)
					continue
				}
			}
		}()
	}
}

// 停止服务器
func (s *Server) Stop() {

	// 停止或回收服务器的资源、状态或者一些已经开辟的链接信息

}

// 运行服务器
func (s *Server) Serve() {

	// 启动服务器
	s.Start()

	// 额外业务

	//阻塞状态
	select {}
}

/*
初始化Server模块的方法
*/
func newServer(name string) mgiface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      "8999",
	}
	return s
}
