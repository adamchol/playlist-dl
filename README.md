# playlist-dl CLI 

playlist-dl is a simple CLI tool for downloading videos from m3u8 playlists, written in Golang for practice.

## Installation

### Install From Source with Go
**Prerequisite:** Ensure you have [Go](https://go.dev/dl/) installed
1. Clone the Repository:
```sh
git clone https://github.com/richardshank/playlist-dl.git
cd playlist-dl
```
2. Build the CLI:
```sh
go build
```
3. Move the Binary to Your PATH:
```sh
mv playlist-dl /usr/local/bin/
```

### Install Using Go Install
**Prerequisite:** Ensure you have [Go](https://go.dev/dl/) installed
```sh
go install github.com/richardshank/playlist-dl@latest
```

### Install with Pre-built Binaries
Download the pre-built binaries for your platform from the [releases page](https://github.com/richardshank/playlist-dl/releases)
1. Download the Binary:
```sh
# MacOS
wget https://github.com/richardshank/playlist-dl/releases/download/0.1.0/playlist-dl-macos.zip

# Linux
wget https://github.com/richardshank/playlist-dl/releases/download/0.1.0/playlist-dl-linux.zip

# Windows
wget https://github.com/richardshank/playlist-dl/releases/download/0.1.0/playlist-dl-windows.zip
```
2. Extract the Binary:
```sh
# MacOS
unzip playlist-dl-macos.zip
```
3. Move the Binary to Your PATH:
```sh
# MacOS/Linux
mv playlist-dl /usr/local/bin/

# Windows
mv playlist-dl C:\Windows\System32
# or
move playlist-dl C:\Windows\System32
```

### Verify the Installation
```sh
playlist-dl -h
```

## Usage
Simple usage of the tool with url to m3u8 playlist:
```sh
playlist-dl url "https://YourUrl.m3u8"
```

Changing the name of the output file:
```sh
playlist-dl url "https://YourUrl.m3u8" -o video
```

