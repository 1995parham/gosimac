# Go Si Mac
## Introduction
Automatically Download and [If you want] Update your wallpaper.
Personally I wrote this for having fun and helping one of my friends :P
## How to change your OS X wallpaper from terminal :)
Form OS X Mavericks, all settings stored in sqlite library in following
path:

`~/Library/Application Support/Dock/desktoppicture.db`
## How to create wallpaper sideshow in ubuntu :)
First put configuration of GoSiMac into gnome desktop configuration folder
with following command:
```sh
dst=$HOME/.local/share/gnome-background-properties
test -d "$dst" || mkdir -p "$dst" && cp gnome/gosimac-config.xml "$dst"
```
