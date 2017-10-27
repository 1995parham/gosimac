# Go Si Mac
[![Travis](https://img.shields.io/travis/1995parham/gosimac.svg?style=flat-square)](https://travis-ci.org/1995parham/gosimac)

## Introduction

Automatically downloads bing daily wallpapers for having beautiful wallpapers on your desktop.
Personally I wrote this for having fun and helping one of my friends :P.

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
