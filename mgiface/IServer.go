package mgiface

// define server interface
type IServer interface {

	// start server
	Start()

	// stop server
	Stop()

	// run server
	Serve()
}
