package unsplash

// Image represents unsplash image information
type Image struct {
	ID   string `json:"id"`
	URLs struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Description string
}
