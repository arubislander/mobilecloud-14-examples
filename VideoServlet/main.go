// Package main implements the main entry point for the Video Servlet
// This package is based on example code for the Programming Mobile Cloud Services for Android Handheld Systems 2014 MOOC
// https://github.com/juleswhite/mobilecloud-14/tree/master/examples/2-VideoServlet
package main

import "github.com/arubislander/mobilecloud-14-examples/VideoServlet/videoservlet"

func main() {
	s := videoservlet.New("localhost", 8080)
	s.Run()
}
