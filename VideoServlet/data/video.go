// Package data implements the Video data type.
// This package is based on example code for the Programming Mobile Cloud Services for Android Handheld Systems 2014 MOOC
// https://github.com/juleswhite/mobilecloud-14/tree/master/examples/2-VideoServlet

package data

type Video struct {
	name     string
	url      string
	duration int64
}

func NewVideo(name string, url string, duration int64) *Video {
	v := new(Video)
	v.name = name
	v.url = url
	v.duration = duration
	return v
}

func (v *Video) Name() string {
	return v.name
}

func (v *Video) setName(name string) {
	v.name = name
}

func (v *Video) Url() string {
	return v.url
}

func (v *Video) setUrl(url string) {
	v.url = url
}

func (v *Video) getDuration() int64 {
	return v.duration
}

func (v *Video) setDuration(duration int64) {
	v.duration = duration
}
