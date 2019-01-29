# Go Si Mac
[![Travis](https://img.shields.io/travis/com/1995parham/gosimac.svg?style=flat-square)](https://travis-ci.com/1995parham/gosimac)
[![Docker Pulls](https://img.shields.io/docker/pulls/1995parham/gosimac.svg?style=flat-square)](https://hub.docker.com/r/1995parham/gosimac/)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/fa84da9d770f4487bb1d6d6d74154267)](https://www.codacy.com/app/1995parham/gosimac?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=1995parham/gosimac&amp;utm_campaign=Badge_Grade)
[![Go Report](https://goreportcard.com/badge/github.com/1995parham/gosimac?style=flat-square)](https://goreportcard.com/report/github.com/1995parham/gosimac)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/1995parham/gosimac)
[![GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

## Introduction

GoSiMac downloads Bing daily wallpapers and random images from Unsplash, etc.
for you to have beautiful wallpaper on your desktop whenever you want.

Personally, I wrote this to have fun and help one of my friends who is not among us right now. :disappointed:

This module is highly customizable and new sources can easily add just by implementing source interface.

## Docker

```
docker run --rm -v $(pwd)/.:/root/Pictures/GoSiMac 1995parham/gosimac ...
```

## Read More

### How to change your OS X wallpaper from terminal :smile:

Form OS X Mavericks, all settings stored in sqlite library in following
path:

`~/Library/Application Support/Dock/desktoppicture.db`

### How to create wallpaper sideshow in ubuntu :smile:

For just the **basic** automatic wallpaper changing feature, you don’t need to install any software.
Just launch the pre-installed *Shotwell* photo manager, choose the pictures you need (you may need to import them first), then go to `Files -> Set as Desktop Slideshow`.
