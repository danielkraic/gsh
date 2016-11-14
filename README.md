# gsh

[![Build Status](https://travis-ci.org/danielkraic/gsh.svg?branch=master)](https://travis-ci.org/danielkraic/gsh)

quickly get ssh/scp commands

## Requirements

* [golang](https://golang.org/doc/install)
* [git](https://git-scm.com/)

## Installation

```bash
export GOPATH=$HOME/go
go get -v github.com/danielkraic/gsh
```

## Configuration

confiruration file `~/.config/gsh.yml` 

```yaml
---
- server:
  name: server1
  username: user1
  hostname: server1.localhost
  port: 22
- server:
  name: server2
  username: user2
  hostname: server2.localhost
- server:
  name: server3
  hostname: server3.localhost
```

## Usage

```bash
# print help
gsh -h
# list all servers (with ssh command) using default config ~/.config/gsh.yml
gsh -p
# list all servers (with ssh command) using custom config file
gsh -c custom_config.yml -p

# Connect to server (choose server using single pattern)
gsh myserver1
# Connect to server (choose server using multiple patterns)
gsh myserver1 user2

# Upload file to server
gsh -f src_file.txt -u dest_path.txt myserver1

# Download file from server
gsh -d remote_path.txt -f dest_file.txt myserver1

# Print ssh command to server (choose server using single pattern)
gsh -p myserver1
# Print ssh command to server (choose server using multiple patterns)
gsh -p myserver1
# Print scp command to upload file to server
gsh -p -f src_file.txt -u dest_path.txt myserver1
# Print scp command to download file from server
gsh -p -d remote_path.txt -f dest_file.txt myserver1
```