package unsplash

// Image represents unsplash image information.
type Image struct {
	ID   string `json:"id"`
	URLs struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Location struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"location"`
	Description string `json:"description"`
}
