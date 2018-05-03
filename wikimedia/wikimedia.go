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
	"net/http"

	log "github.com/sirupsen/logrus"
)

// GetWikimediaPOTD function gets Picture of The Day from wikimedia and stores it in `parh`.
func GetWikimediaPOTD(path string) error {
	// Create HTTP GET request
	resp, err := http.Get(
		fmt.Sprintf("https://commons.wikimedia.org/w/api.php?action=query&Generator=images&titles=Template:Potd/2014-12-02&prop=imageinfo&iiprop=url&format=json&formatVersion=2"),
	)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Errorf("(io.Closer).Close: %v", err)
		}
	}()

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
