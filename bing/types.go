package bing

// Response structure stores bing api respone from json.
type Response struct {
	Images  []Image `json:"images"`
	Tooltip Tooltip `json:"tooltip"`
}

// Request structure stores bing api request to json.
type Request struct {
	Format string `url:"format"`
	Index  int    `url:"idx"`
	Number int    `url:"n"`
	Mkt    string `url:"mkt"`
}

// Image structure stores bing image information.
type Image struct {
	StartDate     string   `json:"startdate"`
	FullStartDate string   `json:"fullstartdate"`
	EndDate       string   `json:"enddate"`
	URL           string   `json:"url"`
	URLBase       string   `json:"urlbase"`
	Copyright     string   `json:"copyright"`
	CopyrightLink string   `json:"copyrightlink"`
	Wallpaper     bool     `json:"wp"`
	Hash          string   `json:"hsh"`
	Drk           int      `json:"drk"`
	Top           int      `json:"top"`
	Bot           int      `json:"bot"`
	HS            []HS     `json:"hs"`
	Msg           []string `json:"msg"`
}

// HS structure ...
type HS struct {
	Description string `json:"desc"`
	Link        string `json:"link"`
}

// Tooltip structure ...
type Tooltip struct {
	Loading        string `json:"loading"`
	Previous       string `json:"previous"`
	Next           string `json:"next"`
	WallpaperSave  string `json:"walls"`
	WallpaperError string `json:"walle"`
}
