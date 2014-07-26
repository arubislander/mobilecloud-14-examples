package echoservlet

import (
	"bytes"
	"net/http"
	"testing"
)

// By default, the test server will be running on localhost and listening to
// port 8080. If the server is running and you can't connect to it with this test,
// ensure that a firewall (e.g. Windows Firewall) isn't blocking access to it.
const TEST_URL = "http://localhost:8080/SimpleServlet/echo"

/**
 * This test sends a GET request with a msg parameter and
 * ensures that the servlet replies with "Echo:" + msg.
 */

func TestMsgEchoing(t *testing.T) {
	servlet := New("localhost", 8080)
	go servlet.Run()
	// The message to send to the EchoServlet
	msg := "1234"

	// Append our message to the URL so that the
	// EchoServlet will send the message back to us
	url := TEST_URL + "?msg=" + msg

	// Send an HTTP GET request to the EchoServer and
	// convert the response body to a String
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(resp.Body)
	content := buffer.String()

	// Ensure that the body of the HTTP response met our
	// expectations (e.g., it was "Echo:" + msg)
	if content != "Echo:"+msg {
		t.Error("Expected Echo:", msg, "got", content)
	}

}
