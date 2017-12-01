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
	"os"
	"os/user"
	"path"

	"github.com/1995parham/gosimac/bing"
	"github.com/1995parham/gosimac/wikimedia"
	"github.com/golang/glog"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 1, "Number of wallpapers that you want")

	var idx int
	flag.IntVar(&idx, "i", 0, "Base index of wallpapers that you want")

	var t string
	flag.StringVar(&t, "type", "bing", "Wallpaper service: bing wikimedia")

	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		glog.Errorf("OS.User: %v", err)
	}

	var p string
	p = path.Join(usr.HomeDir, "Pictures", "Bing")

	if _, err := os.Stat(p); err != nil {
		os.Mkdir(p, 0755)
	}

	switch t {
	case "bing":
		bing.GetBingDesktop(p, idx, num)
	case "wikimedia":
		wikimedia.GetWikimediaPOTD(p)
	}
}
