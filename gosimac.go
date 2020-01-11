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
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/1995parham/gosimac/bing"
	"github.com/1995parham/gosimac/core"
	"github.com/1995parham/gosimac/unsplash"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// FetchedImages is a default number for fetching images from sources
const FetchedImages = 10

func picturesDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Errorf("user.Current: %v", err)
	}

	p := path.Join(usr.HomeDir, "Pictures", "GoSiMac")

	if _, err := os.Stat(p); err != nil {
		if err := os.Mkdir(p, 0755); err != nil {
			log.Fatalf("os.Mkdir: %v", err)
		}
	}

	return p
}

// nolint: funlen
func main() {
	p := picturesDir()

	app := cli.NewApp()
	app.Name = "GoSiMac"
	app.Usage = "Fetch the wallpaper from Bings, Wikimedia ..."
	app.Version = "3.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Parham Alvani",
			Email: "parham.alvani@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "p",
			Usage:       "A path to store the photos",
			Value:       p,
			Destination: &p,
		},
	}
	app.CommandNotFound = func(c *cli.Context, s string) {
		fmt.Printf("Invalid type is used, type %s is unknown\n", s)
	}
	app.Commands = []cli.Command{
		{
			Name:    "unsplash",
			Aliases: []string{"u"},
			Usage:   "fetches images from https://unsplash.org",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n",
					Usage: "The number of photos to return",
					Value: FetchedImages,
				},
				cli.StringFlag{
					Name:  "q",
					Usage: "Limit selection to photos matching a search term.",
					Value: "",
				},
				cli.StringFlag{
					Name:  "o",
					Usage: "Filter search results by photo orientation, Valid values are landscape, portrait, and squarish.",
					Value: "landscape",
				},
			},
			Action: func(c *cli.Context) error {
				s := &unsplash.Source{
					N:           c.Int("n"),
					Query:       c.String("q"),
					Orientation: c.String("o"),
				}
				return run(p, s, c)
			},
		},
		{
			Name:    "bing",
			Aliases: []string{"b"},
			Usage:   "fetches images from https://bing.com",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n",
					Usage: "The number of photos to return",
					Value: FetchedImages,
				},
				cli.IntFlag{
					Name:  "i",
					Usage: "Index of the first image to fetch",
					Value: 0,
				},
			},
			Action: func(c *cli.Context) error {
				s := &bing.Source{
					N:     c.Int("n"),
					Index: c.Int("i"),
				}
				return run(p, s, c)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// run runs given source on given path and waits for its results
func run(p string, s core.Source, _ *cli.Context) error {
	fmt.Println(color.CyanString(">>> Source"))
	fmt.Printf(color.CyanString("%+v\n", s))
	fmt.Println(color.CyanString(">>>"))

	fmt.Println(color.GreenString(">>> Path"))
	fmt.Printf(color.GreenString("%s\n", p))
	fmt.Println(color.GreenString(">>>"))

	a := core.NewApp(p, s)
	if err := a.Run(); err != nil {
		return err
	}

	a.Wait()

	return nil
}
