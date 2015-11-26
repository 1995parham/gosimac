/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 25-11-2015
 * |
 * | File Name:     gosimac.go
 * +===============================================
 */
package main

import (
	"./bing"
	"os/user"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		glog.Errorf("OS.User: %v", err)
	}

	bing.GetBingDesktop(fmt.Sprintf("%s/Pictures/Bing", usr.HomeDir), true)
}
