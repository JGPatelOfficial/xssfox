---
title: Installation
redirect_from: /docs/installation/
nav_order: 2
toc: true
layout: page
---

# Installation Guide

This guide provides detailed instructions on how to install XSSFox using various methods. Choose the method that best suits your environment and technical preferences.

## Using Homebrew

Homebrew is a popular package manager for macOS and Linux. If you're using a system with Homebrew available, this is the quickest and easiest way to install XSSFox.

### Install Homebrew

If you haven't installed Homebrew yet, you can install it by running the following command in your terminal:

```shell
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

### Install XSSFox

Once Homebrew is installed, you can install XSSFox with a single command:

```shell
brew install xssfox
```

After installation, verify it's working by running:

```shell
xssfox --version
```

For more details about the XSSFox Homebrew package, you can visit the [Homebrew Formula page for XSSFox](https://formulae.brew.sh/formula/xssfox).

## Using Snapcraft

Snapcraft is a package manager for Linux that works across many distributions. It provides containerized applications that run consistently regardless of the underlying system.

### Install Snapcraft

To install Snapcraft on your Linux distribution, please refer to the official documentation: [Installing snapd](https://snapcraft.io/docs/installing-snapd).

### Install XSSFox

Once Snapcraft is installed, you can install XSSFox by running:

```shell
sudo snap install xssfox
```

Verify the installation with:

```shell
xssfox --version
```

## Using Nixpkgs (Nix/NixOS)

XSSFox is available in the Nix package collection, making installation straightforward for Nix and NixOS users. Please note that the latest version may only be available in the `unstable` channel.

### Install Nix

Ensure you have Nix installed on your system. You can find installation instructions on the [official Nix website](https://nixos.org/download/).

### Install XSSFox

To install and use XSSFox in a shell session, run:

```bash
nix-shell -p xssfox
 ```

## From Source

Building from source gives you the most up-to-date version of XSSFox and allows for customization if needed.

### Prerequisites

Ensure you have Go (version 1.16 or later recommended) installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

### Install XSSFox

To install the latest version of XSSFox from source, run:

```bash
go install github.com/JGPatelOfficial/xssfox@latest
```

Make sure your Go bin directory is in your PATH. If you haven't set it up, you can add the following to your shell configuration file (e.g., `.bashrc` or `.zshrc`):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Note: The installed version might differ slightly from the latest release as `go install` references the main branch.

## Using Docker

Docker provides a consistent environment for running XSSFox without worrying about dependencies or system configurations. This is especially useful for CI/CD pipelines or isolated testing environments.

### Pull the Latest Docker Image

To pull the latest Docker image of XSSFox, run:

```bash
# From Docker Hub
docker pull JGPatelOfficial/xssfox:latest

# Or from GitHub Container Registry
docker pull ghcr.io/JGPatelOfficial/xssfox:latest
```

### Run XSSFox Using Docker

You can run XSSFox using Docker with the following command:

```bash
# Using Docker Hub image
docker run -it JGPatelOfficial/xssfox:latest /app/xssfox url https://www.example.com

# Using GitHub Container Registry image
docker run -it ghcr.io/JGPatelOfficial/xssfox:latest /app/xssfox url https://www.example.com
```

For scanning local files or directories, you'll need to mount them to the container:

```bash
docker run -it -v $(pwd):/data JGPatelOfficial/xssfox:latest /app/xssfox file /data/targets.txt
```

### Interactive Docker Shell

For an interactive shell within the Docker container (useful for more complex operations):

```bash
# Using Docker Hub image
docker run -it JGPatelOfficial/xssfox:latest /bin/bash

# Using GitHub Container Registry image
docker run -it ghcr.io/JGPatelOfficial/xssfox:latest /bin/bash
```

Once inside the container, you can run XSSFox directly:

```bash
./xssfox --help
```

## Troubleshooting Common Installation Issues

If you encounter issues during installation, try the following:

1. **PATH Issues**: Ensure the installation directory is in your PATH
2. **Permission Errors**: Use `sudo` for commands that require elevated privileges
3. **Version Conflicts**: Check if you have multiple versions installed

For more help, please open an issue on the [XSSFox GitHub repository](https://github.com/JGPatelOfficial/xssfox/issues).
