/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-12-2017
 * |
 * | File Name:     wikimedia.go
 * +===============================================
 */

package wikimedia

import (
	"encoding/json"
	"fmt"

	"github.com/franela/goreq"
)

// GetWikimediaPOTD function gets Picture of The Day from wikimedia and stores it in `parh`.
func GetWikimediaPOTD(path string) error {
	// Create HTTP GET request
	resp, err := goreq.Request{
		Uri: "https://commons.wikimedia.org/w/api.php",
		QueryString: Request{
			Action:        "query",
			Generator:     "images",
			Titles:        "Template:Potd/2014-12-02",
			Prop:          "imageinfo",
			Iiprop:        "url",
			Format:        "json",
			FormatVersion: 2,
		},
		UserAgent: "GoSiMac",
	}.Do()
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var wikimediaResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&wikimediaResp); err != nil {
		return fmt.Errorf("decoding json: %v", err)
	}
	wikimediaResp = wikimediaResp["query"].(map[string]interface{})
	wikimediaResp = wikimediaResp["pages"].([]interface{})[0].(map[string]interface{})
	potd := wikimediaResp["imageinfo"].([]interface{})[0].(map[string]interface{})["url"].(string)

	fmt.Println(potd)

	return nil
}
