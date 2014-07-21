// SimpleServlet project main.go
package main

import (
	"fmt"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {

	// Set the content type header that is going to be returned in the
	// HTTP response so that the client will know how to display the
	// result.
	w.Header().Set("Content-Type", "text/plain")

	// Look inside of the HTTP request for either a query parameter or
	// a url encoded form parameter in the body that is named "msg"
	msg := r.URL.Query().Get("msg")

	// http://foo.bar?msg=asdf

	// Echo a response back to the client with the msg that was sent
	fmt.Fprint(w, "Echo:", msg)

}

func main() {
	http.HandleFunc("/SimpleServlet/echo", echo)
	http.ListenAndServe("localhost:8080", nil)
}
