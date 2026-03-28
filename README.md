<div align="center">
  <br>
  <img src="docs/images/logo.png" alt="xssfox" width="400px;">
</div>
<p align="center">
  <a href="https://github.com/JGPatelOfficial/xssfox/releases/latest"><img src="https://img.shields.io/github/v/release/JGPatelOfficial/xssfox?style=for-the-badge&logoColor=%2330365e&label=xssfox&labelColor=%2330365e&color=%2330365e"></a>
  <a href="https://xssfox.hahwul.com/page/overview/"><img src="https://img.shields.io/badge/documents---.svg?style=for-the-badge&labelColor=%2330365e&color=%2330365e"></a>
  <a href="https://x.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=for-the-badge&logo=x&labelColor=%2330365e&color=%2330365e"></a>
  <a href="https://github.com/JGPatelOfficial/xssfox/blob/main/CONTRIBUTING.md"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=for-the-badge&labelColor=%2330365e&color=%2330365e"></a>
</p>

XSSFox is a powerful open-source tool that focuses on automation, making it ideal for quickly scanning for XSS flaws and analyzing parameters. Its advanced testing engine and niche features are designed to streamline the process of detecting and verifying vulnerabilities.

## Key features

* Modes: `URL`, `SXSS`, `Pipe`, `File`, `Server`, `Payload`
* Discovery: Parameter analysis, static analysis, BAV testing, parameter mining
* XSS Scanning: Reflected, Stored, DOM-based, with optimization and DOM/headless verification
* HTTP Options: Custom headers, cookies, methods, proxy, and more
* Output: JSON/Plain formats, silence mode, detailed reports
* Extensibility: REST API, custom payloads, remote wordlists

And the various options required for the testing :D

## Installation
### Homebrew (macOS/Linux)
```bash
brew install xssfox

# https://formulae.brew.sh/formula/xssfox
```

### Snapcraft (Ubuntu)
```bash
sudo snap install xssfox
```

### Nixpkgs (NixOS)

A package is available for Nix or NixOS users. Keep in mind that the latest releases might only
be present in the `unstable` channel.

```bash
nix-shell -p xssfox
```

### From Source

```bash
go install github.com/JGPatelOfficial/xssfox@latest
```

See [Installation guide](https://xssfox.hahwul.com/docs/installation/) for details.

## Usage
```bash
xssfox [mode] [target] [flags] 
```

* Single URL: `xssfox url http://example.com -b https://callback`
* File Mode: `xssfox file urls.txt --custom-payload mypayloads.txt`
* Pipeline: `cat urls.txt | xssfox pipe -H "AuthToken: xxx"`

Check the [Usage](https://xssfox.hahwul.com/page/usage/) and [Running](https://xssfox.hahwul.com/page/running/) documents for more examples.

## Contributing
if you want to contribute to this project, please see [CONTRIBUTING.md](https://github.com/JGPatelOfficial/xssfox/blob/main/CONTRIBUTING.md) and Pull-Request with cool your contents.

[![](/CONTRIBUTORS.svg)](https://github.com/JGPatelOfficial/xssfox/graphs/contributors)

## About the Name
As for the name, Dal([달](https://en.wiktionary.org/wiki/달)) is the Korean word for "moon," while "Fox" stands for "Finder Of XSS" or 🦊

![](docs/images/illust.jpg)
