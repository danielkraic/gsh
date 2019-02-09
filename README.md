# gsh

[![Build Status](https://travis-ci.org/danielkraic/gsh.svg?branch=master)](https://travis-ci.org/danielkraic/gsh)

quickly ssh connect

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
- name: myserver1
  username: user1
  hostname: server1.localhost
  port: 22
- name: myserver2
  username: user2
  hostname: server2.localhost
- name: myserver3
  hostname: server3.localhost
```

## Usage

```bash
# print help
gsh -h
# select server and connect using ssh
gsh
# select server and print ssh command to connect
gsh -p
```
