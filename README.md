# Go Si Mac
[![Travis](https://img.shields.io/travis/1995parham/gosimac.svg?style=flat-square)](https://travis-ci.org/1995parham/gosimac)
[![Docker Pulls](https://img.shields.io/docker/pulls/1995parham/gosimac.svg?style=flat-square)](https://hub.docker.com/r/1995parham/gosimac/)
[![GoDoc](https://godoc.org/github.com/1995parham/gosimac?status.svg)](http://godoc.org/github.com/1995parham/gosimac)
[![Codacy grade](https://img.shields.io/codacy/grade/fa84da9d770f4487bb1d6d6d74154267.svg?style=flat-square)](https://www.codacy.com/app/1995parham/gosimac?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=1995parham/gosimac&amp;utm_campaign=Badge_Grade)

## Introduction

GoSiMac downloads Bing daily wallpapers and random images from Unsplash, etc.
for you to have beautiful wallpaper on your desktop whenever you want.

Personally, I wrote this to have fun and help one of my friends who is not among us right now. :disappointed:

This module is highly customizable and new sources can easily add just by implementing source interface.

## Docker

```
docker run --rm -v $(pwd)/.:/root/Pictures/Bing 1995parham/gosimac
```

## Read More

### How to change your OS X wallpaper from terminal :smile:

Form OS X Mavericks, all settings stored in sqlite library in following
path:

`~/Library/Application Support/Dock/desktoppicture.db`

### How to create wallpaper sideshow in ubuntu :smile:

First put configuration of GoSiMac into gnome desktop configuration folder
with following command:
```sh
dst=$HOME/.local/share/gnome-background-properties
test -d "$dst" || mkdir -p "$dst" && cp gnome/gosimac-config.xml "$dst"
```
