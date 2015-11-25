/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 25-11-2015
 * |
 * | File Name:     types.go
 * +===============================================
 */
package bing

type BingResponse struct {
	Images  []Image `json:"images"`
	Tooltip Tooltip `json:"tooltip"`
}

type Image struct {
	StartDate     string `json:"startdate"`
	FullStartDate string `json:"fullstartdate"`
	EndDate       string `json:"enddate"`
	URL           string `json:"url"`
	URLBase       string `json:"urlbase"`
	Copyright     string `json:"copyright"`
	CopyrightLink string `json:"copyrightlink"`
	Wallpaper     bool `json:"wp"`
	Hash          string `json:"hsh"`
	Drk           int `json:"drk"`
	Top           int `json:"top"`
	Bot           int `json:"bot"`
	HS            []HS `json:"hs"`
	Msg           []string `json:"msg"`

}

type HS struct {
	Description string `json:"desc"`
	Link        string `json:"link"`
}

type Tooltip struct {
	Loading        string `json:"loading"`
	Previous       string `json:"previous"`
	Next           string `json:"next"`
	WallpaperSave  string `json:"walls"`
	WallpaperError string `json:"walle"`
}