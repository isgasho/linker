# linker
[![Build Status](https://cloud.drone.io/api/badges/domust/linker/status.svg?ref=refs/heads/master)](https://cloud.drone.io/domust/linker)
[![GoDoc](https://godoc.org/github.com/domust/linker?status.svg)](https://godoc.org/github.com/domust/linker)
[![Go Report Card](https://goreportcard.com/badge/domust/linker)](https://goreportcard.com/report/domust/linker)
[![codebeat badge](https://codebeat.co/badges/138238ba-db9d-4c8a-a10e-a4f121cbcc71)](https://codebeat.co/projects/github-com-domust-linker-master)

Linker is a declarative symlink manager in the spirit of terraform.

## Motivation

When customizing Linux distros, one has to deal with a lot of configs.
People figured out that keeping them in one git repository simplifies
things a bit, but one problem still remains - putting the configs into
the appropriate destinations on the system. In order to solve this
problem people started using GNU stow, which has its own share of problems:

1. It's written in Perl, which is not an option when you're dealing with
systems like Alpine Linux, where its not installed by default. Installing
it just to manage configuration removes the whole point of using such a
minimal distro in the first place.

1. Maintaining arbitrary directory structure is annoying and limiting.
Stow takes in an arbitrary directory structure as an argument, parses it
and puts the symlink based on the directory structure one directory up in
the file system tree from where stow was executed. For example, in order
to version `.bashrc` in this way, one would have to create a git repository
in `~/dotfiles` with `.bashrc` being placed in `~/dotfiles/bash/.bashrc`.
Then `cd ~/dotfiles && stow bash/` would be required to execute. Which
in effect would do the same as `cd ~/dotfiles && ln -s ../.bashrc bash/.bashrc`.
This means one of the following:

* All of the configs in the repository have to go to the same target directory
in order to be able to place them efficiently, for example, `~/`.

* One has to keep jumping to various directories and execute stow accordingly.

## Features

* Single-file configuration
* Declarative syntax
* Feels like terraform, but without statefiles

## Installation

Download binary for your system from the release page.

## How to use?

1. Create a file named `main.hcl`. The syntax is:

```hcl
symlink {
    source = "path/to/file1"
    target = "path/to/symlink1"
}

symlink {
    source = "path/to/file2"
    target = "path/to/symlink2"
}

symlink {
    source = "path/to/file3"
    target = "path/to/symlink3"
}
```

1. Run `linker` in the same directory as `main.hcl`.
