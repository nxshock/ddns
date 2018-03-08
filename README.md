# DDNS

Simple dynamic DNS updater written in Go.

## Install

Create package file for Arch Linux:

```sh
curl -O https://raw.githubusercontent.com/nxshock/ddns/master/PKGBUILD
makepkg
```

Install package:

```sh
pacman -U ddns-0.0-0-x86_64.pkg.tar.xz
```

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
