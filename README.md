# Go Si Mac

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/1995parham/gosimac/release?label=release&logo=github&style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/1995parham/gosimac/lint?label=lint&logo=github&style=flat-square)
[![Docker Pulls](https://img.shields.io/docker/pulls/1995parham/gosimac.svg?style=flat-square&logo=docker)](https://hub.docker.com/r/1995parham/gosimac/)
![Docker Image Size (tag)](https://img.shields.io/docker/image-size/1995parham/gosimac/latest?style=flat-square&logo=docker)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/1995parham/gosimac?logo=github&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date/1995parham/gosimac?logo=github&style=flat-square)

## Introduction

_gosimac_ downloads Bing's daily wallpapers, Unsplash's random images, and etc. for you to have a beautiful wallpaper on your desktop whenever you want.
Personally, I wrote this to have fun and help one of my friends who is not among us right now. :disappointed:

## Usage

```sh
Usage:
  GoSiMac [command]

Available Commands:
  bing        fetches images from https://bing.com
  help        Help about any command
  unsplash    fetches images from https://unsplash.org

Flags:
  -h, --help          help for GoSiMac
  -n, --number int    The number of photos to return (default 10)
  -p, --path string   A path to store the photos (default "/home/parham/Pictures/GoSiMac")
  -v, --version       version for GoSiMac
```

As an example, the following command downloads 10 images from unsplash while using Tehran as a search query.

```sh
export http_proxy="http://127.0.0.1:1080"
export https_proxy="http://127.0.0.1:1080"

gosimac u -q Tehran -n 10
```

```powershell
set http_proxy "http://127.0.0.1:1080"
set https_proxy "http://127.0.0.1:1080"
$env:HTTP_PROXY = "http://127.0.0.1:1080"
$env:HTTPS_PROXY = "http://127.0.0.1:1080"

gosimac u -q Tehran -n 10

```

By default, _gosimac_ stores images in `$HOME/Pictures/GoSiMac`.

## Contribution

This module is highly customizable and new sources can easily add just by implementing source interface.

```go
// Source represents source for image background.
type Source interface {
	Init() (int, error)                             // call once on source and return number of available images to fetch
	Name() string                                   // name of source in string format
	Fetch(index int) (string, io.ReadCloser, error) // fetch image from source
}
```

The `Init` method is called on initiation and returns number of available images to download.
Then for each image `Fetch` is called and the result is stored at the user specific location.
By implementing this interface you can create new sources for _gosimac_.
