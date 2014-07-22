// VideoServlet project main.go
package main

func main() {
	s := NewServlet("localhost", 8080)
	s.Run()
}
