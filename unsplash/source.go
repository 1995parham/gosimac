package unsplash

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Source is source implmentation for unsplash image service
type Source struct {
	response []Image
	N        int
}

// Init initiates source and return number of avaiable images
func (s *Source) Init() (int, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/photos/random?count=%d", "https://api.unsplash.com/", s.N), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Accept-Version", "v1")
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %s", "4c483af1b27cf8d55fc29504bc48e3755e47eb7a3dd3a320e92b23fc4e5aa1b8"))

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return 0, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&s.response); err != nil {
		return 0, fmt.Errorf("decoding json: %v", err)
	}

	return len(s.response), err
}

// Name returns source name
func (s *Source) Name() string {
	return "unsplash"
}

// Fetch fetches given index from source
func (s *Source) Fetch(index int) (string, io.ReadCloser, error) {
	image := s.response[index]

	logrus.Infof("Getting %s\n", image.ID)

	resp, err := http.Get(image.URLs.Full)
	if err != nil {
		return "", nil, err
	}

	logrus.Infof("%s was gotten\n", image.ID)

	return fmt.Sprintf("%s.jpg", image.ID), resp.Body, nil
}
