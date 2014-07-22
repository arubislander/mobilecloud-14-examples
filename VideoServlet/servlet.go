// VideoServlet project main.go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Servlet struct {
	address string
	port    int
	videos  []*Video
}

func NewServlet(address string, port int) *Servlet {
	s := new(Servlet)
	s.address = address
	s.port = port

	// An in-memory list that the servlet uses to store the
	// videos that are sent to it by clients
	s.videos = make([]*Video, 0)

	return s
}

const VIDEO_ADDED = "Video added."

/**
* This method processes all of the HTTP GET requests routed to the
* servlet by the web container. This method loops through the lists
* of videos that have been sent to it and generates a plain/text
* list of the videos that is sent back to the client.
*
 */
func (s *Servlet) getHandler(w http.ResponseWriter, r *http.Request) {

	// Make sure and set the content-type header so that the client
	// can properly (and securely!) display the content that you send
	// back
	w.Header().Set("Content-Type", "text/plain")

	// Loop through all of the stored videos and print them out
	// for the client to see.
	for _, v := range s.videos {
		fmt.Fprintln(w, v.Name(), " : ", v.Url())
	}
}

/**
* This method handles all HTTP POST requests that are routed to the
* servlet by the web container.
*
* Sending a post to the servlet with 'name', 'duration', and 'url'
* parameters causes a new video to be created and added to the list of
* videos.
*
* If the client fails to send one of these parameters, the servlet generates
* an HTTP error 400 (Bad request) response indicating that a required request
* parameter was missing.
 */
func (s *Servlet) postHandler(w http.ResponseWriter, r *http.Request) {
	// First, extract the HTTP request parameters that we are expecting
	// from either the URL query string or the url encoded form body
	name := r.FormValue("name")
	url := r.FormValue("url")
	durationStr := r.FormValue("duration")

	// Check that the duration parameter provided by the client
	// is actually a number
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		// The client sent us a duration value that wasn't a number!
		duration = -1
	}

	// Make sure and set the content-type header so that the client knows
	// how to interpret the data that gets sent back
	w.Header().Set("Content-Type", "text/plain")

	// Now, the servlet has to look at each request parameter that the
	// client was expected to provide and make sure that it isn't null,
	// empty, etc.
	if len(strings.Trim(name, " ")) < 1 || len(strings.Trim(url, " ")) < 10 || len(strings.Trim(durationStr, " ")) < 1 || duration <= 0 {

		// If the parameters pass our basic validation, we need to
		// send an HTTP 400 Bad Request to the client and give it
		// a hint as to what it got wrong.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing ['name','duration','url'].")
	} else {
		// It looks like the client provided all of the data that
		// we need, use that data to construct a new Video object
		v := NewVideo(name, url, int64(duration))

		// Add the video to our in-memory list of videos
		s.videos = append(s.videos, v)

		// Let the client know that we successfully added the video
		// by writing a message into the HTTP response body
		fmt.Fprint(w, VIDEO_ADDED)
	}
}

func (s *Servlet) Run() {
	http.HandleFunc("/VideoServlet/video/get", s.getHandler)
	http.HandleFunc("/VideoServlet/video/post", s.postHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", s.address, s.port), nil)
}
