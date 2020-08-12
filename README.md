# Go Si Mac
[![Drone (cloud)](https://img.shields.io/drone/build/1995parham/gosimac.svg?style=flat-square&logo=drone)](https://cloud.drone.io/1995parham/gosimac)
[![Docker Pulls](https://img.shields.io/docker/pulls/1995parham/gosimac.svg?style=flat-square&logo=docker)](https://hub.docker.com/r/1995parham/gosimac/)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/1995parham/gosimac/latest?style=flat-square&logo=docker)
[![GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/1995parham/gosimac?logo=github&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date/1995parham/gosimac?logo=github&style=flat-square)

## Introduction

GoSiMac downloads Bing daily wallpapers, random images from Unsplash, and etc.
for you to have a beautiful wallpaper on your desktop whenever you want.
Following command download unsplash Tehran images for you:

```sh
gosimac u -q Tehran
```

*gosimac* stores images in `$HOME/Pictures/GoSiMac`.

Personally, I wrote this to have fun and help one of my friends who is not among us right now. :disappointed:

This module is highly customizable and new sources can easily add just by implementing source interface.
