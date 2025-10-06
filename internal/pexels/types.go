package pexels

// Response structure stores pexels api response from json.
// nolint: tagliatelle
type Response struct {
	Page         int     `json:"page"`
	PerPage      int     `json:"per_page"`
	Photos       []Photo `json:"photos"`
	TotalResults int     `json:"total_results"`
	NextPage     string  `json:"next_page"`
}

// Photo represents pexels photo information.
// nolint: tagliatelle
type Photo struct {
	ID              int    `json:"id"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	URL             string `json:"url"`
	Photographer    string `json:"photographer"`
	PhotographerURL string `json:"photographer_url"`
	Alt             string `json:"alt"`
	Src             Src    `json:"src"`
}

// Src contains photo URLs in different sizes.
type Src struct {
	Original  string `json:"original"`
	Large2x   string `json:"large2x"`
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}
