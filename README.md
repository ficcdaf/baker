# Baker

A simple file backup/copy utility. Baker automates the simple yet incredibly
common task of creating backup copies of files and directories, such as
`cp file file.bak`. Doing this is not hard, but all that typing can get
annoying, especially if you do it often!

## Installation

First, ensure that Go is installed, and that its `bin/` directory is inside your
`$PATH`. For example in `.zshrc`:

```bash
export PATH="$PATH:$HOME/go/bin"
```

Then, run the following command to install the binary:

```bash
go install github.com/ficcdaf/baker@latest
```

To uninstall Baker, simply delete the binary:

```Bash
# If you Go bin directory is in $HOME:
rm $HOME/go/bin/baker
```

## Usage

Using Baker is simple. It takes a single argument, which can be a file or a
directory. It will create a copy of its target and append to its extension
`.bak` if it's a normal file, and `.d.bak` if it's a directory. By default,
Baker will not overwrite a backup that already exists, but you may force it with
the `-f` flag.

```Bash
# Creates myfile.txt.bak
baker myfile.txt
# Creates mydirectory.d.bak
baker mydirectory
# To overwrite an existing backup, use the -f flag:
baker -f myfile.txt
baker -f mydirectory
# To view help:
baker -h
```

## Why?

This could have been a very simple shell alias or function. But I like
programming, I like Go, so I decided to make utility instead. There's no reason
_not_ to release your own commands and utilities, even if they're dead simple!

## Roadmap

- `-n` flag to create a backup with a unique name if the file already exists.
- Accept multiple input arguments.
- Specify output location.
