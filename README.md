# DDNS

Simple dynamic DNS updater written in Go.

## Install

### ![Arch Linux Logo](https://www.archlinux.org/favicon.ico) Arch Linux

Create package:

```sh
curl -O https://raw.githubusercontent.com/nxshock/ddns/master/PKGBUILD
makepkg
```

Install package:

```sh
pacman -U ddns-0.0-0-x86_64.pkg.tar.xz
```

### Other distros

You can build package manually with [Makefile](https://github.com/nxshock/ddns/blob/master/Makefile).

## Configure

The `ddns` configuration files are stored in `/etc/ddns/` and example
configuration files are available in `/etc/ddns/examples/`.

To use an example configuration, simply copy it from `/etc/ddns/examples/` to
`/etc/ddns/` and configure it with your credentials.

## Usage

### Execute update

One time update can be executed by using:

```sh
systemctl start ddns.service
```

### Automatic updates

To enable automatic updates use:

```sh
systemctl enable ddns.timer
systemctl start ddns.timer
```

### Logs

Logs are accessible by:

```sh
journalctl -u ddns.service
```

## Supported DDNS providers

Name                                | File
----------------------------------- | ----------------------------------
[Duck DNS](https://www.duckdns.org) | /etc/examples/duckdns.toml.example
[hldns.ru](https://hldns.ru)        | /etc/examples/hldns.toml.example
