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
	"github.com/1995parham/gosimac/core"
	"github.com/1995parham/gosimac/unsplash"
	log "github.com/sirupsen/logrus"
)

func main() {
	var num int
	flag.IntVar(&num, "n", 1, "Number of wallpapers that you want")

	var idx int
	flag.IntVar(&idx, "i", 0, "Base index of wallpapers that you want")

	var t string
	flag.StringVar(&t, "type", "bing", "Wallpaper service: bing unsplash")

	flag.Parse()

	usr, err := user.Current()
	if err != nil {
		log.Errorf("user.Current: %v", err)
	}

	var p string
	p = path.Join(usr.HomeDir, "Pictures", "GoSiMac")

	if _, err := os.Stat(p); err != nil {
		if err := os.Mkdir(p, 0755); err != nil {
			log.Fatalf("os.Mkdir: %v", err)
		}
	}

	var s core.Source

	switch t {
	case "bing":
		s = &bing.Source{
			Idx: idx,
			N:   num,
		}
	case "unsplash":
		s = &unsplash.Source{
			N: num,
		}
	}

	a := core.NewApp(p, s)
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

	a.Wait()
}
