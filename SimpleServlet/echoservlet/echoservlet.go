// Package servlet implements the Simple servlet.
// This package is based on example code for the Programming Mobile Cloud Services for Android Handheld Systems 2014 MOOC
// https://github.com/juleswhite/mobilecloud-14/tree/master/examples/1-SimpleServlet

package echoservlet

import (
	"fmt"
	"net/http"
)

// EchoServlet listens on a given address and port and echos messages sent to it.
type EchoServlet struct {
	address string
	port    int
}

// New creates and returns a pointer to an EchoServlet
func New(address string, port int) *EchoServlet {
	s := new(EchoServlet)
	s.address = address
	s.port = port
	return s
}

// echo handles an HTTP request by echoing the msg in the response.
func (s *EchoServlet) echo(w http.ResponseWriter, r *http.Request) {

	// Set the content type header that is going to be returned in the
	// HTTP response so that the client will know how to display the
	// result.
	w.Header().Set("Content-Type", "text/plain")

	// Look inside of the HTTP request for either a query parameter or
	// a url encoded form parameter in the body that is named "msg"
	msg := r.FormValue("msg")

	// http://foo.bar?msg=asdf

	// Echo a response back to the client with the msg that was sent
	fmt.Fprint(w, "Echo:", msg)
}

// Run registers the echo function and starts the server.
func (s *EchoServlet) Run() {
	http.HandleFunc("/SimpleServlet/echo", s.echo)
	http.ListenAndServe(fmt.Sprintf("%s:%d", s.address, s.port), nil)
}
