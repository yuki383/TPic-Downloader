// Package pictures provide routines to create picture data.
package pictures

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Picture is representing a picture data.
type Picture struct {
	Name string
	Data []byte
}

// PictureService has dependencies functions of to make a Picture.
type PictureService struct {
	httpGet func(url string) (*http.Response, error)
	timeNow func() time.Time
}

func newService() *PictureService {
	s := &PictureService{
		httpGet: http.Get,
		timeNow: time.Now,
	}

	return s
}

var defaultService = newService()

// New returns picture data of fetch by url.
func New(name string, url string) (Picture, error) {
	if name == "" || url == "" {
		err := errors.New("invalid arguments.\nexpect: pictures.New(name string, url string)")
		return Picture{}, err
	}

	return defaultService.new(name, url)
}

func (s *PictureService) new(name string, url string) (Picture, error) {
	t := s.timeNow().Local()
	n := s.makeName(name, t)

	data, err := s.fetch(url)
	if err != nil {
		return Picture{}, err
	}

	p := Picture{
		Name: n,
		Data: data,
	}

	return p, nil
}

func (s *PictureService) fetch(url string) ([]byte, error) {
	res, err := s.httpGet(url)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (s *PictureService) makeName(name string, date time.Time) string {
	min := date.Unix()
	path := fmt.Sprintf("%d_%s", min, name)

	return path
}
