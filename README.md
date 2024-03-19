<h1 align="center"> Go Si Mac </h1>

<p align="center">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/1995parham/gosimac?style=for-the-badge&logo=github">
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/1995parham/gosimac/lint.yaml?style=for-the-badge&logo=github">
  <img alt="GitHub Release Date" src="https://img.shields.io/github/release-date/1995parham/gosimac?style=for-the-badge&logo=github">
  <img alt="AUR Version" src="https://img.shields.io/aur/version/gosimac?style=for-the-badge&logo=archlinux">
</p>

## Introduction

_GoSiMac_ downloads Bing's daily wallpapers, Unsplash's random images, etc. for you to have a beautiful wallpaper on your desktop whenever you want.
Personally, I wrote this to have fun and help one of my friends who are not among us right now. :disappointed:

## Usage

```bash
gosimac rev-4cbe101-dirty
Fetch the wallpaper from Bings, Unsplash...

Usage:
  GoSiMac [command]

Available Commands:
  bing        fetches images from https://bing.com
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  unsplash    fetches images from https://unsplash.org

Flags:
  -h, --help          help for GoSiMac
  -p, --path string   A path to where photos are stored (default "/home/parham/Pictures/GoSiMac")

Use "GoSiMac [command] --help" for more information about a command.

```

As an example, the following command downloads 10 images from unsplash while using Tehran as a search query.
Please note that the proxy setup is related to Iranian sanctions and you may not need to setup any proxy
to use gosimac.

```bash
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

By default, _gosimac_ stores images in `$XDG_PICTURES_DIR/GoSiMac (e.g. $HOME/Pictures/GoSiMac)`.

## Contribution

For adding new source you only need to create a new sub-command in `cmd` package
and then calling your new source with provided `path`. Also for saving images
you can use the following helper function:

```go
func (u *Unsplash) Store(name string, content io.ReadCloser) {
        path := path.Join(
                u.Path,
                fmt.Sprintf("%s-%s.jpg", u.Prefix, name),
        )

        if _, err := os.Stat(path); err == nil {
                pterm.Warning.Printf("%s is already exists\n", path)

                return
        }

        file, err := os.Create(path)
        if err != nil {
                pterm.Error.Printf("os.Create: %v\n", err)

                return
        }

        bytes, err := io.Copy(file, content)
        if err != nil {
                pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
        }

        if err := file.Close(); err != nil {
                pterm.Error.Printf("(*os.File).Close: %v", err)
        }

        if err := content.Close(); err != nil {
                pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
        }
}
```
