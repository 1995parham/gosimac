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
	"os/user"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 1, "number of wallpapers that you want from Bing :)")

	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		glog.Errorf("OS.User: %v", err)
	}

	bing.GetBingDesktop(fmt.Sprintf("%s/Pictures/Bing", usr.HomeDir), false, num-1, num)
}
