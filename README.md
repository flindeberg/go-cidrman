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

# Project status and progress

## Findings about the original project

A lot of work was done in the `ipv6-experimental` branch back in Feb and Dec 2017.
In June 2019 some of the *metafiles* was created or updated in the `master` branch, without changes to the code it self.

The continued work in this fork will merge the work done in the `ipv6-experimental` alongside new development.

## Where are we now?

IPv6 merge support has been in the `main` branch since Mar 2022.
The `ipv6-additions` branch was created to create a **PR** into the original project for the IPv6 merge support.

As of Apr 2022 the `merge-experimental` branch was used to merge the old work in `ipv6-experimental`
with the new stuff in `main`. At that point `ipv6-experimental` was removed as it's not relevant in this fork any more.

## Upcoming

Inital work on new `removeCIDRs` fucntions to be able to remove/exclude CIDR blocks or IP ranges. As of Apr 2022 the
internal code supports IPv4. Next step will be IPv6 support.
