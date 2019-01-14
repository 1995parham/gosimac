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
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
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

	app := cli.NewApp()
	app.Name = "GoSiMac"
	app.Usage = "Fetch the wallpaper from Bings, Wikimedia ..."
	app.Version = "3.0.0"
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
					Value: 10,
				},
				cli.StringFlag{
					Name:  "q",
					Usage: "Limit selection to photos matching a search term.",
					Value: "",
				},
			},
			Action: func(c *cli.Context) error {
				s := &unsplash.Source{
					N:     c.Int("n"),
					Query: c.String("q"),
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
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				s := &bing.Source{
					N: c.Int("n"),
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
func run(p string, s core.Source, c *cli.Context) error {
	fmt.Println(">>> Source")
	fmt.Printf("%+v\n", s)
	fmt.Println(">>>")

	fmt.Println(">>> Path")
	fmt.Printf("%s\n", p)
	fmt.Println(">>>")

	a := core.NewApp(p, s)
	if err := a.Run(); err != nil {
		return err
	}

	a.Wait()
	return nil

}
