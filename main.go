package main

import "marcoGo/mgnet"

func main() {
	server := mgnet.NewServer("0.0.0.0")
	server.Start()
}
