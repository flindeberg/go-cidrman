# go-cidrman
Golang CIDR block management module/package.

Inspired by the Python netaddr library:
* https://github.com/drkjam/netaddr/

This is a fork from [EvilSuperstars/go-cidrman](https://github.com/EvilSuperstars/go-cidrman).

## Background

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of
your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your
home directory outside of the standard GOPATH (i.e `$HOME/git/GitHub/Netnod/go-cidrman/`).

# Build instructions

## Clone repository

Clone repository to `$HOME/git/GitHub/Netnod/go-cidrman/`:

```sh
$ mkdir -p $HOME/git/GitHub/Netnod/; cd $HOME/git/GitHub/Netnod/
$ git clone git@github.com:Netnod/go-cidrman.git
```

## Build

```sh
$ cd $HOME/git/GitHub/Netnod/go-cidrman/
$ make build
```

## Test

```sh
$ cd $HOME/git/GitHub/Netnod/go-cidrman/
$ make test
```

# Findings about the original project

A lot of work was done in the `ipv6-experimental` branch, 18 months later some of the *metafiles* was updated in the `master` branch.

The continued work in this fork will merge the work done in the `ipv6-experimental` alongside new development.
