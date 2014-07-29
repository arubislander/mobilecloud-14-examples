package videoservlet

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

/**
*
* This test sends a POST request to the VideoServlet to add a new video and
* then sends a second GET request to check that the video showed up in the list
* of videos.
*
* The test launches the servlet.
*
*  To run this test, navigate to the folder containing the source code in a
*  terminal window and type
 * "go test"
*
* @author arubislander
*
*/

// By default, the test server will be running on localhost and listening to
// port 8080. If the server is running and you can't connect to it with this test,
// ensure that a firewall (e.g. Windows Firewall) isn't blocking access to it.
const TEST_URL = "http://localhost:8080/VideoServlet/video"

func s4() string {
	return fmt.Sprintf("%x", int(math.Floor(1+rand.Float64()*0x10000)))
}

func randomUUID() string {
	return s4() + s4() + "-" + s4() + "-" + s4() + "-" +
		s4() + "-" + s4() + s4() + s4()
}

/**
* This test sends a GET request with a msg parameter and
* ensures that the servlet replies with "Echo:" + msg.
*
* @throws Exception
 */

func TestVideoAddAndList(t *testing.T) {

	// Run the server
	servlet := New("localhost", 8080)
	go servlet.Run()

	// Information about the video
	// We create a random String for the title so that we can ensure
	// that the video is added after every run of this test.
	myRandomID := randomUUID()
	title := "Video - " + myRandomID
	videoUrl := "http://coursera.org/some/video-" + myRandomID
	duration := 60 * 10 * 1000

	// Create the HTTP POST request data to send the video to the server
	data := createVideoPostData(title, videoUrl, duration)

	// Use our HttpClient to send the POST request and obtain the
	// HTTP response that the server sends back.
	resp, err := http.DefaultClient.PostForm(TEST_URL, data)
	if err != nil {
		t.Error("Executing request gave an error:", err)
	}

	// Check that we got an HTTP 200 OK response code
	if resp.StatusCode != http.StatusOK {
		t.Error("Expected statuscode: %d got %d", http.StatusOK, resp.StatusCode)
	}

	// Retrieve the HTTP response body from the HTTP response
	responseBody, err := extractResponseBody(resp)
	if err != nil {
		t.Error("Executing request gave an error:", err)
	}

	// Make sure that the response is what we expect. Rather than trying to
	// keep the response message from the VideoServlet in synch with this
	// test, we simply use a public static final variable on the VideoServlet so
	// that we can refer to the message in both places and avoid the test and
	// servlet definition of the message drifting out of synch.
	if VIDEO_ADDED != responseBody {
		t.Error("Expected response:", VIDEO_ADDED, "got", responseBody)
	}

	// Now that we have posted the video to the server, we construct
	// an HTTP GET request to fetch the list of videos from the VideoServlet
	// Execute our GET request and obtain the server's HTTP response
	listResponse, err := http.DefaultClient.Get(TEST_URL)
	if err != nil {
		t.Error("Executing request gave an error:", err)
	}

	// Check that we got an HTTP 200 OK response code
	if listResponse.StatusCode != http.StatusOK {
		t.Error("Expected statuscode: %d got %d", http.StatusOK, listResponse.StatusCode)
	}

	// Extract the HTTP response body from the HTTP response
	receivedVideoListData, err := extractResponseBody(listResponse)
	if err != nil {
		t.Error("Executing request gave an error:", err)
	}

	// Construct a representation of the text that we are expecting
	// to see in the response representing our video
	expectedVideoEntry := title + " : " + videoUrl + "\n"

	// Check that our video shows up in the list by searching for the
	// expectedVideoEntry in the text of the response body
	if !strings.Contains(receivedVideoListData, expectedVideoEntry) {
		t.Error("Expected entry:", expectedVideoEntry, "not found in response ", receivedVideoListData)
	}
}

/**
 * This test sends a POST request to the VideoServlet and supplies an
 * empty String for the "name" parameter, which should cause the
 * VideoServlet to generate an error 400 Bad request response.
 *
 */

func TestMissingRequestParameter(t *testing.T) {
	// Information about the video
	// We create a random String for the title so that we can ensure
	// that the video is added after every run of this test.

	// We are going to purposely send an empty String for the title
	// in this test and ensure that the VideoServlet generates a 400
	// response code.

	var title string
	myRandomID := randomUUID()
	videoUrl := "http://coursera.org/some/video-" + myRandomID
	duration := 60 * 10 * 1000 // 10 min in milliseconds

	// Create the HTTP POST request data to send the video to the server
	data := createVideoPostData(title, videoUrl, duration)

	// Use the default HttpClient to send the POST request and obtain the
	// HTTP response that the server sends back
	response, err := http.DefaultClient.PostForm(TEST_URL, data)
	if err != nil {
		t.Error("Executing request gave an error:", err)
	}

	// The VideoServlet should generate an error 400 Bad request response
	if response.StatusCode != http.StatusBadRequest {
		t.Error("Expected statuscode: %d got %d", http.StatusBadRequest, response.StatusCode)
	}
}

/*
 * This simple method extracts the HTTP response body
 */
func extractResponseBody(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", buffer), nil
}

/*
 * This method constructs a properly formatted POST request
 * that can be sent to the VideoServlet to add a video.
 *
 * @param title
 * @param videoUrl
 * @param duration
 * @return
 */
func createVideoPostData(title string, videoUrl string, duration int) url.Values {
	data := url.Values{}
	data.Set("name", title)
	data.Set("url", videoUrl)
	data.Set("duration", fmt.Sprintf("%d", duration))
	return data
}
