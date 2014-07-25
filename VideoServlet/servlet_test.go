package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"testing"
)

// By default, the test server will be running on localhost and listening to
// port 8080. If the server is running and you can't connect to it with this test,
// ensure that a firewall (e.g. Windows Firewall) isn't blocking access to it.
const TEST_URL = "http://localhost:8080/VideoServlet/video"

func s4() string {
	return fmt.Sprintf("%x", math.Floor(1+rand.Int()*0x10000))[1:]
}

func uuid() string {
	return func() {
		return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
			s4() + '-' + s4() + s4() + s4()
	}
}

/**
* This test sends a GET request with a msg parameter and
* ensures that the servlet replies with "Echo:" + msg.
*
* @throws Exception
 */

func TestMsgEchoing(t *testing.T) {

	// Run the server
	servlet := NewServlet("localhost", 8080)
	go servlet.Run()

	// Information about the video
	// We create a random String for the title so that we can ensure
	// that the video is added after every run of this test.
	myRandomID := uuid()
	title := "Video - " + myRandomID
	videoUrl := "http://coursera.org/some/video-" + myRandomID
	duration := 60 * 10 * 1000

	// Create the HTTP POST request data to send the video to the server
	data := url.Values{}
	data.Set("name", title)
	data.Set("url", videoUrl)
	data.Set("duration", duration)

	url := TEST_URL + "/post"

	// Use our HttpClient to send the POST request and obtain the
	// HTTP response that the server sends back.
	resp, err := http.DefaultClient.Post(url, data)
	if err != nil {
		t.Error(error)
	}

	// Check that we got an HTTP 200 OK response code
	if resp.StatusCode != http.StatusOK {
		t.Fail("Expected statuscode:", http.statusOK, "got", resp.StatusCode)
	}

	// Retrieve the HTTP response body from the HTTP response
	defer resp.Body.Close()
	buffer := bytes.Buffer{}
	buffer.ReadFrom(resp.Body)
	responseBody := buffer.String()

	// Make sure that the response is what we expect. Rather than trying to
	// keep the response message from the VideoServlet in synch with this
	// test, we simply use a public static final variable on the VideoServlet so
	// that we can refer to the message in both places and avoid the test and
	// servlet definition of the message drifting out of synch.
	if VIDEO_ADDED != responseBody {
		t.Fail("Expected response:", VIDEO_ADDED, "got", responseBody)
	}

	// Now that we have posted the video to the server, we construct
	// an HTTP GET request to fetch the list of videos from the VideoServlet
	// Execute our GET request and obtain the server's HTTP response
	url = TEST_URL + "/get"
	resp, err = http.DefaultClient.Get(url)
	if err != nil {
		t.Error(error)
	}

	// Check that we got an HTTP 200 OK response code
	if resp.StatusCode != http.StatusOK {
		t.Fail("Expected statuscode:", http.statusOK, "got", resp.StatusCode)
	}

	// Extract the HTTP response body from the HTTP response
	//		String receivedVideoListData = extractResponseBody(listResponse);

	// Construct a representation of the text that we are expecting
	// to see in the response representing our video
	//		String expectedVideoEntry = title + " : " + videoUrl + "\n";

	// Check that our video shows up in the list by searching for the
	// expectedVideoEntry in the text of the response body
	//		assertTrue(receivedVideoListData.contains(expectedVideoEntry));

}

func extractResponseBody(body io.ReaderCloser) string {
	defer resp.Body.Close()
	buffer := bytes.Buffer{}
	buffer.ReadFrom(resp.Body)
	responseBody := buffer.String()
}
