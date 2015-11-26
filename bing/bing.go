/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     bing.go
 * +===============================================
 */
package bing

import (
	"encoding/json"
	"fmt"
	"github.com/1995parham/gosimac/gosimac"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func GetBingDesktop(path string, change bool) error {
	resp, err := http.Get("http://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-US")
	if err != nil {
		glog.Errorf("Net.HTTP: %v\n", err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("IO.IOUtil: %v\n", err)
	}
	var bing_resp BingResponse
	json.Unmarshal(body, &bing_resp)

	for _, image := range bing_resp.Images {
		resp, err = http.Get(fmt.Sprintf("http://www.bing.com/%s", image.URL))
		if err != nil {
			glog.Errorf("Net.HTTP: %v\n", err)
			return err
		}

		defer resp.Body.Close()

		dest_file, err := os.Create(fmt.Sprintf("%s/%s.jpg", path, image.FullStartDate))
		if err != nil {
			glog.Errorf("OS: %v\n", err)
			return err
		}

		defer dest_file.Close()

		io.Copy(dest_file, resp.Body)

		if change {
			err := gosimac.ChangeDesktopBackground(fmt.Sprintf("%s/%s.jpg", path, image.FullStartDate))
			if err != nil {
				glog.Errorf("GoSiMac: %v", err)
				return err
			}
			exec.Command("killall", "Dock")

		}
	}
	return nil
}
