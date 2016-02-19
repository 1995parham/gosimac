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
	"flag"
	"fmt"
	"github.com/1995parham/gosimac/bing"
	"github.com/golang/glog"
	"os"
	"os/user"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 1, "Number of wallpapers that you want from Bing :)")

	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		glog.Errorf("OS.User: %v", err)
	}

	var path string
	path = fmt.Sprintf("%s/Pictures/Bing", usr.HomeDir)

	if _, err := os.Stat(path); err != nil {
		os.Mkdir(path, 0755)
	}

	bing.GetBingDesktop(path, 0, num)
}
