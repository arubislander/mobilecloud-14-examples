// Package main implements main entry point for the Simple servlet.
// This package is based on example code for the Programming Mobile Cloud Services for Android Handheld Systems 2014 MOOC
// https://github.com/juleswhite/mobilecloud-14/tree/master/examples/1-SimpleServlet
package main

import "github.com/arubislander/mobilecloud-14-examples/SimpleServlet/echoservlet"

func main() {
	s := echoservlet.New("localhost", 8080)
	s.Run()
}
