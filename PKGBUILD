# Maintainer: NXShock <nxshock@gmail.com>

pkgname=ddns
pkgver=0.0
pkgrel=0
pkgdesc="Simple dynamic DNS updater"
arch=('x86_64' 'armv7h' 'aarch64')
url="https://github.com/nxshock/$pkgname"
license=('MIT')
makedepends=('git' 'go')
backup=()
options=("!strip")
source=('git://github.com/nxshock/ddns.git')
sha256sums=('SKIP')

prepare() {
	msg2 "Download dependencies"
	
	export GOPATH="$startdir"
	
	cd "$srcdir/$pkgname"
	make deps
}

build() {
	cd "$srcdir/$pkgname"
	make build
}

package() {
	cd "$srcdir/$pkgname"
	
	export DESTDIR="$pkgdir"
	make install
}
