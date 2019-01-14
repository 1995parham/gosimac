package unsplash

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	resty "gopkg.in/resty.v1"
)

// Source is source implmentation for unsplash image service
type Source struct {
	response []Image
	N        int
	Query    string
}

// Init initiates source and return number of avaiable images
func (s *Source) Init() (int, error) {
	resp, err := resty.New().
		SetHeader("Accept-Version", "v1").
		SetHeader("Authorization", fmt.Sprintf("Client-ID %s", "4c483af1b27cf8d55fc29504bc48e3755e47eb7a3dd3a320e92b23fc4e5aa1b8")).
		SetHostURL("https://api.unsplash.com").
		R().
		SetResult(&s.response).
		SetQueryParam("count", strconv.Itoa(s.N)).
		SetQueryParam("orientation", "landscape").
		SetQueryParam("query", s.Query).
		Get("/photos/random")
	if err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("Invalid response: %s", resp.Status())
	}

	return len(s.response), nil
}

// Name returns source name
func (s *Source) Name() string {
	return "unsplash"
}

// Fetch fetches given index from source
func (s *Source) Fetch(index int) (string, io.ReadCloser, error) {
	image := s.response[index]

	logrus.Infof("Getting %s (%s)", image.ID, image.Description)

	resp, err := http.Get(image.URLs.Full)
	if err != nil {
		return "", nil, err
	}

	logrus.Infof("%s was gotten\n", image.ID)

	return fmt.Sprintf("%s.jpg", image.ID), resp.Body, nil
}
